package helper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/louisun/allure-go-v2/allure"
)

type tRequireMock struct {
}

func (p *tRequireMock) Step(step *allure.Step) {
}

func (p *tRequireMock) Errorf(format string, msgAndArgs ...interface{}) {
}

func (p *tRequireMock) FailNow() {
}

func TestNewRequireHelper(t *testing.T) {
	h := NewAssertsHelper(&tRequireMock{})
	require.NotNil(t, h)
}
