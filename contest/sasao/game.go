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
	d, _ := strconv.Atoi(input[1])

	if (a+1)*d > a*(d+1) {
		fmt.Println((a + 1) * d)
	} else {
		fmt.Println(a * (d + 1))
	}
}
