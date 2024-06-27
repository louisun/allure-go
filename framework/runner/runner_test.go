package runner

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/louisun/allure-go-v2/allure"
	"github.com/louisun/allure-go-v2/framework/core/common"
	"github.com/louisun/allure-go-v2/framework/core/constants"
	"github.com/louisun/allure-go-v2/framework/provider"
	"github.com/stretchr/testify/require"
)

type executionContextRunnerMock struct {
	name        string
	steps       []*allure.Step
	attachments []*allure.Attachment
}

func newExecContextRunnerMock(name string) *executionContextRunnerMock {
	return &executionContextRunnerMock{
		name:        name,
		steps:       []*allure.Step{},
		attachments: []*allure.Attachment{},
	}
}

func (m *executionContextRunnerMock) AddStep(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *executionContextRunnerMock) AddAttachments(attachments ...*allure.Attachment) {
	m.attachments = append(m.attachments, attachments...)
}

func (m *executionContextRunnerMock) GetName() string {
	return m.name
}

type providerMockRunner struct {
	provider.AllureForwardFull

	testMetaMock  provider.TestMeta
	suiteMetaMock *suiteMetaMockRunner
	executionMock *executionContextRunnerMock
}

func (m *providerMockRunner) GetResult() *allure.Result {
	return m.testMetaMock.GetResult()
}

func (m *providerMockRunner) UpdateResultStatus(msg string, trace string) {}

func (m *providerMockRunner) StopResult(status allure.Status) {}

func (m *providerMockRunner) GetTestMeta() provider.TestMeta {
	return m.testMetaMock
}

func (m *providerMockRunner) SetTestMeta(meta provider.TestMeta) {
	m.testMetaMock = meta
}

func (m *providerMockRunner) GetSuiteMeta() provider.SuiteMeta {
	return m.suiteMetaMock
}

func (m *providerMockRunner) ExecutionContext() provider.ExecutionContext {
	return m.executionMock
}

func (m *providerMockRunner) Step(step *allure.Step) {
	m.ExecutionContext().AddStep(step)
}

func (m *providerMockRunner) NewStep(stepName string, params ...*allure.Parameter) {
	m.ExecutionContext().AddStep(allure.NewSimpleStep(stepName, params...))
}

func (m *providerMockRunner) TestContext() {
	m.executionMock.name = constants.TestContextName
}

func (m *providerMockRunner) BeforeEachContext() {
	m.executionMock.name = constants.BeforeEachContextName
}

func (m *providerMockRunner) AfterEachContext() {
	m.executionMock.name = constants.AfterEachContextName
}

func (m *providerMockRunner) BeforeAllContext() {
	m.executionMock.name = constants.BeforeAllContextName
}

func (m *providerMockRunner) AfterAllContext() {
	m.executionMock.name = constants.AfterAllContextName
}

func (m *providerMockRunner) NewTest(testName, packageName string, tags ...string) {}
func (m *providerMockRunner) FinishTest() error {
	return nil
}

type suiteMetaMockRunner struct {
	namePrefix string
	name       string
	container  *allure.Container
	hookBa     func(t provider.T)
	hookAa     func(t provider.T)
}

func (m *suiteMetaMockRunner) GetPackageName() string {
	return m.name
}

func (m *suiteMetaMockRunner) GetRunner() string {
	return m.name
}

func (m *suiteMetaMockRunner) GetSuiteName() string {
	return m.name
}

func (m *suiteMetaMockRunner) GetParentSuite() string {
	return m.name
}

func (m *suiteMetaMockRunner) GetSuiteFullName() string {
	return fmt.Sprintf("%s/%s", m.namePrefix, m.name)
}

func (m *suiteMetaMockRunner) GetContainer() *allure.Container {
	return m.container
}

func (m *suiteMetaMockRunner) SetBeforeAll(hook func(provider.T)) {
	m.hookBa = hook
}

func (m *suiteMetaMockRunner) SetAfterAll(hook func(provider.T)) {
	m.hookAa = hook
}

func (m *suiteMetaMockRunner) GetBeforeAll() func(provider.T) {
	return m.hookBa
}

func (m *suiteMetaMockRunner) GetAfterAll() func(provider.T) {
	return m.hookAa
}

type testMetaMockRunner struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockRunner) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockRunner) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockRunner) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockRunner) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockRunner) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockRunner) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockRunner) GetAfterEach() func(t provider.T) {
	return m.ae
}

type runnerTMock struct {
	testing.TB

	t        *testing.T
	run      bool
	parallel bool
}

func (m *runnerTMock) Name() string {
	return "testName"
}

func (m *runnerTMock) Run(testName string, testBody func(t *testing.T)) bool {
	m.run = true
	testBody(m.t)
	return true
}

func (m *runnerTMock) Parallel() {

}

func newInternalTMock(name string) *common.Common {
	return &common.Common{
		TestingT: &runnerTMock{t: new(testing.T)},
		Provider: &providerMockRunner{
			executionMock: newExecContextRunnerMock(name),
			testMetaMock:  &testMetaMockRunner{container: allure.NewContainer()},
			suiteMetaMock: &suiteMetaMockRunner{container: allure.NewContainer()},
		},
	}
}

