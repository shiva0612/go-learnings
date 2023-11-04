package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
)

func req() {
	b, _ := ioutil.ReadFile("req")
	sourceCode := string(b)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "example.go", sourceCode, parser.ParseComments)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	// Define a visitor function to traverse the AST and add a type assertion.
	ast.Inspect(node, func(n ast.Node) bool {
		assignStmt, ok := n.(*ast.AssignStmt)
		if !ok {
			return true
		}

		// Check if the assignment statement has an identifier and a function call.
		if len(assignStmt.Lhs) == 1 && len(assignStmt.Rhs) == 1 {
			ident, ok := assignStmt.Lhs[0].(*ast.Ident)
			callExpr, ok := assignStmt.Rhs[0].(*ast.CallExpr)
			if ok && ident != nil && callExpr != nil {
				// Check if the function being called is "ListGet."
				if selExpr, ok := callExpr.Fun.(*ast.Ident); ok {
					if selExpr.Name == "ListGet" {
						// Determine the function call's type by analyzing the argument's value.
						sliceType := callExpr.Args[0].(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Rhs[0].(*ast.CompositeLit).Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
						typeAssert := &ast.TypeAssertExpr{
							X:    assignStmt.Rhs[0],
							Type: &ast.Ident{Name: sliceType},
						}
						// Replace the assignment expression with the one including the type assertion.
						assignStmt.Rhs[0] = typeAssert
					}
				}
			}
		}
		return true
	})

	// Print the modified AST.
	printer.Fprint(os.Stdout, fset, node)
}
