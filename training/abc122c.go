package main

import (
	"fmt"
)

func main() {

	var (
		n, q int
		s    string
	)
	fmt.Scan(&n, &q, &s)

	count := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		count[i] = count[i-1]
		if s[i-1] == 'A' && s[i] == 'C' {
			count[i]++
		}
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Println(count[r-1] - count[l-1])
	}

}
