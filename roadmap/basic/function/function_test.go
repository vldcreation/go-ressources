package function_test

import (
	"testing"

	"github.com/vldcration/go-ressources/roadmap/basic/function"
)

func TestSimpleFunction(t *testing.T) {
	ts := function.NewTestSuite("TestSuite").WithInvoker("TestSimpleFunction")
	// Call SimpleFunction from functions package
	ts.Prepare("Call Simple Function from functions package").SetSimpleFunc(function.SimpleFunction).Run()

	// Call SimpleFunctionWithArguments from functions package
	ts.Prepare("Call Simple Function with arguments from functions package").SetSimpleFuncWithArguments(func(a ...any) {
		if len(a) == 2 {
			function.SimpleFunctionWithArguments(a[0].(int), a[1].(int))
		}
	}).Run()

	// Call SimpleFunctionWithReturn from functions package
	result := function.SimpleFunctionWithReturn()
	if result != function.SimpleFunctionWithReturnResult {
		t.Errorf("Expected %d, got %d", function.SimpleFunctionWithReturnResult, result)
	}

	// Call SimpleFunctionWithArgumentsAndReturn from functions package
	result = function.SimpleFunctionWithArgumentsAndReturn(10, 20)
	if result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}

}

func TestMultipleReturnValues(t *testing.T) {
	ts := function.NewTestSuite("TestSuite").WithInvoker("TestMultipleReturnValues")
	// Call MultipleReturnValues from functions package
	ts.Prepare("Call Multiple Return Values from functions package").SetSimpleFunc(func() {
		a, b := function.MultipleReturnValues()
		if a != 10 || b != 20 {
			t.Errorf("Expected 10 and 20, got %d and %d", a, b)
		}
	}).Run()
}
