package govaluate

import (
	"testing"
)

func TestGetFuncName(t *testing.T) {
	eva, err := NewEvaluableExpressionWithFunctions("testA(testB(var1,var2),var3,var4)", map[string]ExpressionFunction{
		"testA": func(arguments ...interface{}) (interface{}, error) {
			return "testA", nil
		},
		"testB": func(arguments ...interface{}) (interface{}, error) {
			return "testB", nil
		},
	})
	if err != nil {
		t.Fatal("PARSE func error", err)
	}
	var tokens = eva.Tokens()
	if fn, ok := tokens[0].Value.(*Function); ok {
		if fn.Name != "testA" {
			t.Fatalf("get func name error want 'testA' got '%s'", fn.Name)
		}
	}

}
