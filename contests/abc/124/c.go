package main

import (
	"bufio"
	"fmt"
	"os"
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
	s := strings.Split(readLine(), "")

	a := s[0]

	var sei []string

	if a == "0" {
		sei = append(sei, "0")
		for i := 1; i < len(s); i++ {
			if sei[i-1] == "0" {
				sei = append(sei, "1")
			} else {
				sei = append(sei, "0")
			}
		}
	} else {
		sei = append(sei, "1")
		for i := 1; i < len(s); i++ {
			if sei[i-1] == "0" {
				sei = append(sei, "1")
			} else {
				sei = append(sei, "0")
			}
		}
	}

	var count int

	for i := 0; i < len(s); i++ {
		if s[i] != sei[i] {
			count++
		}
	}

	fmt.Println(count)
}
