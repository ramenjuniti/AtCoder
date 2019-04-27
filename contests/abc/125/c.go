package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 10000)

func readLine() string {
	buf := make([]byte, 0, 10000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	_ = readLine()
	input := strings.Split(readLine(), " ")
	var a []int

	for i := 0; i < len(input); i++ {
		avalue, _ := strconv.Atoi(input[i])
		a = append(a, avalue)
	}

	sort.Sort(sort.IntSlice(a))

	c := 2
	var index int

L:
	for {
		for i := 0; i < len(a); i++ {
			if a[i]%c != 0 {
				index = i
				break L
			}
		}
		c++
	}

	for i := 0; i < len(a); i++ {
		if i != index {
			a[index] = a[i]
			break
		}
	}

	min := 1000000000

	for i := 1; i < len(a); i++ {
		if ans := gcd(a[0], a[i]); min > ans {
			min = ans
		}
	}

	fmt.Println(min)
}
