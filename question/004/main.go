package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var answer int = 1
	for i := 0; i < 3; i++ {
		sc.Scan()

		input, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println("fail to covert int")
			return
		}
		answer = answer * input
	}

	fmt.Printf("%d\n", answer)
}
