package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(n, y, man, gosen, sen int) (int, int, int) {
	for i := 0; i <= man; i++ {
		for j := 0; j <= gosen; j++ {
			if sen-(i*10+j*5) >= 0 && i+j+sen-(i*10+j*5) == n && i*10000+j*5000+(sen-(i*10+j*5))*1000 == y {
				return i, j, sen - (i*10 + j*5)
			}
		}
	}
	return -1, -1, -1
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	input := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(input[0])
	y, _ := strconv.Atoi(input[1])
	man := y / 10000
	gosen := y / 5000
	sen := y / 1000
	manCount, gosenCount, senCount := solve(n, y, man, gosen, sen)

	fmt.Printf("%d %d %d\n", manCount, gosenCount, senCount)
}
