package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	n, _ := strconv.Atoi(readLine())
	as := strings.Split(readLine(), " ")

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(as[i])
	}

	x := make([]int, n)
	var asum int
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			asum += a[i]
		} else {
			asum -= a[i]
		}
	}
	x[0] = asum / 2

	for i := 0; i < len(a)-1; i++ {
		x[i+1] = a[i] - x[i]
	}

	for i := 0; i < len(x); i++ {
		x[i] *= 2
	}

	for i, v := range x {
		if i == len(x)-1 {
			fmt.Println(v)
		} else {
			fmt.Print(fmt.Sprintf("%v ", v))
		}
	}
}
