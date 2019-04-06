package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var s []string

	for i := 0; i < 12; i++ {
		sc.Scan()
		s = append(s, sc.Text())
	}

	var count int
	for _, v := range s {
		if strings.Contains(v, "r") {
			count++
		}
	}

	fmt.Println(count)
}
