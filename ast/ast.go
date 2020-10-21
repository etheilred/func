package ast

import (
	"fmt"
	"func/token"
	"strconv"
	"strings"
)

// Attrib is used for parser internal representation
type Attrib interface{}

// NewInt64 creates new int64 from []byte
func NewInt64(n Attrib) (Attrib, error) {
	res, err := strconv.Atoi(string(n.(*token.Token).Lit))
	return Attrib(&Int64Node{int64(res)}), err
}

// NewAddNode creates new addition node
func NewAddNode(left, right Attrib) (Attrib, error) {
	return Attrib(&AddNode{left.(Evaluator), right.(Evaluator)}), nil
}

// NewSubNode creates new subtraction node
func NewSubNode(left, right Attrib) (Attrib, error) {
	return Attrib(&SubNode{left.(Evaluator), right.(Evaluator)}), nil
}

// NewMulNode creates new subtraction node
func NewMulNode(left, right Attrib) (Attrib, error) {
	return Attrib(&MulNode{left.(Evaluator), right.(Evaluator)}), nil
}

// NewDivNode creates new subtraction node
func NewDivNode(left, right Attrib) (Attrib, error) {
	return Attrib(&DivNode{left.(Evaluator), right.(Evaluator)}), nil
}

// NewExprList creates new expression list
func NewExprList(acc, next Attrib) (Attrib, error) {
	if acc == nil {
		return &ExprList{[]Evaluator{next.(Evaluator)}}, nil
	}
	tmp := acc.(*ExprList)
	tmp.addExpr(next.(Evaluator))
	return tmp, nil
}

// NewLocalNode creates new node for local variable
func NewLocalNode(id, val Attrib) (Attrib, error) {
	return &LocalNode{string(id.(*token.Token).Lit), val.(Evaluator)}, nil
}

// NewStmtNode creates new node for statement
func NewStmtNode(val Attrib) (Attrib, error) {
	return &StmtNode{val.(Evaluator)}, nil
}

// NewVarNode creates new node for variable access
func NewVarNode(id Attrib) (Attrib, error) {
	return &VarNode{string(id.(*token.Token).Lit)}, nil
}

// NewParamsNode creates new node for parameter list
func NewParamsNode(id Attrib) (Attrib, error) {
	return &ParamsNode{[]string{string(id.(*token.Token).Lit)}}, nil
}

// AppendParamsNode appends new parameter to parameter list
func AppendParamsNode(paramsNode, id Attrib) (Attrib, error) {
	tmp := paramsNode.(*ParamsNode)
	tmp.ParamNames = append(tmp.ParamNames, string(id.(*token.Token).Lit))
	return tmp, nil
}

// NewFuncNode creates new function node
func NewFuncNode(id, params, body Attrib) (Attrib, error) {
	if params == nil {
		params = &ParamsNode{}
	}
	tmp := &FuncNode{
		string(id.(*token.Token).Lit),
		params.(*ParamsNode),
		body.(*ExprList),
	}
	addFunction(tmp)
	return tmp, nil
}

// NewFuncCallNode creates new function call node
func NewFuncCallNode(id, args Attrib) (Attrib, error) {
	return &FuncCallNode{
		string(id.(*token.Token).Lit),
		args.(*ExprList),
	}, nil
}

// NewStmtListNode creates new node for statement list
func NewStmtListNode(stmt Attrib) (Attrib, error) {
	return &StmtListNode{[]*StmtNode{stmt.(*StmtNode)}}, nil
}

// AppendStmtListNode appends statement to statement list
func AppendStmtListNode(lst, stmt Attrib) (Attrib, error) {
	tmp := lst.(*StmtListNode)
	tmp.Statements = append(tmp.Statements, stmt.(*StmtNode))
	return tmp, nil
}

// NewGtNode creates new '>' node
func NewGtNode(left, right Attrib) (Attrib, error) {
	return &GtNode{left.(Evaluator), right.(Evaluator)}, nil
}

// NewLtNode creates new '<' node
func NewLtNode(left, right Attrib) (Attrib, error) {
	return &LtNode{left.(Evaluator), right.(Evaluator)}, nil
}

// NewEqNode creates new '==' node
func NewEqNode(left, right Attrib) (Attrib, error) {
	return &EqNode{left.(Evaluator), right.(Evaluator)}, nil
}

