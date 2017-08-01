package main

import (
	"fmt"
	"math"
)

// Only primes less than or equal to limit will be generated
const limit = 1000000

func main() {
	//https://en.wikipedia.org/wiki/Sieve_of_Atkin
	var x, y, n, sum int
	nsqrt := math.Sqrt(limit)
	primeness := [limit]bool{}
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

	// primes is now a slice that contains all primes numbers up to limit
	// so let's print them
	for _, x := range primes {
		sum += x
	}
	fmt.Println(sum)
}
