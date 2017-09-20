//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	func() {
		sum := 0

		for i := 1; i <= 333; i++ {
			sum += i * 3
			if i <= 199 && ((i*5)%3) != 0 {
				sum += i * 5
			}
		}
		fmt.Println("The sum of all the multiples of 3 or 5 below 1000 is", sum)
	}()
	fmt.Println("Execution time was:", time.Since(t))
}
