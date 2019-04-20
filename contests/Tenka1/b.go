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
	_ = readLine()
	s := strings.Split(readLine(), "")
	k, _ := strconv.Atoi(readLine())

	match := s[k-1]

	for i := 0; i < len(s); i++ {
		if s[i] != match {
			fmt.Print("*")
		} else {
			fmt.Print(s[i])
		}
	}

	fmt.Println()
}
