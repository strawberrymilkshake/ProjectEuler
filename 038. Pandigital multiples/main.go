/*
Take the number 192 and multiply it by each of 1, 2, and 3:
    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576.
We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3,
4, and 5, giving the pandigital, 918273645, which is the concatenated
product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be
formed as the concatenated product of an integer with (1,2, ... , n)
where n > 1?
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ispandigital(s string) (is bool) {
	is = true
	if len(s) != 9 {
		is = false
		return
	}
	for i := 1; i <= 9; i++ {
		subs := strconv.Itoa(i)
		if !strings.Contains(s, subs) {
			is = false
		}
	}
	return
}

func main() {
	var max int
	for x := 1; x <= 9999; x++ {
		d := strconv.Itoa(x)
		for i := 2; i <= 5; i++ {
			d += strconv.Itoa(i * x)
			if ispandigital(d) {
				pd, _ := strconv.Atoi(d)
				if pd > max {
					max = pd
				}
			}
		}
	}
	fmt.Println(max)
}
