package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a, _ := strconv.Atoi(nextLine())
	b, _ := strconv.Atoi(nextLine())
	c, _ := strconv.Atoi(nextLine())
	x, _ := strconv.Atoi(nextLine())
	count := 0

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
