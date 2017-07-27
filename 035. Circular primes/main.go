/*
The number, 197, is called a circular prime because all rotations of the digits: 1
97, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71,
73, 79, and 97.

How many circular primes are there below one million?
*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

func decompose(n int) []int {
	var s = []int{}

	sl := len(strconv.Itoa(n))
	for i := 0; i < sl; i++ {
		s = append(s, n%10)
		n /= 10
	}
	for j := 0; j < sl/2; j++ {
		s[j], s[sl-j-1] = s[sl-j-1], s[j]
	}
	return s
}

func iscircular(n []int) (result bool) {
	result = true
	for i := 0; i < len(n)-1; i++ {
		n = rotate(n)
		r := 0
		for j := 0; j < len(n); j++ {
			r = r*10 + n[j]
		}
		if !primeness[r] {
			result = false
			return
		}
	}
	return
}

func rotate(n []int) []int {
	l := len(n)
	for i := 0; i < l-1; i++ {
		n[i], n[i+1] = n[i+1], n[i]
	}
	return n
}

func generateprimes() []int {
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
	primes := make([]int, 0, limit)
	for x = 2; x < limit; x++ {
		if primeness[x] {
			primes = append(primes, x)
		}
	}
	fmt.Println("Primeness length is", len(primeness))
	return primes
}

const limit = 10000000

var primeness [limit]bool

func main() {
	var i, counter int
	primes := generateprimes()
	for primes[i] < 1000000 {
		if iscircular(decompose(primes[i])) {
			counter++
			fmt.Println(counter, "-", primes[i])
		}
		i++
	}
	fmt.Println(counter)
}
