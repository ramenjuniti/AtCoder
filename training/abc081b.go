package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	nextLine()
	min := 0
	isFirst := true
	for _, v := range strings.Split(nextLine(), " ") {
		intV, _ := strconv.Atoi(v)
		count := 0
		if intV%2 != 0 {
			min = 0
			break
		}
		for intV%2 == 0 {
			intV /= 2
			count++
		}
		if isFirst {
			min = count
			isFirst = false
		} else if count < min {
			min = count
		}
	}
	fmt.Println(min)
}
