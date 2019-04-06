package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var input []int
	for i := 0; i < 6; i++ {
		in, _ := strconv.Atoi(readLine())
		input = append(input, in)
	}
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if input[j]-input[i] > input[5] {
				fmt.Println(":(")
				return
			}
		}
	}
	fmt.Println("Yay!")
}
