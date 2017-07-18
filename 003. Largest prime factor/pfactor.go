// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func largestPrimeFactor(num int) int {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			num = num / i
		}
	}
	return num
}

func main() {
	num := 600851475143
	fmt.Println("The largest prime factor is ", largestPrimeFactor(num))
}
