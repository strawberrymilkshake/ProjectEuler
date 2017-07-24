/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

/*
Given any positive integers m and n where m > n > 0, the integers that will form a Pythagorean triplet are:
a =m^2-n^2
b =2m*n
c =m^2+n^2
m^2-n^2 + 2m*n + m^2+n^2  = 1000
2*m^2 + 2mn = 1000
m(m+n) = 500
m+n > m, thus m < sqrt(500) and m <= 22
*/

package main

import "fmt"

func main() {
	var a, b, c int
	for m := 22; m >= 1; m-- {
		if (500 % m) == 0 {
			n := 500/m - m
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if a+b+c == 1000 {
				fmt.Println("a*b*c=", a*b*c)
				break
			}
		}
	}
}
