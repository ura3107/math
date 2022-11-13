package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputN int
	fmt.Println("Nの値を入力してください。")
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()

	var err error
	if inputN, err = strconv.Atoi(sc.Text()); err != nil {
		fmt.Println("fail to convert int.")
		return
	}

	inputA := make([]int, inputN)
	for i := 1; i <= inputN; i++ {
		var sc = bufio.NewScanner(os.Stdin)
		sc.Scan()
		if inputA[i-1], err = strconv.Atoi(sc.Text()); err != nil {
			fmt.Println("fail to convert int.")
			return
		}
	}

	var total int
	for _, n := range inputA {
		total = total + n
	}

	fmt.Printf("%d\n", total)
}
