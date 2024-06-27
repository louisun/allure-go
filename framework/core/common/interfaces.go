package common

import (
	"sync"

	"github.com/louisun/allure-go-v2/allure"
	"github.com/louisun/allure-go-v2/framework/provider"
)

type ParentT interface {
	GetProvider() provider.Provider
	GetResult() *allure.Result
}

type HookProvider interface {
	BeforeEachContext()
	AfterEachContext()
	BeforeAllContext()
	AfterAllContext()

	GetSuiteMeta() provider.SuiteMeta
	GetTestMeta() provider.TestMeta
}

type InternalT interface {
	provider.T
	SetRealT(realT provider.TestingT)
	WG() *sync.WaitGroup
}
