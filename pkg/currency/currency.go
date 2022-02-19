package currency

import (
	"fmt"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// FormatIntNegativeSign returns a string with dollar sign, commas, and a minus sign if
// negative.
func FormatIntNegativeSign(i int) string {
	return negativeFormat(i, NegativeSign)
}

// FormatIntParens returns a string with dollar sign, commas, and in parenthesis
// if negative.
func FormatIntParens(i int) string {
	return negativeFormat(i, NegativeParen)
}

// FormatFloat64NegativeSign reurns a string with dollar sign, commas, and in
// parenthesis if negative.
func FormatFloat64NegativeSign(i float64) string {
	return negativeFormat(i, NegativeSign)
}

// FormatFloat64Parens reurns a string with dollar sign, commas, and a minus
// sign if negative.
func FormatFloat64Parens(i float64) string {
	return negativeFormat(i, NegativeParen)
}

// Negative is format of showing a negative currency.
type Negative int64

const (
	// NegativeSign shows currency with a minus sign.
	NegativeSign Negative = iota
	// NegativeParen shows currency in parenthesis.
	NegativeParen
)

func negativeFormat(i interface{}, negativeType Negative) string {
	p := message.NewPrinter(language.English)

	raw := ""
	negative := false

	switch v := i.(type) {
	case int:
		var abs int
		abs, negative = absInt(v)
		raw = p.Sprintf("$%d", abs)
	case float64:
		var abs float64
		abs, negative = absFloat64(v)
		raw = p.Sprintf("$%.2f", abs)
	}

	if negative {
		switch negativeType {
		case NegativeSign:
			return fmt.Sprintf("-%s", raw)
		case NegativeParen:
			return fmt.Sprintf("(%s)", raw)
		}
	}

	return raw
}

// absInt returns absolute value for int and whether it was negative or not.
func absInt(i int) (absInt int, negative bool) {
	if i < 0 {
		return -i, true
	}
	return i, false
}

// absFloat64 returns absolute value for a float and whether it was negative or not.
func absFloat64(i float64) (absFloat64 float64, negative bool) {
	abs := math.Abs(i)
	if abs != i {
		return abs, true
	}
	return i, false
}
