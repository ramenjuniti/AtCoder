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
	input := strings.Split(readLine(), " ")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	t, _ := strconv.Atoi(input[2])

	var ans, time, count int
	count++

	for {
		time = a * count
		if float64(time) > float64(t)+0.5 {
			break
		}
		ans += b
		count++
	}

	fmt.Println(ans)
}
