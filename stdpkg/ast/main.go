package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("main...")
	b, _ := ioutil.ReadFile("now")
	sourceCode := string(b)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "example.go", sourceCode, parser.ParseComments)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	// Define a visitor function to traverse the AST and add a type assertion.
	ast.Inspect(node, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BinaryExpr:
			// Handle binary expressions like '==', '!=', etc.
			if callExpr, ok := n.X.(*ast.CallExpr); ok {
				handleListGet(callExpr, fset, node)
			}
			if callExpr, ok := n.Y.(*ast.CallExpr); ok {
				handleListGet(callExpr, fset, node)
			}
		case *ast.CallExpr:
			// Handle function calls.
			handleListGet(n, fset, node)
		}
		return true
	})

	// Print the modified AST.
	printer.Fprint(os.Stdout, fset, node)

	// Optionally, you can also generate source code from the modified AST.
	var outputCode strings.Builder
	printer.Fprint(&outputCode, fset, node)
	fmt.Println("\n\nCorrected code:")
	fmt.Println(outputCode.String())
}

func handleListGet(callExpr *ast.CallExpr, fset *token.FileSet, node *ast.File) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("handleListGet panic : ", r)
		}
	}()
	if selExpr, ok := callExpr.Fun.(*ast.Ident); ok {
		if selExpr.Name == "ListGet" {
			var typeAssert *ast.TypeAssertExpr
			if _, ok := callExpr.Args[0].(*ast.Ident); ok {
				sliceType := callExpr.Args[0].(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Rhs[0].(*ast.CompositeLit).Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
				typeAssert = &ast.TypeAssertExpr{
					X:    callExpr,
					Type: &ast.Ident{Name: sliceType},
				}
			} else if _, ok := callExpr.Args[0].(*ast.CompositeLit); ok {
				sliceType := callExpr.Args[0].(*ast.CompositeLit).Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
				typeAssert = &ast.TypeAssertExpr{
					X:    callExpr,
					Type: &ast.Ident{Name: sliceType},
				}

			} else {
				panic("cannot define the arg[0] ka type")
			}
			replaceNode(fset, node, callExpr, typeAssert)
		}
	}
}

func replaceNode(fset *token.FileSet, node *ast.File, oldNode, newNode ast.Node) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("replaceNode panic : ", r)
		}
	}()

	for _, decl := range node.Decls {
		ast.Inspect(decl, func(n ast.Node) bool {
			if n == oldNode {
				oldNode = newNode
			}
			return true
		})
	}
}
