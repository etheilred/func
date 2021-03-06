// Code generated by gocc; DO NOT EDIT.

package parser

import (
		"func/ast"
	)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : StmtList	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StmtList : StmtList Stmt	<< ast.AppendStmtListNode(X[0], X[1]) >>`,
		Id:         "StmtList",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStmtListNode(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `StmtList : Stmt	<< ast.NewStmtListNode(X[0]) >>`,
		Id:         "StmtList",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtListNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Stmt : ExprList ";"	<< ast.NewStmtNode(X[0]) >>`,
		Id:         "Stmt",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Stmt : Func ";"	<< ast.NewStmtNode(X[0]) >>`,
		Id:         "Stmt",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Func : fn id Params arrow ExprList	<< ast.NewFuncNode(X[1], X[2], X[4]) >>`,
		Id:         "Func",
		NTType:     3,
		Index:      5,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFuncNode(X[1], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `Func : fn id arrow ExprList	<< ast.NewFuncNode(X[1], nil, X[3]) >>`,
		Id:         "Func",
		NTType:     3,
		Index:      6,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFuncNode(X[1], nil, X[3])
		},
	},
	ProdTabEntry{
		String: `Params : Params id	<< ast.AppendParamsNode(X[0], X[1]) >>`,
		Id:         "Params",
		NTType:     4,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendParamsNode(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Params : id	<< ast.NewParamsNode(X[0]) >>`,
		Id:         "Params",
		NTType:     4,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewParamsNode(X[0])
		},
	},
	ProdTabEntry{
		String: `ExprList : ExprList "," Expr	<< ast.NewExprList(X[0], X[2]) >>`,
		Id:         "ExprList",
		NTType:     5,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExprList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `ExprList : Expr	<< ast.NewExprList(nil, X[0]) >>`,
		Id:         "ExprList",
		NTType:     5,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExprList(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : T1	<< X[0], nil >>`,
		Id:         "Expr",
		NTType:     6,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Local	<< X[0], nil >>`,
		Id:         "Expr",
		NTType:     6,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : IfExpr	<< X[0], nil >>`,
		Id:         "Expr",
		NTType:     6,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfExpr : if RelExpr then Expr else Expr	<< ast.NewIfNode(X[1], X[3], X[5]) >>`,
		Id:         "IfExpr",
		NTType:     7,
		Index:      14,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfNode(X[1], X[3], X[5])
		},
	},
	ProdTabEntry{
		String: `RelExpr : Expr gt Expr	<< ast.NewGtNode(X[0], X[2]) >>`,
		Id:         "RelExpr",
		NTType:     8,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGtNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `RelExpr : Expr lt Expr	<< ast.NewLtNode(X[0], X[2]) >>`,
		Id:         "RelExpr",
		NTType:     8,
		Index:      16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLtNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `RelExpr : Expr eq Expr	<< ast.NewEqNode(X[0], X[2]) >>`,
		Id:         "RelExpr",
		NTType:     8,
		Index:      17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEqNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `RelExpr : Expr le Expr	<< ast.NewLeNode(X[0], X[2]) >>`,
		Id:         "RelExpr",
		NTType:     8,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLeNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `RelExpr : Expr ge Expr	<< ast.NewGeNode(X[0], X[2]) >>`,
		Id:         "RelExpr",
		NTType:     8,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGeNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `RelExpr : Expr ne Expr	<< ast.NewNeNode(X[0], X[2]) >>`,
		Id:         "RelExpr",
		NTType:     8,
		Index:      20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNeNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Local : let id assign T1	<< ast.NewLocalNode(X[1], X[3]) >>`,
		Id:         "Local",
		NTType:     9,
		Index:      21,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLocalNode(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `T1 : T1 "+" T2	<< ast.NewAddNode(X[0], X[2]) >>`,
		Id:         "T1",
		NTType:     10,
		Index:      22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAddNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `T1 : T1 "-" T2	<< ast.NewSubNode(X[0], X[2]) >>`,
		Id:         "T1",
		NTType:     10,
		Index:      23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSubNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `T1 : T2	<< X[0], nil >>`,
		Id:         "T1",
		NTType:     10,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `T2 : T2 "*" Factor	<< ast.NewMulNode(X[0], X[2]) >>`,
		Id:         "T2",
		NTType:     11,
		Index:      25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMulNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `T2 : T2 "/" Factor	<< ast.NewDivNode(X[0], X[2]) >>`,
		Id:         "T2",
		NTType:     11,
		Index:      26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDivNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `T2 : Factor	<< X[0], nil >>`,
		Id:         "T2",
		NTType:     11,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expr ")"	<< X[1], nil >>`,
		Id:         "Factor",
		NTType:     12,
		Index:      28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Factor : int64	<< ast.NewInt64(X[0]) >>`,
		Id:         "Factor",
		NTType:     12,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInt64(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : id	<< ast.NewVarNode(X[0]) >>`,
		Id:         "Factor",
		NTType:     12,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : id "(" ExprList ")"	<< ast.NewFuncCallNode(X[0], X[2]) >>`,
		Id:         "Factor",
		NTType:     12,
		Index:      31,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFuncCallNode(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Factor : Unary	<< X[0], nil >>`,
		Id:         "Factor",
		NTType:     12,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Unary : "-" Factor	<< ast.NewUnaryMinusNode(X[1]) >>`,
		Id:         "Unary",
		NTType:     13,
		Index:      33,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewUnaryMinusNode(X[1])
		},
	},
}
