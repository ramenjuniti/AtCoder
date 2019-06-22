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

type cost struct {
	a int
	b int
}

type costs []cost

func (p costs) Len() int {
	return len(p)
}

func (p costs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p costs) Less(i, j int) bool {
	return p[i].b < p[j].b
}

func main() {
	ins := readLine()
	n, _ := strconv.Atoi(ins)

	c := make([]cost, n)
	for i := 0; i < n; i++ {
		inabs := strings.Split(readLine(), " ")
		ai, _ := strconv.Atoi(inabs[0])
		bi, _ := strconv.Atoi(inabs[1])
		c[i].a = ai
		c[i].b = bi
	}
	sort.Sort(costs(c))

	var sum int
	for i := range c {
		sum += c[i].a
		if sum > c[i].b {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
