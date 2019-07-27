package main

import (
	"bufio"
	"fmt"
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
	s := strings.Split(readLine(), "")
	n := 13
	dp := make([]int, n)
	dp[0] = 1

	mod := 1000000007
	mul := 1
	for i := len(s) - 1; i >= 0; i-- {
		nextDP := make([]int, n)

		if s[i] == "?" {
			for k := 0; k < 10; k++ {
				for j := 0; j < n; j++ {
					nextDP[(k*mul+j)%n] += dp[j]
					nextDP[(k*mul+j)%n] %= mod
				}
			}
		} else {
			k, _ := strconv.Atoi(s[i])
			for j := 0; j < n; j++ {
				nextDP[(k*mul+j)%n] += dp[j]
				nextDP[(k*mul+j)%n] %= mod
			}
		}

		mul *= 10
		mul %= n
		dp = nextDP
	}
	fmt.Println(dp[5])
}
