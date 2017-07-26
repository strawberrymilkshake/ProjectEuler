/*
A number chain is created by continuously adding the square of the digits in a
number to form a new number until it has been seen before.

For example,

44 → 32 → 13 → 10 → 1 → 1
85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop.
What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.

How many starting numbers below ten million will arrive at 89?
*/

package main

import "fmt"

func process(num int) bool {
	var sum int
	i := num
	for i > 0 {
		sum += (i % 10) * (i % 10)
		i = i / 10
	}
	if sum == 1 {
		return false
	} else if sum == 89 {
		return true
	}
	return process(sum)
}

func main() {
	var counter int
	for i := 9999999; i >= 2; i-- {
		if process(i) {
			counter++
		}
	}
	fmt.Println(counter)
}
