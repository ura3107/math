package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := make([]int, 3)
	fmt.Println("A1,A2,A3の値を入力してください。")
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()

	var err error
	if input[0], err = strconv.Atoi(sc.Text()); err != nil {
		fmt.Println("fail to convert int.")
		return
	}

	sc.Scan()
	if input[1], err = strconv.Atoi(sc.Text()); err != nil {
		fmt.Println("fail to convert int.")
		return
	}

	sc.Scan()
	if input[2], err = strconv.Atoi(sc.Text()); err != nil {
		fmt.Println("fail to convert int.")
		return
	}

	var total int
	for _, n := range input {
		total = total + n
	}

	fmt.Printf("%d\n", total)
}
