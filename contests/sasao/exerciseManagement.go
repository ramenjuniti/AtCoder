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
	l, _ := strconv.Atoi(input[0])
	h, _ := strconv.Atoi(input[1])
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var a []int
	for i := 0; i < n; i++ {
		sc.Scan()
		in, _ := strconv.Atoi(sc.Text())
		if in > h {
			a = append(a, -1)
		} else if in < l {
			a = append(a, l-in)
		} else {
			a = append(a, 0)
		}
	}

	for _, v := range a {
		fmt.Printf("%v\n", v)
	}
}
