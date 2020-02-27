package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int64
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{-1, -1, -2},
	}

	for _, test := range tests {
		got := add(test.a, test.b)
		if got != test.want {
			t.Errorf("comma(%q, %q), got %q, want %q", test.a, test.b, got, test.want)
		}
	}
}

func TestComma(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"1234561", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
		{"1234567890", "1,234,567,890"},
	}
	for _, test := range tests {
		got := comma(test.s)
		if got != test.want {
			t.Errorf("comma(%q), got %q, want %q", test.s, got, test.want)
		}
	}
}
