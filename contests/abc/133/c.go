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
	in := readLine()
	ins := strings.Split(in, " ")

	l, _ := strconv.Atoi(ins[0])
	r, _ := strconv.Atoi(ins[1])

	if r-l >= 2019 {
		fmt.Println(0)
		return
	}

	min := 2019
	for i := l; i <= r-1; i++ {
		for j := i + 1; j <= r; j++ {
			if a := (i * j) % 2019; min > a {
				min = a
			}
		}
	}

	fmt.Println(min)
}
