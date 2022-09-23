package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

// Given an expression containing only int types, evaluate
// the expression and return the result.
func Evaluate(expr ast.Expr) (int, error) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind != token.INT {
			return 0, errors.New(fmt.Sprintf("%v is not an integer", expr))
		}
		val, err := strconv.Atoi(e.Value)
		if err != nil {
			return 0, err
		}
		return val, nil
	case *ast.BinaryExpr:
		left, err := Evaluate(e.X)
		if err != nil {
			return 0, nil
		}
		right, err := Evaluate(e.Y)
		if err != nil {
			return 0, nil
		}
		switch e.Op {
		case token.ADD:
			return left + right, nil
		case token.SUB:
			return left - right, nil
		case token.MUL:
			return left * right, nil
		}
	case *ast.ParenExpr:
		return Evaluate(e.X)
	case *ast.UnaryExpr:
		val, err := Evaluate(e.X)
		if err != nil {
			return 0, err
		}
		if e.Op == token.SUB {
			return -1 * val, nil
		}
	default:
		fmt.Printf("Do not recognize type %e\n", e)
	}
	return 0, nil
}

func main() {
	expr, err := parser.ParseExpr("1 + 2 - 3 * 4")
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	err = ast.Print(fset, expr)
	if err != nil {
		log.Fatal(err)
	}
}
