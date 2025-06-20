package a17countthenumberofvowelstringsinrange2586

import "testing"

func TestVowelStrings(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		left   int
		right  int
		expect int
	}{
		{
			name:   "basic vowels",
			words:  []string{"apple", "orange", "ice", "umbrella", "cat"},
			left:   0,
			right:  4,
			expect: 4,
		},
		{
			name:   "no vowels",
			words:  []string{"cat", "dog", "bat"},
			left:   0,
			right:  2,
			expect: 0,
		},
		{
			name:   "single element vowel",
			words:  []string{"apple"},
			left:   0,
			right:  0,
			expect: 1,
		},
		{
			name:   "single element non-vowel",
			words:  []string{"cat"},
			left:   0,
			right:  0,
			expect: 0,
		},
		{
			name:   "partial range",
			words:  []string{"apple", "orange", "ice", "umbrella", "cat"},
			left:   1,
			right:  3,
			expect: 3,
		},
		{
			name:   "mixed case",
			words:  []string{"Apple", "orange", "Ice", "umbrella", "cat"},
			left:   0,
			right:  4,
			expect: 2, // only "orange", "umbrella" (case sensitive)
		},
		{
			name:   "empty words",
			words:  []string{},
			left:   0,
			right:  -1,
			expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vowelStrings3(tt.words, tt.left, tt.right)
			if got != tt.expect {
				t.Errorf("vowelStrings(%v, %d, %d) = %d; want %d", tt.words, tt.left, tt.right, got, tt.expect)
			}
		})
	}
}
