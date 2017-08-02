/*Euler discovered the remarkable quadratic formula:

n^2+n+41

It turns out that the formula will produce 40 primes for the consecutive
integer values 0≤n≤39. However, when n=40,402+40+41=40(40+1)+41 is divisible
by 41, and certainly when n=41,412+41+41

is clearly divisible by 41.

The incredible formula n2−79n+1601
was discovered, which produces 80 primes for the consecutive values 0≤n≤79

. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

    n2+an+b

, where |a|<1000 and |b|≤1000

where |n|
is the modulus/absolute value of n
e.g. |11|=11 and |−4|=4

Find the product of the coefficients, a
and b, for the quadratic expression that produces the maximum number of primes
for consecutive values of n, starting with n=0.
*/

package main

import (
	"fmt"
	"math"
)

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return num * (-1)
}

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

const limit = 10000000

var primeness [limit]bool

func main() {
	generateprimes()
	var a, b, maxn, maxprod int
	for a = 999; a >= -999; a-- {
		for b = 1000; b >= -1000; b-- {
			var n, prod int
			for {
				prod = n*(n+a) + b
				if primeness[abs(prod)] {
					n++
				} else {
					break
				}
			}
			if n > maxn {
				maxn = n
				maxprod = a * b
			}
		}
	}
	fmt.Println("Maximum product is", maxprod)
}
