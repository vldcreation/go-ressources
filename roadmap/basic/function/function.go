package function

import "fmt"

type TestSuite struct {
	Invoker                 string
	name                    string
	keyword                 string
	simpleFunc              func()
	simpleFuncWithArguments func(...any)
}

func NewTestSuite(name string) *TestSuite {
	return &TestSuite{
		name: name,
	}
}

func (ts *TestSuite) Prepare(keyword string) *TestSuite {
	ts.keyword = keyword
	return ts
}

func (ts *TestSuite) WithInvoker(invoker string) *TestSuite {
	ts.Invoker = invoker
	if ts.Invoker == "" {
		ts.Invoker = "unknown"
	}
	return ts
}

func (ts *TestSuite) SetSimpleFunc(f func()) *TestSuite {
	ts.simpleFunc = f

	return ts
}

func (ts *TestSuite) SetSimpleFuncWithArguments(f func(a ...any)) *TestSuite {
	ts.simpleFuncWithArguments = f

	return ts
}

func (ts *TestSuite) Run() {
	if ts.simpleFunc != nil {
		if ts.keyword == "" {
			fmt.Printf("Running %s test suite with invoker %s\n", ts.name, ts.Invoker)
		} else {
			fmt.Printf("%s: \n", ts.keyword)
		}

		ts.simpleFunc()
	}

	if ts.simpleFuncWithArguments != nil {
		if ts.keyword == "" {
			fmt.Printf("Running %s test suite with invoker %s\n", ts.name, ts.Invoker)
		} else {
			fmt.Printf("%s: \n", ts.keyword)
		}

		ts.simpleFuncWithArguments()
	}
}
