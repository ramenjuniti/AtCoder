package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	readLine()
	ins := strings.Split(readLine(), " ")

	n := len(ins)
	var d []int
	for _, v := range ins {
		di, _ := strconv.Atoi(v)
		d = append(d, di)
	}
	sort.Ints(d)

	var count int
	if d[(n-1)/2] == d[(n-1)/2+1] {
		fmt.Println(count)
		return
	}

	fmt.Println(d[(n-1)/2+1] - d[(n-1)/2])
}
