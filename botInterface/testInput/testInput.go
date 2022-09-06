package testInput

import "fmt"

const (
	TheUserIsNotOnTheRequestList = "it's not command"

	ThereAreNoExpectedRequests = "it's not command"

	WrongChoice = "wrong choice"
)

type testInput struct {
	userRequest     map[int64]bool
	expectationList map[int64][]string
}

type TestInput interface {
	TestInput(user int64, messageText string) error
	SetRequest(user int64, expectationList []string)
}

func New() TestInput {
	test := new(testInput)
	test.userRequest = make(map[int64]bool)
	test.expectationList = make(map[int64][]string)
	return test
}

func (t *testInput) TestInput(user int64, messageText string) error {
	test, err := t.userRequest[user]
	if !err {
		return fmt.Errorf(TheUserIsNotOnTheRequestList)
	}
	testBase, err := t.expectationList[user]
	if !err || !test {
		return fmt.Errorf(ThereAreNoExpectedRequests)
	}
	for _, tests := range testBase {
		if messageText == tests {
			delete(t.userRequest, user)
			delete(t.expectationList, user)
			return nil
		}
	}
	return fmt.Errorf(WrongChoice)
}

func (t *testInput) SetRequest(user int64, expectationList []string) {
	t.userRequest[user] = true
	t.expectationList[user] = expectationList
}
