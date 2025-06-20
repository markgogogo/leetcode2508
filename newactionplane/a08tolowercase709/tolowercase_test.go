package a08tolowercase709

import "testing"

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "hello"},
		{"WORLD", "world"},
		{"GoLang", "golang"},
		{"12345", "12345"},
		{"", ""},
		{"MiXeD123", "mixed123"},
	}

	for _, test := range tests {
		result := toLowerCase2(test.input)
		if result != test.expected {
			t.Errorf("toLowerCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
