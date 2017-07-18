//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func primes(max int) []int {
	pr := []int{}
	var primeness bool
	for i := 2; i <= max; i++ {
		primeness = true
		for j := i - 1; j > 1; j-- {
			if (i % j) == 0 {
				primeness = false
				break
			}
		}
		if primeness {
			pr = append(pr, i)
		}
	}
	return pr
}

func devisibles(primes []int) (result int) {
	result = 1
	for _, curr := range primes {
		pow := curr
		i := 1
		for pow < 20 {
			pow = pow * curr
			i++
		}
		result = result * pow / curr
	}
	return
}

func main() {
	fmt.Println("The number is ", devisibles(primes(20)))
}
