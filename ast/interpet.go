package ast

import "fmt"

type scope map[string]int64

type tid []scope

type function struct {
	Params []string
	Body   Evaluator
}

var functions map[string]*function = make(map[string]*function)

var scopes tid = make(tid, 1)

func addFunction(f *FuncNode) {
	functions[f.ID] = &function{f.ParamList.ParamNames, f.Body}
}

func callFunction(f *FuncCallNode) (int64, error) {
	newScope()
	if len(f.Args.Expressions) != len(functions[f.ID].Params) {
		return 0, fmt.Errorf("Invalid parameter list")
	}
	for i := 0; i < len(f.Args.Expressions); i++ {
		getTopScope()[functions[f.ID].Params[i]] = f.Args.Expressions[i].Evaluate()
	}
	tmp := functions[f.ID].Body.Evaluate()
	popScope()
	return tmp, nil
}

func newScope() {
	scopes = append(scopes, make(scope))
}

func getTopScope() scope {
	return scopes[len(scopes)-1]
}

func popScope() {
	scopes = scopes[:len(scopes)-1]
}

func getValue(id string) (int64, error) {
	// fmt.Println(getTopScope())
	for i := len(scopes) - 1; i >= 0; i-- {
		if _, ok := scopes[i][id]; ok {
			return scopes[i][id], nil
		}
	}
	return 0, fmt.Errorf("No such variable: %s", id)
}

func setValue(id string, val int64) {
	getTopScope()[id] = val
}
