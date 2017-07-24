//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10 001st prime number?

package main

import "fmt"

func main() {
	var pr int
	var primeness bool
	i := 2
	for {
		primeness = true
		for j := i - 1; j > 1; j-- {
			if (i % j) == 0 {
				primeness = false
			}
		}
		if primeness {
			pr++
			if pr == 10001 {
				fmt.Println(i)
				break
			}
		}
		i++
	}
}

//too slow, needs redesing
