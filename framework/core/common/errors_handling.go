package common

import (
	"github.com/louisun/allure-go-v2/allure"
	"github.com/louisun/allure-go-v2/framework/core/constants"
)

type ErrorT interface {
	Errorf(format string, args ...interface{})
	Logf(format string, args ...interface{})
	FailNow()
}

type ErrorProvider interface {
	StopResult(status allure.Status)
	UpdateResultStatus(msg string, trace string)
}

func TestError(t ErrorT, provider ErrorProvider, contextName, errMsg string) {
	short := errMsg
	if len(errMsg) > 100 {
		short = errMsg[:100]
	}
	switch contextName {
	case constants.TestContextName, constants.BeforeEachContextName:
		provider.StopResult(allure.Broken)
		provider.UpdateResultStatus(short, errMsg)
		t.Errorf(errMsg)
		t.FailNow()
	case constants.AfterEachContextName, constants.AfterAllContextName:
		t.Logf(errMsg)
		provider.UpdateResultStatus(short, errMsg)
	case constants.BeforeAllContextName:
		t.Logf(errMsg)
		provider.UpdateResultStatus(short, errMsg)
		t.FailNow()
	}
}
