package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "h1.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.InterfaceType)
		if ok {
			//fmt.Printf("%d:\n\t", fset.Position(ret.Pos()).Line)
			// printer.Fprint(os.Stdout, fset, ret)
			HandleInterface(ret)
			return true
		}
		return true
	})
}

func HandleInterface(t *ast.InterfaceType) {
	m := t.Methods
	for _, m1 := range m.List {
		fmt.Printf("%s(", m1.Names[0])
		//fmt.Printf("Method params: %#v \n", m1.Type)
		ft, ok := m1.Type.(*ast.FuncType)
		if ok {
			HandleParams(ft.Params)
			fmt.Printf(")(")
			HandleResults(ft.Results)
			fmt.Println(")")
		}

	}
	_ = m
}

func HandleParams(fl *ast.FieldList) {
	for index, m1 := range fl.List {
		//fmt.Printf("Param names:%s\n", m1.Names)
		ident := m1.Type.(*ast.Ident)
		if index > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%s %s", m1.Names[0], ident.Name)
	}
}

func HandleResults(fl *ast.FieldList) {
	for index, m1 := range fl.List {
		// fmt.Printf(":%s\n", m1.Names)
		ident := m1.Type.(*ast.Ident)
		if index > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%s", ident.Name)
	}
}
