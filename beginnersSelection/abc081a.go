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
	inputLine := scanLine()
	inputArray := strings.Split(inputLine, "")
	ans := 0
	for _, i := range inputArray {
		if i == "1" {
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
}
