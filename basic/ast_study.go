package basic

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func AstStudy() {

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset

	// 解析ast_study.go的代码
	f, err := parser.ParseFile(fset, "ast_study.go", nil, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

}