// NewLeNode creates new '<=' node
func NewLeNode(left, right Attrib) (Attrib, error) {
	return &LeNode{left.(Evaluator), right.(Evaluator)}, nil
}

// NewGeNode creates new '>=' node
func NewGeNode(left, right Attrib) (Attrib, error) {
	return &GeNode{left.(Evaluator), right.(Evaluator)}, nil
}

// NewNeNode creates new '!=' node
func NewNeNode(left, right Attrib) (Attrib, error) {
	return &NeNode{left.(Evaluator), right.(Evaluator)}, nil
}

// NewIfNode return new if-then-else node
func NewIfNode(cond, t, f Attrib) (Attrib, error) {
	return &IfNode{
		cond.(Evaluator),
		t.(Evaluator),
		f.(Evaluator),
	}, nil
}

// Evaluator used to evaluate ast
type Evaluator interface {
	Evaluate() int64
	fmt.Stringer
}

// Int64Node used to store int64 constants in ast
type Int64Node struct {
	Val int64
}

func (i64 *Int64Node) String() string {
	return fmt.Sprintf("%d", i64.Val)
}

// Evaluate evaluates contants
func (i64 *Int64Node) Evaluate() int64 {
	return i64.Val
}

// AddNode used to store Addition node
type AddNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates sum
func (addNode *AddNode) Evaluate() int64 {
	return (addNode.Left).Evaluate() + (addNode.Right).Evaluate()
}

func (addNode *AddNode) String() string {
	return fmt.Sprintf("(+ %s %s)", addNode.Left.String(), addNode.Right.String())
}

// SubNode represents "-" node
type SubNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates subtraction
func (subNode *SubNode) Evaluate() int64 {
	return subNode.Left.Evaluate() - subNode.Right.Evaluate()
}

func (subNode *SubNode) String() string {
	return fmt.Sprintf("(- %s %s)", subNode.Left.String(), subNode.Right.String())
}

// MulNode represents "*" node
type MulNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates subtraction
func (n *MulNode) Evaluate() int64 {
	return n.Left.Evaluate() * n.Right.Evaluate()
}

func (n *MulNode) String() string {
	return fmt.Sprintf("(* %s %s)", n.Left.String(), n.Right.String())
}

// DivNode represents "/" node
type DivNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates subtraction
func (n *DivNode) Evaluate() int64 {
	return n.Left.Evaluate() / n.Right.Evaluate()
}

func (n *DivNode) String() string {
	return fmt.Sprintf("(/ %s %s)", n.Left.String(), n.Right.String())
}

// ExprList stores sequential expressions
type ExprList struct {
	Expressions []Evaluator
}

// Evaluate evaluates all expressions in expr-list, returning last evaluation result
func (e *ExprList) Evaluate() int64 {
	var tmp int64
	for _, expr := range e.Expressions {
		tmp = expr.Evaluate()
	}
	return tmp
}

func mapStrings(v []Evaluator) []string {
	a := make([]string, 0)
	for _, k := range v {
		a = append(a, k.String())
	}
	return a
}

func (e *ExprList) addExpr(a Evaluator) {
	e.Expressions = append(e.Expressions, a)
}

func (e *ExprList) String() string {
	return fmt.Sprintf("(expressions %s)", strings.Join(mapStrings(e.Expressions), " "))
}

// LocalNode represents "let x = y" in ast
type LocalNode struct {
	ID  string
	Val Evaluator
}

// Evaluate returns variable value
func (e *LocalNode) Evaluate() int64 {
	tmp := e.Val.Evaluate()
	setValue(e.ID, tmp)
	return tmp
}

func (e *LocalNode) String() string {
	return fmt.Sprintf("(let (%s %s))", e.ID, e.Val.String())
}

// StmtNode represents statement
type StmtNode struct {
	Val Evaluator
}

// Evaluate returns variable value
func (e *StmtNode) Evaluate() int64 {
	return e.Val.Evaluate()
}

func (e *StmtNode) String() string {
	return fmt.Sprintf("(stmt %s)", e.Val.String())
}

// VarNode represents variable
type VarNode struct {
	ID string
}

