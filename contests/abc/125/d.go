package main

import (
	"bufio"
	"fmt"
	"math"
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
	astr := strings.Split(readLine(), " ")

	var a []int
	var ncount int

	for _, v := range astr {
		avalue, _ := strconv.Atoi(v)
		a = append(a, avalue)
		if avalue < 0 {
			ncount++
		}
	}

	var ans int

	if ncount%2 == 0 {
		for _, v := range a {
			if v < 0 {
				ans += -v
			} else {
				ans += v
			}
		}
	} else {
		amin := 1000000000
		var aminindex int
		for i, v := range a {
			if amin > int(math.Abs(float64(v))) {
				amin = int(math.Abs(float64(v)))
				aminindex = i
			}
		}
		for i, v := range a {
			if i == aminindex {
				if v > 0 {
					ans += -v
				} else {
					ans += v
				}
			} else {
				if v < 0 {
					ans += -v
				} else {
					ans += v
				}
			}
		}
	}

	fmt.Println(ans)
}
