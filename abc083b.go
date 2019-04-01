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

	input := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(input[0])
	a, _ := strconv.Atoi(input[1])
	b, _ := strconv.Atoi(input[2])

	sum := 0

	for i := 1; i <= n; i++ {
		j := i
		dsum := 0
		for j/10 != 0 {
			d := j % 10
			dsum += d
			j /= 10
		}
		dsum += j
		if a <= dsum && dsum <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}
