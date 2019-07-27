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

func judge(a []int) bool {
	mae := a[0]
	for i := 1; i < len(a); i++ {
		if mae+1 != a[i] {
			return false
		}
		mae = a[i]
	}
	return true
}

func main() {
	_ = readInt()
	p := readInts()

	if judge(p) {
		fmt.Println("YES")
		return
	}

	var a int
	var b int
	for i := 0; i < len(p); i++ {
		if p[i] != i+1 {
			a = i
		}
	}
	for i := 0; i < len(p); i++ {
		if p[i] != i+1 && p[i] != p[a] {
			b = i
		}
	}
	p[a], p[b] = p[b], p[a]
	if judge(p) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
