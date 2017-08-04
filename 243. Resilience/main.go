// https://projecteuler.net/problem=243
package main

import (
	"fmt"
	"math"
	"sync"
)

func generateprimes() []int {
	var x, y, n int
	var primeness [limit]bool
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
	return primes
}

func getdivscnt(n int) float64 {
	f := float64(n)
	i := 0
	for primes[i] <= n {
		if n%primes[i] == 0 {
			f *= float64(primes[i] - 1)
			f /= float64(primes[i])
		}
		i++
	}
	return f
}

const limit = 3000000

var primes []int
var wg sync.WaitGroup
var mut sync.Mutex
var lowest float64

//bruteforce
func main() {
	lowest = 1
	primes = generateprimes()
	fmt.Println(getdivscnt(94745))
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 30000*(i-1) + 1; j <= 30000*i; j++ {
				rate := getdivscnt(j) / float64(j-1)
				mut.Lock()
				if rate < lowest {
					//15499.0/94744.0 {
					lowest = rate
					fmt.Printf("New lowest rate - %f for n - %v\n", rate, j)
				}
				mut.Unlock()
			}
			wg.Done()
		}(i)

	}
	//15499/94744
	wg.Wait()
}
