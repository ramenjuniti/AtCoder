package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n, _ := strconv.Atoi(nextLine())
	var riceCake, count int
	var riceCakes []int

	for i := 0; i < n; i++ {
		riceCake, _ = strconv.Atoi(nextLine())
		riceCakes = append(riceCakes, riceCake)
	}
	sort.Sort(sort.IntSlice(riceCakes))

	for i := 0; i < n; i++ {
		if i == 0 {
			count++
		} else if riceCakes[i] > riceCakes[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