// Evaluate returns variable value
func (e *VarNode) Evaluate() int64 {
	tmp, err := getValue(e.ID)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (e *VarNode) String() string {
	return fmt.Sprintf("%s", e.ID)
}

// ParamsNode represents parameter list
type ParamsNode struct {
	ParamNames []string
}

// Evaluate returns variable value
func (e *ParamsNode) Evaluate() int64 {
	return 0
}

func (e *ParamsNode) String() string {
	return fmt.Sprintf("(params (%s))", strings.Join(e.ParamNames, " "))
}

// FuncNode represents function
type FuncNode struct {
	ID        string
	ParamList *ParamsNode
	Body      *ExprList
}

// Evaluate returns 0
func (e *FuncNode) Evaluate() int64 {
	return 0
}

func (e *FuncNode) String() string {
	return fmt.Sprintf("(defun %s %s %s)", e.ID, e.ParamList.String(), e.Body.String())
}

// FuncCallNode represents function call node
type FuncCallNode struct {
	ID   string
	Args *ExprList
}

// Evaluate evaluates function call
func (e *FuncCallNode) Evaluate() int64 {
	tmp, err := callFunction(e)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (e *FuncCallNode) String() string {
	return fmt.Sprintf("(funcall %s %s)", e.ID, e.Args.String())
}

// StmtListNode represents list of statements
type StmtListNode struct {
	Statements []*StmtNode
}

// Evaluate todo
func (e *StmtListNode) Evaluate() int64 {
	var tmp int64
	for _, e := range e.Statements {
		tmp = e.Evaluate()
	}
	return tmp
}

func mapStringsStmt(v []*StmtNode) []string {
	a := make([]string, 0)
	for _, s := range v {
		a = append(a, s.String())
	}
	return a
}

func (e *StmtListNode) String() string {
	return fmt.Sprintf("(statements %s)", strings.Join(mapStringsStmt(e.Statements), " "))
}

// GtNode represents '>'
type GtNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates '>'
func (e *GtNode) Evaluate() int64 {
	if e.Left.Evaluate() > e.Right.Evaluate() {
		return 1
	}
	return 0
}

func (e *GtNode) String() string {
	return fmt.Sprintf("(> %s %s)", e.Left.String(), e.Right.String())
}

// LtNode represents '<'
type LtNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates '<'
func (e *LtNode) Evaluate() int64 {
	if e.Left.Evaluate() < e.Right.Evaluate() {
		return 1
	}
	return 0
}

func (e *LtNode) String() string {
	return fmt.Sprintf("(< %s %s)", e.Left.String(), e.Right.String())
}

// EqNode represents '=='
type EqNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates '=='
func (e *EqNode) Evaluate() int64 {
	if e.Left.Evaluate() == e.Right.Evaluate() {
		return 1
	}
	return 0
}

func (e *EqNode) String() string {
	return fmt.Sprintf("(== %s %s)", e.Left.String(), e.Right.String())
}

// LeNode represents '<='
type LeNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates '<='
func (e *LeNode) Evaluate() int64 {
	if e.Left.Evaluate() <= e.Right.Evaluate() {
		return 1
	}
	return 0
}

func (e *LeNode) String() string {
	return fmt.Sprintf("(<= %s %s)", e.Left.String(), e.Right.String())
}

// GeNode represents '>='
type GeNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates '>='
func (e *GeNode) Evaluate() int64 {
	if e.Left.Evaluate() >= e.Right.Evaluate() {
		return 1
	}
	return 0
}

func (e *GeNode) String() string {
	return fmt.Sprintf("(>= %s %s)", e.Left.String(), e.Right.String())
}

// NeNode represents '!='
type NeNode struct {
	Left, Right Evaluator
}

// Evaluate evaluates '!='
func (e *NeNode) Evaluate() int64 {
	if e.Left.Evaluate() != e.Right.Evaluate() {
		return 1
	}
	return 0
}

func (e *NeNode) String() string {
	return fmt.Sprintf("(!= %s %s)", e.Left.String(), e.Right.String())
}

// IfNode represents if-then-else
type IfNode struct {
	Condition, TrueFlow, FalseFlow Evaluator
}

// Evaluate evaluates if-then-else expression
func (e *IfNode) Evaluate() int64 {
	if e.Condition.Evaluate() != 0 {
		return e.TrueFlow.Evaluate()
	}
	return e.FalseFlow.Evaluate()
}

func (e *IfNode) String() string {
	return fmt.Sprintf("(if %s (then %s) (else %s))", e.Condition.String(), e.TrueFlow.String(), e.FalseFlow.String())
}
