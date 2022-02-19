package bytesize_test

import (
	"math"
	"testing"

	"github.com/ambientkit/plugin/pkg/bytesize"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "1.23 B", bytesize.String(1.23, 2))
	assert.Equal(t, "1.3 B", bytesize.String(1.34, 1))
	assert.Equal(t, "1.4 B", bytesize.String(1.35, 1))
	assert.Equal(t, "1.2 B", bytesize.String(1.23, 1))
	assert.Equal(t, "1.00 KB", bytesize.String(math.Pow(1024, 1), 2))
	assert.Equal(t, "1.00 MB", bytesize.String(math.Pow(1024, 2), 2))
	assert.Equal(t, "1.00 GB", bytesize.String(math.Pow(1024, 3), 2))
	assert.Equal(t, "1.00 TB", bytesize.String(math.Pow(1024, 4), 2))
	assert.Equal(t, "1.00 PB", bytesize.String(math.Pow(1024, 5), 2))
	assert.Equal(t, "1.00 EB", bytesize.String(math.Pow(1024, 6), 2))
	assert.Equal(t, "1.00 ZB", bytesize.String(math.Pow(1024, 7), 2))
	assert.Equal(t, "1.00 YB", bytesize.String(math.Pow(1024, 8), 2))
	assert.Equal(t, "1024.00 YB", bytesize.String(math.Pow(1024, 9), 2))
}
