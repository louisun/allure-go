package runner

import (
	"sync"
	"testing"

	"github.com/louisun/allure-go-v2/allure"
	"github.com/louisun/allure-go-v2/framework/provider"
)

// AllureBeforeTest has a BeforeEach method, which will run before each
// test in the suite.
type AllureBeforeTest interface {
	BeforeEach(t provider.T)
}

// AllureAfterTest has a AfterEach method, which will run after
// each test in the suite.
type AllureAfterTest interface {
	AfterEach(t provider.T)
}

// AllureBeforeSuite has a BeforeAll method, which will run before the
// tests in the suite are run.
type AllureBeforeSuite interface {
	BeforeAll(t provider.T)
}

// AllureAfterSuite has a AfterAll method, which will run after
// all the tests in the suite have been run.
type AllureAfterSuite interface {
	AfterAll(t provider.T)
}

// WithTestPramsSuite has an InitTestParams method, which will run before
// collecting the tests in the suite.
type WithTestPramsSuite interface {
	InitTestParams()
}

type TestSuite interface {
	GetRunner() TestRunner
	SetRunner(runner TestRunner)
	AddAllureIDMapping(testName, allureID string)
	FindAllureID(testName string) (id string, ok bool)
}

type TestingT interface {
	testing.TB
	Parallel()
	Run(testName string, testBody func(t *testing.T)) bool
}

type TestRunner interface {
	NewTest(testName string, testBody func(provider.T), tags ...string)
	BeforeEach(hookBody func(provider.T))
	AfterEach(hookBody func(provider.T))
	BeforeAll(hookBody func(provider.T))
	AfterAll(hookBody func(provider.T))
	RunTests() SuiteResult
}

type Test interface {
	GetBody() TestBody
	GetMeta() provider.TestMeta
}

type SuiteResult interface {
	NewResult(result TestResult)
	GetContainer() *allure.Container
	GetAllTestResults() []TestResult
	GetResultByName(name string) TestResult
	GetResultByUUID(uuid string) TestResult
	ToJSON() ([]byte, error)
}

type TestResult interface {
	GetResult() *allure.Result
	GetContainer() *allure.Container
	Print() error
	ToJSON() ([]byte, error)
}

type internalT interface {
	provider.T

	SetRealT(t provider.TestingT)
	GetProvider() provider.Provider
	WG() *sync.WaitGroup
	GetResult() *allure.Result
}
