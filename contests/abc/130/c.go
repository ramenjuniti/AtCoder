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

func main() {
	ins := strings.Split(readLine(), " ")
	wi, _ := strconv.Atoi(ins[0])
	hi, _ := strconv.Atoi(ins[1])
	xi, _ := strconv.Atoi(ins[2])
	yi, _ := strconv.Atoi(ins[3])

	w := float64(wi)
	h := float64(hi)
	x := float64(xi)
	y := float64(yi)

	var j int

	if w/2 == x && h/2 == y {
		j = 1
	}

	fmt.Println(w*h/2, j)
}
