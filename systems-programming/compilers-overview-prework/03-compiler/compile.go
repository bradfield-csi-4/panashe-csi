package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

// Given an AST node corresponding to a function (guaranteed to be
// of the form `func f(x, y byte) byte`), compile it into assembly
// code.
//
// Recall from the README that the input parameters `x` and `y` should
// be read from memory addresses `1` and `2`, and the return value
// should be written to memory address `0`.
func compile(node *ast.FuncDecl) (string, error) {
	var result strings.Builder
	// locals maps variable names to memory addresses (1-based index)
	locals := []string{"x", "y", "", "", "", "", ""}
	for _, stmt := range node.Body.List {
		err := compileStmt(stmt, &result, locals)
		if err != nil {
			return "", err
		}
	}
	result.WriteString("pop 0\nhalt\n")
	return result.String(), nil
}

func find(s string, sSlice []string) (uint, error) {
	for i := range sSlice {
		if sSlice[i] == s {
			return uint(i), nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Could not find %s in %v", s, sSlice))
}

func addVars(varNames, locals []string) error {
	var j int
	for i := 0; i < len(locals); i++ {
		if j >= len(varNames) {
			break
		}
		if locals[i] == "" {
			locals[i] = varNames[j]
			j++
		}
	}
	if j < len(varNames) {
		return errors.New(fmt.Sprintf("Not enough space to store %v", varNames[j:]))
	}
	return nil
}

func compileStmt(stmt ast.Stmt, result *strings.Builder, locals []string) error {
	switch stmt := stmt.(type) {
	case *ast.BlockStmt:
		for _, s := range stmt.List {
			compileStmt(s, result, locals)
		}
	case *ast.ReturnStmt:
		// TODO: handle multiple return
		val, err := compileExpr(stmt.Results[0], locals)
		if err != nil {
			return err
		}
		result.WriteString(val)
	case *ast.AssignStmt:
		// TODO: handle multiple assignment
		compiled, err := compileExpr(stmt.Rhs[0], locals)
		if err != nil {
			return err
		}
		result.WriteString(compiled)
		switch lhs := stmt.Lhs[0].(type) {
		case *ast.Ident:
			pos, err := find(lhs.Name, locals)
			if err != nil {
				return err
			}
			result.WriteString(fmt.Sprintf("pop %d\n", pos+1))
		default:
			return errors.New(fmt.Sprintf("%v cannot be on LHS of assignment", lhs))
		}
	case *ast.IfStmt:
		cond, err := compileExpr(stmt.Cond, locals)
		if err != nil {
			return err
		}
		result.WriteString(cond)
		result.WriteString("jeqz else\n")
		compileStmt(stmt.Body, result, locals)
		result.WriteString("jump if_end\n")
		result.WriteString("label else\n")
		compileStmt(stmt.Else, result, locals)
		result.WriteString("label if_end\n")
	case *ast.DeclStmt:
		switch decl := stmt.Decl.(type) {
		case *ast.GenDecl:
			var varNames []string
			for _, spec := range decl.Specs {
				switch s := spec.(type) {
				case *ast.ValueSpec:
					for _, name := range s.Names {
						varNames = append(varNames, name.Name)
					}
				default:
					return errors.New(fmt.Sprintf("Unrecognized spec %v of type %T", s, s))
				}
			}
			err := addVars(varNames, locals)
			if err != nil {
				return err
			}
		default:
			return errors.New(fmt.Sprintf("Unrecognized declaration %v of type %T", decl, decl))
		}
	case *ast.ForStmt:
		cond, err := compileExpr(stmt.Cond, locals)
		if err != nil {
			return err
		}
		result.WriteString("label loop_start\n")
		result.WriteString(cond)
		result.WriteString("jeqz after_loop\n")
		compileStmt(stmt.Body, result, locals)
		result.WriteString("jump loop_start\n")
		result.WriteString("label after_loop\n")
	}
	return nil
}

func compileExpr(expr ast.Expr, locals []string) (string, error) {
	switch expr := expr.(type) {
	case *ast.BasicLit:
		if !(expr.Kind == token.INT) {
			return "", errors.New(fmt.Sprintf("Bad literal expression %v", expr))
		} else {
			parsedVal, err := strconv.Atoi(expr.Value)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("pushi %d\n", parsedVal), nil
		}
	case *ast.ParenExpr:
		return compileExpr(expr.X, locals)
	case *ast.Ident:
		pos, err := find(expr.Name, locals)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("push %d\n", pos+1), nil
	case *ast.BinaryExpr:
		xComp, err := compileExpr(expr.X, locals)
		if err != nil {
			return "", err
		}
		yComp, err := compileExpr(expr.Y, locals)
		if err != nil {
			return "", err
		}
		var result strings.Builder
		result.WriteString(fmt.Sprintf("%s%s", xComp, yComp))
		asmOp, err := translateOpTokenToAsm(expr.Op)
		result.WriteString(asmOp)
		return result.String(), nil
	}
	return "", nil
}

func translateOpTokenToAsm(tok token.Token) (string, error) {
	switch tok {
	case token.ADD:
		return "add\n", nil
	case token.SUB:
		return "sub\n", nil
	case token.MUL:
		return "mul\n", nil
	case token.QUO:
		return "div\n", nil
	case token.EQL:
		return "eq\n", nil
	case token.LSS:
		return "lt\n", nil
	case token.GTR:
		return "gt\n", nil
	case token.NEQ:
		return "neq\n", nil
	case token.LEQ:
		return "leq\n", nil
	case token.GEQ:
		return "geq\n", nil

	default:
		return "", errors.New(fmt.Sprintf("Unrecognized token %v", tok))
	}

}
