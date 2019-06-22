package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
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

func gojo(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b < 1 {
		return -1
	}

	if a%b == 0 {
		return b
	}

	return gojo(b, a%b)
}

func main() {
	ins := strings.Split(readLine(), " ")

	a, _ := strconv.Atoi(ins[0])
	b, _ := strconv.Atoi(ins[1])
	c, _ := strconv.Atoi(ins[2])
	d, _ := strconv.Atoi(ins[3])

	a--

	ca := a / c
	da := a / d
	cda := a / (c * d / gojo(c, d))

	cb := b / c
	db := b / d
	cdb := b / (c * d / gojo(c, d))

	l := b - (cb + db) + cdb
	r := a - (ca + da) + cda

	fmt.Println(l - r)
}
