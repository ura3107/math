package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("10進数の値を入力してください。")
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	input, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println("fail to convert int")
		return
	}

	var answer string
	for input >= 1 {
		if (input % 2) == 0 {
			answer = "0" + answer
		}
		if (input % 2) == 1 {
			answer = "1" + answer
		}
		input = input / 2
	}

	fmt.Printf("%s\n", answer)
}
