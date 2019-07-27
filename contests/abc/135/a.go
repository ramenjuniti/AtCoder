package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// input に使うやつ
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

// string から int へ
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// 一つの文字を int として読み込み
func readInt() int {
	return atoi(readLine())
}

// 複数の文字を int のスライスとして読み込み
func readInts() []int {
	in := readLine()
	ins := strings.Split(in, " ")
	out := make([]int, len(ins))
	for i, v := range ins {
		out[i] = atoi(v)
	}
	return out
}

// sortしたスライスを返す
func sorted(a []int) []int {
	sort.Ints(a)
	b := a
	return b
}

// 逆順のsort
func reverse(a []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	b := a
	return b
}

func main() {
	n := readInts()
	a := n[0]
	b := n[1]
	var k int
	if a > b {
		if (a-b)%2 == 0 {
			k = (a - b) / 2
			if a+k == b-k {
				fmt.Println(int(math.Abs(float64(a + k))))
			} else {
				fmt.Println(int(math.Abs(float64(a - k))))
			}
		} else {
			fmt.Println("IMPOSSIBLE")
		}
	} else {
		if (b-a)%2 == 0 {
			k = (b - a) / 2
			if b+k == a-k {
				fmt.Println(int(math.Abs(float64(b + k))))
			} else {
				fmt.Println(int(math.Abs(float64(b - k))))
			}
		} else {
			fmt.Println("IMPOSSIBLE")
		}
	}
}
