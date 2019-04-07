package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

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
	s := readLine()

	fmt.Println(math.Min(float64(strings.Count(s, "1")), float64(strings.Count(s, "0"))) * 2)
}
