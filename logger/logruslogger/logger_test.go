package logruslogger_test

import (
	"bufio"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/logger/logruslogger"
	"github.com/ambientkit/plugin/pkg/loggertestsuite"
)

// Run the standard logger test suite.
func TestMain(t *testing.T) {
	rt := loggertestsuite.New()
	rt.Run(t, func(writer *bufio.Writer) ambient.AppLogger {
		return logruslogger.NewLogger("app", "1.0", writer)
	})
}
