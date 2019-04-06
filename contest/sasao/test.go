package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	input := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])

	if a < b {
		fmt.Println("Better")
	} else {
		fmt.Println("Worse")
	}
}
