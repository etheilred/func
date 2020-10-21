package main

import (
	"fmt"
	"func/ast"
	"func/lexer"
	"func/parser"
	"os"
)

var inp = `
fn fib n => if n == 0 then 1 else if n == 1 then 1 else fib(n-1) + fib(n-2);
fib(6);
`

func main() {
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(inp))
	sum, err := p.Parse(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(sum)
	fmt.Println(sum.(ast.Evaluator).Evaluate())
	// fmt.Println(sum.(ast.Evaluator).Evaluate())
	// exprs := sum.(*ast.ExprList)
	// for _, e := range exprs.Expressions {
	// 	fmt.Println(e.Evaluate())
	// }
}
