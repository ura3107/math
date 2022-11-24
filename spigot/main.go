package main

import "fmt"

func main() {
	var numerator [8400]int // 分子
	base := 10000           // 基底
	n := 8400               // 計算項数

	for i := 0; i < n; i++ {
		numerator[i] = base / 5
	}

	var denom int
	out := 0

	for n = 8400; n > 0; n -= 14 {
		temp := 0
		for i := n - 1; i > 0; i-- {
			denom = 2*i - 1
			temp = temp*i + numerator[i]*base
			numerator[i] = temp % denom
			temp /= denom
		}
		fmt.Printf("%04d", out+temp/base)
		out = temp % base
	}
}
