package loggertestsuite

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/stretchr/testify/assert"
)

// TestSuite performs standard tests.
type TestSuite struct{}

// New returns a router test suite.
func New() *TestSuite {
	return new(TestSuite)
}

// Run all the tests.
func (ts *TestSuite) Run(t *testing.T, l func(writer *bufio.Writer) ambient.AppLogger) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	logger := l(writer)

	ts.TestLogLevel(t, logger, writer, &buffer)
	ts.TestDebug(t, logger, writer, &buffer)
	ts.TestInfo(t, logger, writer, &buffer)
	ts.TestWarn(t, logger, writer, &buffer)
	ts.TestError(t, logger, writer, &buffer)
	ts.TestFatal(t, logger, writer, &buffer)
	ts.TestOutputLevel(t, logger, writer, &buffer)
}

// TestLogLevel .
func (ts *TestSuite) TestLogLevel(t *testing.T, logger ambient.AppLogger, writer *bufio.Writer, buffer *bytes.Buffer) {
	logger.SetLogLevel(ambient.LogLevelDebug)
	writer.Flush()
	assert.Contains(t, buffer.String(), "debug")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelInfo)
	writer.Flush()
	assert.Contains(t, buffer.String(), "info")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelWarn)
	writer.Flush()
	assert.Contains(t, buffer.String(), "warn")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelError)
	writer.Flush()
	assert.Contains(t, buffer.String(), "error")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelFatal)
	writer.Flush()
	assert.Contains(t, buffer.String(), "fatal")
	buffer.Reset()
}

// TestDebug .
func (ts *TestSuite) TestDebug(t *testing.T, logger ambient.AppLogger, writer *bufio.Writer, buffer *bytes.Buffer) {
	logger.SetLogLevel(ambient.LogLevelDebug)
	writer.Flush()
	buffer.Reset()

	logger.Debug("test1 %v %v", "test2", "test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.Debug("test1 test2 test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelInfo)
	writer.Flush()
	buffer.Reset()
	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelWarn)
	writer.Flush()
	buffer.Reset()
	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelError)
	writer.Flush()
	buffer.Reset()
	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelFatal)
	writer.Flush()
	buffer.Reset()
	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()
}

// TestInfo .
func (ts *TestSuite) TestInfo(t *testing.T, logger ambient.AppLogger, writer *bufio.Writer, buffer *bytes.Buffer) {
	logger.SetLogLevel(ambient.LogLevelInfo)
	writer.Flush()
	buffer.Reset()

	logger.Info("test1 %v %v", "test2", "test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.Info("test1 test2 test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelWarn)
	writer.Flush()
	buffer.Reset()
	logger.Info("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelError)
	writer.Flush()
	buffer.Reset()
	logger.Info("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelFatal)
	writer.Flush()
	buffer.Reset()
	logger.Info("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()
}

// TestWarn .
func (ts *TestSuite) TestWarn(t *testing.T, logger ambient.AppLogger, writer *bufio.Writer, buffer *bytes.Buffer) {
	logger.SetLogLevel(ambient.LogLevelWarn)
	writer.Flush()
	buffer.Reset()

	logger.Warn("test1 %v %v", "test2", "test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.Warn("test1 test2 test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelError)
	writer.Flush()
	buffer.Reset()
	logger.Warn("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelFatal)
	writer.Flush()
	buffer.Reset()
	logger.Warn("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()
}

// TestError .
func (ts *TestSuite) TestError(t *testing.T, logger ambient.AppLogger, writer *bufio.Writer, buffer *bytes.Buffer) {
	logger.SetLogLevel(ambient.LogLevelError)
	writer.Flush()
	buffer.Reset()

	logger.Error("test1 %v %v", "test2", "test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.Error("test1 test2 test3")
	writer.Flush()
	assert.Contains(t, buffer.String(), "test1 test2 test3")
	buffer.Reset()

	logger.SetLogLevel(ambient.LogLevelFatal)
	writer.Flush()
	buffer.Reset()
	logger.Error("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()
}

// TestFatal .
func (ts *TestSuite) TestFatal(t *testing.T, logger ambient.AppLogger, writer *bufio.Writer, buffer *bytes.Buffer) {
	logger.SetLogLevel(ambient.LogLevelFatal)
	writer.Flush()
	buffer.Reset()
}

// TestOutputLevel .
func (ts *TestSuite) TestOutputLevel(t *testing.T, logger ambient.AppLogger, writer *bufio.Writer, buffer *bytes.Buffer) {
	logger.SetLogLevel(ambient.LogLevelDebug)
	writer.Flush()
	buffer.Reset()

	logger.Debug("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	logger.Info("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	logger.Warn("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	logger.Error("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	// *************************************************************************

	logger.SetLogLevel(ambient.LogLevelInfo)
	writer.Flush()
	buffer.Reset()

	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Info("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	logger.Warn("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	logger.Error("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	// *************************************************************************

	logger.SetLogLevel(ambient.LogLevelWarn)
	writer.Flush()
	buffer.Reset()

	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Info("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Warn("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	logger.Error("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	// *************************************************************************

	logger.SetLogLevel(ambient.LogLevelError)
	writer.Flush()
	buffer.Reset()

	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Info("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Warn("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Error("test")
	writer.Flush()
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()

	// *************************************************************************

	logger.SetLogLevel(ambient.LogLevelFatal)
	writer.Flush()
	buffer.Reset()

	logger.Debug("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Info("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Warn("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()

	logger.Error("test")
	writer.Flush()
	assert.Empty(t, buffer.String())
	buffer.Reset()
}
