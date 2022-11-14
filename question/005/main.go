package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	inputs := strings.Split(sc.Text(), " ")
	var sum int
	for _, i := range inputs {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("fail to convert int")
			return
		}
		sum = sum + num
	}

	fmt.Printf("%d\n", sum%100)
}
