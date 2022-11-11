package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"syscall/js"
)

type Ast struct {}

func (*Ast) Print(this js.Value, args []js.Value) any {
	inputSrc := args[0].String()
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", inputSrc, 0)
	s := ""
	buf := bytes.NewBufferString(s)
	ast.Fprint(buf, fset, f, ast.NotNilFilter)
	return buf.String()
}

func main() {
	a := Ast{}
	
	js.Global().Set("ast", js.ValueOf(
		map[string]any{
			"print": js.FuncOf(a.Print),
		},
	))

	select {} // keep running
}
