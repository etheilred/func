package main

import (
	"func/ast"
	"func/lexer"
	"func/parser"
	"testing"
)

func TestExpression(t *testing.T) {
	inp := "((177*(5/23)/44-5)*(6----3));"
	ans := int64(-45)
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(inp))
	q, err := p.Parse(s)
	if err != nil {
		t.Error(err)
	}
	res := q.(ast.Evaluator).Evaluate()
	if res != ans {
		t.Errorf("Expected: %d\nGot: %d", ans, res)
	}
}

func TestFunction(t *testing.T) {
	inp := `
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
	ans := int64(7)
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(inp))
	q, err := p.Parse(s)
	if err != nil {
		t.Error(err)
	}
	res := q.(ast.Evaluator).Evaluate()
	if res != ans {
		t.Errorf("Expected: %d\nGot: %d", ans, res)
	}
}
