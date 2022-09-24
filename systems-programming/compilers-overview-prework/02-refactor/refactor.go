package main

import (
	"log"
	"os"

	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

const src string = `package foo

import (
	"fmt"
	"time"
)

func baz() {
	fmt.Println("Hello, world!")
}

type A int

const b = "testing"

func bar() {
	fmt.Println(time.Now())
}`

type decls []dst.Decl

func (s decls) Len() int {
	return len(s)
}
func (s decls) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s decls) Less(i, j int) bool {
	switch first := s[i].(type) {
	case *dst.FuncDecl:
		switch second := s[j].(type) {
		case *dst.FuncDecl:
			return first.Name.Name < second.Name.Name
		default:
			return false
		}
	default:
		switch s[j].(type) {
		case *dst.FuncDecl:
			return true
		}
		return false
	}
}

// Moves all top-level functions to the end, sorted in alphabetical order.
// The "source file" is given as a string (rather than e.g. a filename).
func SortFunctions(src string) (string, error) {
	f, err := decorator.Parse(src)
	if err != nil {
		return "", err
	}
	sort.Stable(decls(f.Decls))
	var result strings.Builder
	decorator.Fprint(&result, f)
	return result.String(), nil
}

func main() {
	f, err := decorator.Parse(src)
	if err != nil {
		log.Fatal(err)
	}

	// Print AST
	err = dst.Fprint(os.Stdout, f, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Convert AST back to source
	err = decorator.Print(f)
	if err != nil {
		log.Fatal(err)
	}
}
