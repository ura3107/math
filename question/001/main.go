package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	fmt.Println("みかんの個数を入力してください。")
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()

	input = sc.Text()

	orangeNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("fail to convert int.")
		return
	}

	if orangeNum <= 1 && orangeNum >= 100 {
		fmt.Println("number of oranges is incorrect.")
		return
	}

	appleNum := 5

	total := appleNum + orangeNum

	fmt.Printf("%d\n", total)
}
