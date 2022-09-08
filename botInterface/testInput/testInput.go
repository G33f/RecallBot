package testInput

import "fmt"

const (
	WrongChoice = "wrong choice"
)

type testInput struct {
	userRequest     bool
	expectationList []string
}

type TestInput interface {
	TestInput(messageText string) error
	SetRequest(expectationList []string)
}

func New() TestInput {
	test := new(testInput)
	return test
}

func (t *testInput) TestInput(messageText string) error {
	testBase := t.expectationList
	for _, tests := range testBase {
		if messageText == tests {
			return nil
		}
	}
	return fmt.Errorf(WrongChoice)
}

func (t *testInput) SetRequest(expectationList []string) {
	t.userRequest = true
	t.expectationList = expectationList
}
