package main

import (
	"fmt"
	"math"
	"unicode"
)

// 単項式から係数と次数を表示する
func main() {
	var input string
	fmt.Println("単項式の係数と次数を表示します。単項式を入力してください。\n冪乗は「^」で表現してください。\n例:5x^6y^7")
	fmt.Scan(&input)

	var str []rune

	for _, v := range input {
		str = append(str, v)
	}

	// マイナス
	minus := false
	// 係数
	var coefficientArr []int
	coefficientFlag := true
	// 冪乗
	var word string
	exponents := make(map[string][]int)

	// 文字列を一文字ずつ判定する
	for i, v := range str {
		// マイナスの有無
		if i == 0 && string(v) == "-" {
			minus = true
			continue
		}

		// 係数の取得
		if unicode.IsNumber(v) && coefficientFlag {
			coefficientArr = append([]int{int(v - '0')}, coefficientArr...)
			continue
		}

		// 変数(文字)の取得
		if unicode.IsLower(v) || unicode.IsUpper(v) {
			coefficientFlag = false
			word = string(v)
			// 変数かつ末尾の場合は、forを抜ける
			if len(str) == i+1 {
				exponents[word] = append(exponents[word], 1)
				break
			}
			continue
		}

		// 冪乗の取得
		if unicode.IsNumber(v) {
			exponents[word] = append([]int{int(v - '0')}, exponents[word]...)
		}
	}

	// 係数
	var coefficient int
	for i, c := range coefficientArr {
		digit := math.Pow(10, float64(i))
		coefficient = coefficient + (c * int(digit))
	}

	// 次数
	var degree int
	for _, m := range exponents {
		for i, e := range m {
			digit := math.Pow(10, float64(i))
			degree = degree + (e * int(digit))
		}
	}

	if coefficient == 0 {
		fmt.Printf("%sの係数は0、次数は定めません", input)
	} else {
		if minus {
			fmt.Printf("%sの係数は-%d、次数は%dです。", input, coefficient, degree)
		} else {
			fmt.Printf("%sの係数は%d、次数は%dです。", input, coefficient, degree)
		}
	}
}
