/*


The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/
package main

import (
	"fmt"
	"math"
)

func iscircular(n int) (result bool) {
	var a, b int
	a = n
	result = true
	if n < 10 {
		return
	}
	for b != n {
		b = rotate(a)
		a = b
		if !primeness[b] {
			result = false
		}
	}
	return
}

func rotate(n int) int {
	//not ready
	l := len(string(n))
	rotation := n / 10
	if rotation == 0 {
		return n
	}
	rotation += (n%10)*10 ^ (l - 1)
	return rotation
}

func generateprimes() []int {
	var x, y, n int
	nsqrt := math.Sqrt(limit)
	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				primeness[n] = !primeness[n]
			}
			n = 3*(x*x) + y*y
			if n <= limit && n%12 == 7 {
				primeness[n] = !primeness[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= limit && n%12 == 11 {
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
	primes := make([]int, 0, 1270606)
	for x = 2; x < len(primeness)-1; x++ {
		if primeness[x] {
			primes = append(primes, x)
		}
	}
	return primes
}

const limit = 1000000

var primeness [limit]bool

func main() {
	var i, counter int
	primes := generateprimes()
	for primes[i] < limit {
		if iscircular(primes[i]) {
			counter++
			fmt.Println(counter)
		}
		fmt.Println(i)
		i++
	}
}
