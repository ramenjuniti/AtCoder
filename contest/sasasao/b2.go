package main

import (
	"bufio"
	"fmt"
	"os"
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
	s := sc.Text()

	var ca, cb, cc, cd, ce, cf int

	ca = strings.Count(s, "A")
	cb = strings.Count(s, "B")
	cc = strings.Count(s, "C")
	cd = strings.Count(s, "D")
	ce = strings.Count(s, "E")
	cf = strings.Count(s, "F")

	fmt.Printf("%v %v %v %v %v %v\n", ca, cb, cc, cd, ce, cf)
}
