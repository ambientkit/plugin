package currency

import (
	"fmt"
	"strconv"
	"strings"
)

// Format numbers with commas and dollar signs. Only works below
// $1 million.
func Format(i int) string {
	s := strconv.Itoa(i)

	if i < 1000 {
		return fmt.Sprintf("$%v", s)
	} else if i < 10000 {
		return fmt.Sprintf("$%v,%v", s[0:1], s[1:])
	} else if i < 100000 {
		return fmt.Sprintf("$%v,%v", s[0:2], s[2:])
	} else if i < 1000000 {
		return fmt.Sprintf("$%v,%v", s[0:3], s[3:])
	}

	return fmt.Sprintf("$%v", s)
}

// FormatFloat converts a float into a currency with commas and a dollar sign.
func FormatFloat(i float64) string {
	f := fmt.Sprintf("%.2f", i)

	final := make([]string, 0)

	small := ""
	counter := 0

	for i := 0; i < len(f); i++ {
		char := string(f[len(f)-i-1])

		small = char + small

		if i < 3 {
			continue
		}

		counter++

		if counter < 3 {
			continue
		}

		final = append([]string{small}, final...)
		small = ""
		counter = 0
	}

	if len(small) > 0 {
		final = append([]string{small}, final...)
	}

	return "$" + strings.Join(final, ",")
}
