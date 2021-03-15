package main

func rand10() int {
	a, b := rand7(), rand7()
	for a == 7 {
		// 1,2,3,4,5,6
		// a有一般概率为奇数 一半概率为偶数
		// 1/2
		a = rand7()
	}
	for b > 5 {
		// 1/2/3/4/5
		// 所以b是1-5
		// 1/5
		b = rand7()
	}
	// 1/2 * 1/5 = 1/10
	// 判端a的奇偶 奇数取0 偶数取5 概率1/2
	if a&1 == 1 {
		return b
	}
	return b + 5
}
