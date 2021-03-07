package main

import "math"

func mySqrt(x int) int {
	//i := 0
	//for i*i <= x {
	//	i++
	//}
	//return i - 1
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans + 1) * (ans + 1) <= x {
		return ans + 1
	}
	return ans

}
