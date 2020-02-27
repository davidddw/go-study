package calc

import (
	"bytes"
)

func add(a, b int64) int64 {
	return a + b
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	b := &bytes.Buffer{}
	pre := len(s) % 3
	// Write the first group of up to 3 digits.
	if pre == 0 {
		pre = 3
	}
	b.WriteString(s[:pre])
	// Deal with the rest.
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
