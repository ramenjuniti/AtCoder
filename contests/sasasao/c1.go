package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	input := strings.Split(sc.Text(), " ")
	var in []int
	for _, v := range input {
		t, _ := strconv.Atoi(v)
		in = append(in, t)
	}

	if in[4]+in[3]+in[0] > in[4]+in[2]+in[1] {
		fmt.Println(in[4] + in[3] + in[0])
	} else {
		fmt.Println(in[4] + in[2] + in[1])
	}
}
