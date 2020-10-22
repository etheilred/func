package main

import (
	"fmt"
	"func/ast"
	"func/lexer"
	"func/parser"
	"os"
)

var inp = `
fn fib n => if n == 0 then 
				1 
			else if n == 1 then 
				1 
			else 
				fibT(n);

fn fibT n => fib(n-1) + fib(n-2);

fn sqrt n =>	if n < 0 then 
					-1 
				else 
					_sqrt(0, n, n);

fn sqr n => n * n;
fn _sqrt l r n => 	if l >= r-1 then 
						l
					else if sqr((l+r)/2) < n then
						_sqrt((l+r)/2, r, n)
					else if sqr((l+r)/2) > n then
						_sqrt(l, (l+r)/2, n)
					else 
						(l+r)/2;

sqrt(55);
`

var inp2 string = `
fn sqr n => n * n;
fn binpow n x => 	if x == 0 then
						1
					else if (x/2)*2 == x then 
						sqr(binpow(n, x/2))
					else
						n * sqr(binpow(n, x/2));
binpow(-2, 3);
`

func main() {
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(inp))
	sum, err := p.Parse(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(sum)
	fmt.Println(sum.(ast.Evaluator).Evaluate())
	// fmt.Println(sum.(ast.Evaluator).Evaluate())
	// exprs := sum.(*ast.ExprList)
	// for _, e := range exprs.Expressions {
	// 	fmt.Println(e.Evaluate())
	// }
}
