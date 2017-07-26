/*
We shall say that an n-digit number is pandigital if it makes use of all the digits
1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be
written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only
include it once in your sum.
*/
package main

import (
	"fmt"
	"strings"
)

func swap(n []int, i int, j int) {
	n[i], n[j] = n[j], n[i]
}

func permute(n []int, b int, e int) {
	if b == e {
		pandigitizer(n)
	} else {
		for i := b; i <= e; i++ {
			swap(n, b, i)
			permute(n, b+1, e)
			swap(n, b, i)
		}
	}
}

func pandigitizer(n []int) {
	a := 10*n[0] + n[1]
	b := 100*n[2] + 10*n[3] + n[4]
	c := 1000*n[5] + 100*n[6] + 10*n[7] + n[8]
	if a*b == c {
		if !strings.Contains(s, string(c)) {
			s = s + string(c) + " "
			sum += c
		}
	}
	a = n[0]
	b = 1000*n[1] + 100*n[2] + 10*n[3] + n[4]
	c = 1000*n[5] + 100*n[6] + 10*n[7] + n[8]
	if a*b == c {
		if !strings.Contains(s, string(c)) {
			s = s + string(c) + " "
			sum += c
		}
	}
}

var sum int
var s string

func main() {
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	permute(num, 0, len(num)-1)
	fmt.Println(sum)
}
