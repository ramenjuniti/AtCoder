package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
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
	d, _ := strconv.Atoi(ins[1])

	var x [][]int

	for i := 0; i < n; i++ {
		xs := strings.Split(readLine(), " ")
		var xi []int
		for _, v := range xs{
			vint, _ := strconv.Atoi(v)
			xi = append(xi, vint)
		}
		x = append(x, xi)
	}

	var count int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var a float64
			for k := 0; k < d; k++ {
				a += math.Pow(math.Abs(float64(x[i][k] - x[j][k])), 2)
			}
			b := math.Sqrt(a)
			if math.Mod(b, 1.0) == 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}
