package zaplogger_test

import (
	"bufio"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/logger/zaplogger"
	"github.com/ambientkit/plugin/pkg/loggertestsuite"
)

// Run the standard logger test suite.
func TestMain(t *testing.T) {
	rt := loggertestsuite.New()
	rt.Run(t, func(writer *bufio.Writer) ambient.AppLogger {
		return zaplogger.NewLogger("app", "1.0", writer)
	})
}
