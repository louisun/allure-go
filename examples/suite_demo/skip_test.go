//go:build examples_new
// +build examples_new

package suite_demo

import (
	"testing"

	"github.com/louisun/allure-go-v2/framework/provider"
	"github.com/louisun/allure-go-v2/framework/suite"
)

type SkipDemoSuite struct {
	suite.Suite
}

func (s *SkipDemoSuite) TestSkip(t provider.T) {
	t.Epic("Demo")
	t.Feature("Skip Test")
	t.Title("Skip test")
	t.Description(`
		This test will be skipped`)

	t.Tags("Test", "Skip")
	t.Skip("Skip Reason")
}

func TestSkipDemo(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(SkipDemoSuite))
}
