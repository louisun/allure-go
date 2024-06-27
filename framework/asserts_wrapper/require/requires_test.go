package require

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/louisun/allure-go-v2/allure"
)

type providerTMock struct {
	steps        []*allure.Step
	errorF       bool
	errorFString string
	failNow      bool
}

func newMock() *providerTMock {
	return &providerTMock{steps: make([]*allure.Step, 0)}
}

func (p *providerTMock) Step(step *allure.Step) {
	p.steps = append(p.steps, step)
}

func (p *providerTMock) Errorf(format string, msgAndArgs ...interface{}) {
	p.errorFString = format
	p.errorF = true
}

func (p *providerTMock) FailNow() {
	p.failNow = true
}

func TestRequireExactly_Success(t *testing.T) {
	mockT := newMock()
	Exactly(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Exactly", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireExactly_Fail(t *testing.T) {
	mockT := newMock()
	Exactly(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Exactly", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireSame_Success(t *testing.T) {
	mockT := newMock()
	type someStr struct {
	}
	exp := &someStr{}
	act := exp
	Same(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Same", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireSame_Fail(t *testing.T) {
	mockT := newMock()
	type someStr struct {
		someField string
	}
	exp := &someStr{}
	act := &someStr{}

	Same(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Same", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotSame_Success(t *testing.T) {
	mockT := newMock()
	type someStr struct {
		someField string
	}
	exp := &someStr{}
	act := &someStr{}

	NotSame(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Same", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotSame_Fail(t *testing.T) {
	mockT := newMock()
	type someStr struct {
	}
	exp := &someStr{}
	act := exp
	NotSame(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Same", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEqual_Success(t *testing.T) {
	mockT := newMock()
	Equal(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEqual_Fail(t *testing.T) {
	mockT := newMock()
	Equal(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEqual_Success(t *testing.T) {
	mockT := newMock()
	NotEqual(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotEqual_Fail(t *testing.T) {
	mockT := newMock()
	NotEqual(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEqualValues_Success(t *testing.T) {
	mockT := newMock()
	EqualValues(mockT, uint32(123), int32(123))
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "uint32(0x7b)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "int32(123)", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEqualValues_Fail(t *testing.T) {
	mockT := newMock()
	EqualValues(mockT, 1, "test")
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "int(1)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "string(\"test\")", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEqualValues_Success(t *testing.T) {
	mockT := newMock()
	NotEqualValues(mockT, 1, "test")
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "int(1)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "string(\"test\")", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotEqualValues_Fail(t *testing.T) {
	mockT := newMock()
	NotEqualValues(mockT, uint32(123), int32(123))
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "uint32(0x7b)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "int32(123)", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireError_Success(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	Error(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireError_Fail(t *testing.T) {
	mockT := newMock()
	Error(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNoError_Success(t *testing.T) {
	mockT := newMock()
	NoError(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNoError_Fail(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NoError(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotNil_Success(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NotNil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEqualError_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := errors.New(exp)
	EqualError(mockT, err, exp)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, err.Error(), params[0].GetValue())
	require.Equal(t, "Expected", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEqualError_Fail(t *testing.T) {
	mockT := newMock()
	exp := "testErr2"
	actual := "testErr"
	err := errors.New(actual)
	EqualError(mockT, err, exp)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, err.Error(), params[0].GetValue())
	require.Equal(t, "Expected", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireErrorIs_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := fmt.Errorf(exp)
	errNew := errors.Wrap(err, "NewMessage")
	ErrorIs(mockT, errNew, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error Is", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, errNew.Error(), params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

type fakeError struct {
	input string
}

func (f *fakeError) Error() string {
	return fmt.Sprintf("fake error: %s", f.input)
}

func TestRequireErrorIs_Fail(t *testing.T) {
	mockT := newMock()

	var err = fakeError{"some"}
	errNew := errors.Wrap(fmt.Errorf("other"), "NewMessage")
	ErrorIs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error Is", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, "NewMessage: other", params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, "fake error: some", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireErrorAs_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := fmt.Errorf(exp)
	errNew := errors.Wrap(err, "NewMessage")
	ErrorAs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error As", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, errNew.Error(), params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireErrorAs_Fail(t *testing.T) {
	mockT := newMock()

	var err *fakeError
	errNew := errors.Wrap(fmt.Errorf("other"), "NewMessage")
	ErrorAs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error As", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, "NewMessage: other", params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, fmt.Sprintf("**require.fakeError((**require.fakeError)(%+v))", &err), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotNil_Failed(t *testing.T) {
	mockT := newMock()

	NotNil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNil_Success(t *testing.T) {
	mockT := newMock()

	Nil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNil_Failed(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	Nil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLen_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test\")", params[0].GetValue())
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireLen_Failed(t *testing.T) {
	mockT := newMock()
	str := "test1"

	Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test1\")", params[0].GetValue())
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NotContains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Not Contain", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	NotContains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Not Contain", params[1].Name)
	require.Equal(t, "est", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	Contains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Contain", params[1].Name)
	require.Equal(t, "est", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	Contains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Contain", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireGreater_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	Greater(mockT, test, 3)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "3", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireGreater_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	Greater(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireGreaterOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	GreaterOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireGreaterOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	GreaterOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLess_Success(t *testing.T) {
	mockT := newMock()
	test := 3

	Less(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "3", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireLess_Fail(t *testing.T) {
	mockT := newMock()
	test := 5

	Less(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "5", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLesOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	LessOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireLessOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 6

	LessOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "6", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

type testStructSuc struct {
}

func (t *testStructSuc) test() {
}

func TestRequireImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.testStructSuc(%#v)", ts), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireImplements_Failed(t *testing.T) {
	type testInterface interface {
		test2()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.testStructSuc(%#v)", ts), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := ""
	Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEmpty_False(t *testing.T) {
	mockT := newMock()

	test := "123"
	Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := "123"
	NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotEmpty_False(t *testing.T) {
	mockT := newMock()

	test := ""
	NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireWithDuration_Success(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Within Duration", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireWithDuration_Fail(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	test = test.Add(1000000)
	WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Within Duration", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireJSONEq_Success(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	JSONEq(mockT, exp, exp)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireJSONEq_Fail(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	JSONEq(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireJSONContains_Success(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 123, "key3": ["foo", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`
	JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireJSONContains_Fail(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 321, "key3": ["foobar", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireSubset_Success(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireSubset_Fail(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireIsType_Success(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	test := new(testStruct)

	IsType(mockT, test, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireIsType_Fail(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	type failTestStruct struct {
	}
	test := new(testStruct)
	act := new(failTestStruct)

	IsType(mockT, test, act)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testStruct(%#v)", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.failTestStruct(%#v)", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireTrue_Success(t *testing.T) {
	mockT := newMock()

	True(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireTrue_Fail(t *testing.T) {
	mockT := newMock()

	True(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireFalse_Success(t *testing.T) {
	mockT := newMock()

	False(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireFalse_Fail(t *testing.T) {
	mockT := newMock()

	False(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireRegexp_Success(t *testing.T) {
	mockT := newMock()

	rx := `^start`
	str := "start of the line"
	Regexp(mockT, rx, str)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Regexp", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, rx, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, str, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireRegexp_Failed(t *testing.T) {
	mockT := newMock()

	rx := `^end`
	str := "start of the line"
	Regexp(mockT, rx, str)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Regexp", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, rx, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, str, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireElementsMatch_Success(t *testing.T) {
	mockT := newMock()

	listA := []int{1, 2, 3}
	listB := []int{1, 2, 3}
	ElementsMatch(mockT, listA, listB)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Elements Match", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "ListA", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", listA), params[0].GetValue())

	require.Equal(t, "ListB", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", listB), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireElementsMatch_Fail(t *testing.T) {
	mockT := newMock()

	listA := []int{1, 2, 3}
	listB := []int{4, 3}
	ElementsMatch(mockT, listA, listB)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Elements Match", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "ListA", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", listA), params[0].GetValue())

	require.Equal(t, "ListB", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", listB), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireDirExists_Success(t *testing.T) {
	dirName := "test"
	err := os.Mkdir(dirName, 0644)
	require.NoError(t, err, "Can't create folder to begin test")
	defer os.RemoveAll(dirName)

	mockT := newMock()
	DirExists(mockT, dirName)
	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Dir Exists", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Path", params[0].Name)
	require.Equal(t, dirName, params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireDirExists_Fail(t *testing.T) {
	dirName := "test"

	mockT := newMock()
	DirExists(mockT, dirName)
	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Dir Exists", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Path", params[0].Name)
	require.Equal(t, dirName, params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireCondition_Success(t *testing.T) {
	test := false
	conditionFunc := func() bool {
		test = true
		return test
	}
	mockT := newMock()
	Condition(mockT, conditionFunc)
	steps := mockT.steps
	require.True(t, test)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Condition", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Signature", params[0].Name)
	require.Equal(t, "assert.Comparison", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireCondition_Fail(t *testing.T) {
	test := false
	conditionFunc := func() bool {
		test = true
		return !test
	}
	mockT := newMock()
	Condition(mockT, conditionFunc)
	steps := mockT.steps
	require.True(t, test)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Condition", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Signature", params[0].Name)
	require.Equal(t, "assert.Comparison", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireZero_Success(t *testing.T) {
	mockT := newMock()

	Zero(mockT, 0)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Zero", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "0", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireZero_Fail(t *testing.T) {
	mockT := newMock()

	Zero(mockT, 1)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Zero", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotZero_Success(t *testing.T) {
	mockT := newMock()

	NotZero(mockT, 1)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Zero", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotZero_Fail(t *testing.T) {
	mockT := newMock()

	NotZero(mockT, 0)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Zero", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "0", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}
