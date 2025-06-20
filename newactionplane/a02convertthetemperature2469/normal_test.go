package a02convertthetemperature2469

import (
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		celsius  float64
		expected []float64
	}{
		{0, []float64{273.15, 32}},
		{100, []float64{373.15, 212}},
		{-40, []float64{233.15, -40}},
		{37, []float64{310.15, 98.6}},
	}

	for _, test := range tests { // TODO：浮点数相等校验
		result := convertTemperature(test.celsius)
		for i := 0; i < len(result); i++ {
			if result[i] != test.expected[i] {
				t.Errorf("convertTemperature(%f) = %v; expected %v", test.celsius, result, test.expected)
			}
		}
		// if !reflect.DeepEqual(result, test.expected) {
		// 	t.Errorf("convertTemperature(%f) = %v; expected %v", test.celsius, result, test.expected)
		// }
	}
}
