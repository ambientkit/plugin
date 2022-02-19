package bytesize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "1.23 B", ByteSize.String(1.23))
	assert.Equal(t, "1.00 KB", ByteSize.String(1024.56))
	assert.Equal(t, "1.00 MB", ByteSize.String(1024.56*1024))
}
