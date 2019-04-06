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
	var max int
	var ans int
	for i := 0; i < 5; i++ {
		in, _ := strconv.Atoi(readLine())
		if max < (10-in%10)%10 {
			max = (10 - in%10) % 10
		}
		ans += in + (10-in%10)%10
	}
	ans = ans - max
	fmt.Println(ans)
}
