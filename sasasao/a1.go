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

	input := strings.Split(sc.Text(), "")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])
	d, _ := strconv.Atoi(input[3])

	if a == b && b == c && c == d {
		fmt.Println("SAME")
	} else {
		fmt.Println("DIFFERENT")
	}
}
