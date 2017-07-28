/*
It was proposed by Christian Goldbach that every odd composite
number can be written as the sum of a prime and twice a square.

9 = 7 + 2×1^2
15 = 7 + 2×2^2
21 = 3 + 2×3^2
25 = 7 + 2×3^2
27 = 19 + 2×2^2
33 = 31 + 2×1^2

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the
sum of a prime and twice a square?
*/

package main

import (
	"fmt"
	"math"
)

const limit = 1000000

var primeness [limit]bool

//genodds generates and returns slice of odd composites <=limit
//It also fills primeness slice with index representing the number, and value represents its primeness
func genodds() []int {
	var x, y, n int
	nsqrt := math.Sqrt(limit)
	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if (n <= limit) && (n%12 == 1 || n%12 == 5) {
				primeness[n] = !primeness[n]
			}

			n = 3*(x*x) + y*y
			if (n <= limit) && (n%12 == 7) {
				primeness[n] = !primeness[n]
			}

			n = 3*(x*x) - y*y
			if (x > y) && (n <= limit) && (n%12 == 11) {
				primeness[n] = !primeness[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if primeness[n] {
			for y = n * n; y < limit; y += n * n {
				primeness[y] = false
			}
		}
	}

	primeness[2] = true
	primeness[3] = true
	odds := make([]int, 0, limit)
	for x = 2; x < limit; x++ {
		if !primeness[x] && x%2 != 0 {
			odds = append(odds, x)
		}
	}
	return odds
}

//dummy
func main() {
	odds := genodds()
	//iterate throug odd composites
	for _, v := range odds {
		i := 1
		fits := false
		//trying to find twice a square that will satisfy the requirement
		for {
			if (v - 2*i*i) < 0 {
				break
			} else {
				if primeness[v-2*i*i] {
					fits = true
					break
				}
			}
			i++
		}
		//if non was found - print the number and break
		if fits == false {
			fmt.Println(i, v)
			break
		}
	}
}
