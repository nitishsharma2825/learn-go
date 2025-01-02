package word

import "testing"

func TestPalindrome(t *testing.T) {
	// Table driven testing
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrom(%q) = %v", test.input, got)
		}
	}
}
