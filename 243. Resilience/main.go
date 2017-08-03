// https://projecteuler.net/problem=243
package main

import (
	"fmt"
	"math"
)

func generateprimes() {
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
}

const limit = 100000

var primeness [limit]bool

func isin(list []int, element int) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}
func getdivscnt(n int) int {
	var list []int
	for i := 2; i <= n/2; i++ {
		if primeness[i] {
			if n%i == 0 {
				for j := 1; j <= n/i-1; j++ {
					el := j * i
					if !isin(list, el) {
						list = append(list, el)
					}
				}
			}
		}
	}
	return len(list) + 1
}

//bruteforce
func main() {
	generateprimes()
	i := 3
	for {
		if !primeness[i] {
			k := i - getdivscnt(i)
			rate := float64(k) / float64(i-1)
			if rate < 15499.0/94744.0 {
				fmt.Println("num=", i-k, "/", i-1)
				fmt.Println(i)
				break
			}
		}
		fmt.Println(i)
		i++
	}
	//15499/94744
}
