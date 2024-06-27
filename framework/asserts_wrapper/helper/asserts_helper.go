package helper

import "github.com/louisun/allure-go-v2/framework/asserts_wrapper/wrapper"

// NewAssertsHelper inits new Assert interface
func NewAssertsHelper(t ProviderT) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewAsserts(t),
	}
}
