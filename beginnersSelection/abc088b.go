package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n, _ := strconv.Atoi(nextLine())
	inputStr := strings.Split(nextLine(), " ")
	inputInt := make([]int, n)
	a, b := 0, 0

	for i := range inputStr {
		inputInt[i], _ = strconv.Atoi(inputStr[i])
	}
	sort.Ints(inputInt)

	j := 0
	for i := n - 1; i >= 0; i-- {
		if j%2 == 0 {
			a += inputInt[i]
		} else {
			b += inputInt[i]
		}
		j++
	}
	fmt.Println(a - b)
}
