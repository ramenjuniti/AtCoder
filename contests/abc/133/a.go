package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

	n, _ := strconv.Atoi(ins[0])
	a, _ := strconv.Atoi(ins[1])
	b, _ := strconv.Atoi(ins[2])

	den := n * a
	tax := b
	if den > b {
		fmt.Println(tax)
	} else {
		fmt.Println(den)
	}
}
