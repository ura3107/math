package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	input := make([]int, 3)
	for k, _ := range input {
		sc.Scan()
		var err error
		input[k], err = strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	var cnt int
	x := input[1]
	y := input[2]
	for n := input[0]; n > 0; n-- {
		if (n%x) == 0 || (n%y) == 0 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
