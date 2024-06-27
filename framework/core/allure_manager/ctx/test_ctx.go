package ctx

import (
	"github.com/louisun/allure-go-v2/allure"
	"github.com/louisun/allure-go-v2/framework/core/constants"
	"github.com/louisun/allure-go-v2/framework/provider"
)

type testCtx struct {
	name   string
	result *allure.Result
}

func NewTestCtx(result *allure.Result) provider.ExecutionContext {
	return &testCtx{result: result, name: constants.TestContextName}
}

func (ctx *testCtx) AddStep(newStep *allure.Step) {
	ctx.result.Steps = append(ctx.result.Steps, newStep)
}

func (ctx *testCtx) GetName() string {
	return ctx.name
}

func (ctx *testCtx) AddAttachments(attachments ...*allure.Attachment) {
	ctx.result.Attachments = append(ctx.result.Attachments, attachments...)
}
