package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for {
		if n%a == 0 && n%b == 0 {
			fmt.Println(n)
			break
		}
		n++
	}
}