type iT interface {
	t() internalT
}

func TestNewRunner(t *testing.T) {
	result := NewRunner(t, "suiteTest")

	require.NotNil(t, result)
	it := result.(iT).t()
	require.NotNil(t, it)
	require.Equal(t, "TestNewRunner", it.Name())
	require.NotNil(t, it.RealT())
	require.Equal(t, t, it.RealT())

	provider := it.GetProvider()
	require.NotNil(t, provider)

	testMeta := provider.GetTestMeta()
	require.NotNil(t, testMeta)

	suiteMeta := provider.GetSuiteMeta()
	require.NotNil(t, suiteMeta)
	require.Equal(t, "suiteTest", suiteMeta.GetSuiteName())
	require.Equal(t, "TestNewRunner", suiteMeta.GetSuiteFullName())
	require.Equal(t, "TestNewRunner", suiteMeta.GetRunner())
	require.Equal(t, "github.com/louisun/allure-go-v2/framework/runner", suiteMeta.GetPackageName())
	require.NotNil(t, suiteMeta.GetContainer())
}

func TestRunner_BeforeEach_noStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int
	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.BeforeEachContextName)}

	meta := &testMetaMockRunner{result: &allure.Result{}, container: allure.NewContainer(), be: func(t provider.T) {
		counter++
		flag = true
	}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, 2, counter)
}

func TestRunner_BeforeEach_withStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.BeforeEachContextName)}
	r.BeforeEach(func(t provider.T) {
		t.NewStep("stepName")
		counter++
		flag = true
	})
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, counter, 2)
}

func TestRunner_AfterEach_noStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.AfterEachContextName)}
	r.AfterEach(func(t provider.T) {
		flag = true
		counter++
	})
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, counter, 2)
}

func TestRunner_AfterEach_withStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.AfterEachContextName)}

	meta := &testMetaMockRunner{result: &allure.Result{}, container: allure.NewContainer(), ae: func(t provider.T) {
		t.NewStep("stepName")
		flag = true
		counter++
	}}

	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, 2, counter)
}

func TestRunner_BeforeAll_noStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int
	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.BeforeAllContextName)}
	r.BeforeAll(func(t provider.T) {
		counter++
		flag = true
	})
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, 1, counter)
}

func TestRunner_BeforeAll_withStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.BeforeAllContextName)}
	r.BeforeAll(func(t provider.T) {
		t.NewStep("stepName")
		counter++
		flag = true
	})
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, 1, counter)
}

func TestRunner_AfterAll_noStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.AfterAllContextName)}
	r.AfterAll(func(t provider.T) {
		flag = true
		counter++
	})
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, 1, counter)
}

func TestRunner_AfterAll_withStep(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.AfterAllContextName)}
	r.AfterAll(func(t provider.T) {
		t.NewStep("stepName")
		flag = true
		counter++
	})
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) {}}

	r.RunTests()

	require.True(t, flag)
	require.Equal(t, counter, 1)
}

func TestRunner_RunTests(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.AfterAllContextName)}
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(t provider.T) { counter++ }}
	r.tests["test2"] = &testFunc{testMeta: meta, testBody: func(t provider.T) { counter++ }}

	r.RunTests()

	require.Equal(t, counter, 2)
}

func TestRunner_RunTestsPanic(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var counter int

	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)
	wg := sync.WaitGroup{}
	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.AfterAllContextName)}
	wg.Add(1)
	meta := &testMetaMockRunner{result: &allure.Result{}}
	r.tests["test"] = &testFunc{testMeta: meta, testBody: func(mockT provider.T) {
		defer wg.Done()
		counter++
		panic("whoops")
	}}

	wg.Add(1)
	go require.NotPanics(t, func() {
		defer wg.Done()
		r.RunTests()
	})
	wg.Wait()
	require.Equal(t, 1, counter)
}

func TestGetPackage(t *testing.T) {
	require.Equal(t, "github.com/louisun/allure-go-v2/framework/runner", getPackage(1))
}

func TestRunner_NewTest(t *testing.T) {
	t.Skipf("This test need to be reworked")
	var flag bool
	r := runner{tests: make(map[string]Test), internalT: newInternalTMock(constants.AfterAllContextName)}
	r.NewTest("TestName", func(t provider.T) {
		flag = true
	}, "tag1", "tag2")

	testKey := fmt.Sprintf("%s/%s", r.t().Name(), "TestName")
	tagList := r.tests[testKey].GetMeta().GetResult().GetLabels(allure.Tag)
	require.NotEmpty(t, r.tests)
	require.NotNil(t, r.tests[testKey])
	require.NotNil(t, r.tests[testKey].GetBody())
	require.NotEmpty(t, r.tests[testKey].GetMeta().GetResult().GetLabels(allure.Tag))
	require.Len(t, tagList, 2)
	require.Equal(t, "tag1", tagList[0])
	require.Equal(t, "tag2", tagList[1])
	require.Equal(t, "TestName", r.tests[testKey].GetMeta().GetResult().Name)

	r.tests[testKey].GetBody()(r.t())
	require.True(t, flag)
}
