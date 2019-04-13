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
	inputH := strings.Split(readLine(), " ")

	var max, count int

	max, _ = strconv.Atoi(inputH[0])
	count = 1

	for i := 1; i < n; i++ {
		h, _ := strconv.Atoi(inputH[i])
		if h >= max {
			count++
			max = h
		}
	}

	fmt.Println(count)
}
