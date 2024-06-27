package helper

import (
	"github.com/louisun/allure-go-v2/framework/asserts_wrapper/wrapper"
)

// NewRequireHelper inits new Require interface
func NewRequireHelper(t ProviderT) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewRequire(t),
	}
}
