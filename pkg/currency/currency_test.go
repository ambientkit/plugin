package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatIntMinus(t *testing.T) {
	assert.Equal(t, "$1", FormatIntNegativeSign(1))
	assert.Equal(t, "$1,000", FormatIntNegativeSign(1000))
	assert.Equal(t, "-$1", FormatIntNegativeSign(-1))
}

func TestFormatIntParens(t *testing.T) {
	assert.Equal(t, "$1", FormatIntParens(1))
	assert.Equal(t, "$1,000", FormatIntParens(1000))
	assert.Equal(t, "($1)", FormatIntParens(-1))
	assert.Equal(t, "($1,000)", FormatIntParens(-1000))
}

func TestFormatFloat64(t *testing.T) {
	assert.Equal(t, "$1.00", FormatFloat64NegativeSign(1))
	assert.Equal(t, "$1,000.00", FormatFloat64NegativeSign(1000))
	assert.Equal(t, "-$1.00", FormatFloat64NegativeSign(-1))
	assert.Equal(t, "-$1,000.00", FormatFloat64NegativeSign(-1000))
}
