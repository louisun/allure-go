package ctx

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/louisun/allure-go-v2/allure"
	"github.com/louisun/allure-go-v2/framework/core/constants"
)

func TestNewAfterAllCtx(t *testing.T) {
	ctx := NewAfterAllCtx(allure.NewContainer())
	require.NotNil(t, ctx)
}

func TestNewAfterEachCtx(t *testing.T) {
	ctx := NewAfterEachCtx(allure.NewContainer())
	require.NotNil(t, ctx)
}

func TestNewBeforeEachCtx(t *testing.T) {
	ctx := NewBeforeEachCtx(allure.NewContainer())
	require.NotNil(t, ctx)
}

func TestNewBeforeAllCtx(t *testing.T) {
	ctx := NewBeforeAllCtx(allure.NewContainer())
	require.NotNil(t, ctx)
}

func TestHooksCtx_GetName(t *testing.T) {
	th := hooksCtx{name: "test"}
	require.Equal(t, "test", th.GetName())
}

func TestHooksCtx_AddStep(t *testing.T) {
	testStep := allure.NewSimpleStep("test")
	beforeEach := hooksCtx{name: constants.BeforeEachContextName, container: &allure.Container{}}
	beforeEach.AddStep(testStep)
	require.NotEmpty(t, beforeEach.container.Befores)
	require.Len(t, beforeEach.container.Befores, 1)
	require.Equal(t, testStep, beforeEach.container.Befores[0])

	beforeAll := hooksCtx{name: constants.BeforeAllContextName, container: &allure.Container{}}
	beforeAll.AddStep(testStep)
	require.NotEmpty(t, beforeAll.container.Befores)
	require.Len(t, beforeAll.container.Befores, 1)
	require.Equal(t, testStep, beforeAll.container.Befores[0])

	afterEach := hooksCtx{name: constants.AfterEachContextName, container: &allure.Container{}}
	afterEach.AddStep(testStep)
	require.NotEmpty(t, afterEach.container.Afters)
	require.Len(t, afterEach.container.Afters, 1)
	require.Equal(t, testStep, afterEach.container.Afters[0])

	afterAll := hooksCtx{name: constants.AfterAllContextName, container: &allure.Container{}}
	afterAll.AddStep(testStep)
	require.NotEmpty(t, afterAll.container.Afters)
	require.Len(t, afterAll.container.Afters, 1)
	require.Equal(t, testStep, afterAll.container.Afters[0])
}

func TestHooksCtx_AddAttachment(t *testing.T) {
	attach := allure.NewAttachment("testAttach", allure.Text, []byte("test"))
	beforeAll := hooksCtx{name: constants.BeforeAllContextName, container: &allure.Container{}}
	beforeAll.AddAttachments(attach)
	require.NotEmpty(t, beforeAll.container.Befores)
	require.Len(t, beforeAll.container.Befores, 1)
	require.NotEmpty(t, beforeAll.container.Befores[0].Attachments)
	require.Len(t, beforeAll.container.Befores[0].Attachments, 1)
	require.Equal(t, attach, beforeAll.container.Befores[0].Attachments[0])

	beforeEach := hooksCtx{name: constants.BeforeEachContextName, container: &allure.Container{}}
	beforeEach.AddAttachments(attach)
	require.NotEmpty(t, beforeEach.container.Befores)
	require.Len(t, beforeEach.container.Befores, 1)
	require.NotEmpty(t, beforeEach.container.Befores[0].Attachments)
	require.Len(t, beforeEach.container.Befores[0].Attachments, 1)
	require.Equal(t, attach, beforeEach.container.Befores[0].Attachments[0])

	afterAll := hooksCtx{name: constants.AfterAllContextName, container: &allure.Container{}}
	afterAll.AddAttachments(attach)
	require.NotEmpty(t, afterAll.container.Afters)
	require.Len(t, afterAll.container.Afters, 1)
	require.NotEmpty(t, afterAll.container.Afters[0].Attachments)
	require.Len(t, afterAll.container.Afters[0].Attachments, 1)
	require.Equal(t, attach, afterAll.container.Afters[0].Attachments[0])

	afterEach := hooksCtx{name: constants.AfterEachContextName, container: &allure.Container{}}
	afterEach.AddAttachments(attach)
	require.NotEmpty(t, afterEach.container.Afters)
	require.Len(t, afterEach.container.Afters, 1)
	require.NotEmpty(t, afterEach.container.Afters[0].Attachments)
	require.Len(t, afterEach.container.Afters[0].Attachments, 1)
	require.Equal(t, attach, afterEach.container.Afters[0].Attachments[0])
}
