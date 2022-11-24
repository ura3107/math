package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(2*n + 3)
}
