//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

//Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "fmt"

func mirror(num int) (mirr int) {
	position := 100
	for i := 0; i <= 2; i++ {
		mirr = mirr + (num%10)*(position)
		num = num / 10
		position = position / 10
	}
	return
}

func isPalindrome(num int) bool {
	if (num / 1000) == mirror(num%1000) {
		return true
	}
	return false
}

func main() {
	var largest int
	for i := 999; i >= 111; i-- {
		for j := 999; j >= 111; j-- {
			if isPalindrome(i*j) && i*j > largest {
				largest = i * j
			}
		}
	}
	fmt.Println(largest)
}
