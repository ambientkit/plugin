package router

import (
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/routertestsuite"
)

// Run the standard router test suite.
func TestMain(t *testing.T) {
	rt := routertestsuite.New()
	rt.Run(t, func() ambient.AppRouter {
		return New()
	})
}
