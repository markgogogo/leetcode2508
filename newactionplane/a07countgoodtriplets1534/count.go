package a07countgoodtriplets1534

import "math"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	for i := 0; i+2 < len(arr); i++ {
		for j := i + 1; j+1 < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if math.Abs(float64(arr[i]-arr[j])) <= float64(a) &&
					math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
					math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
					res++
				}
			}
		}
	}
	return res
}
