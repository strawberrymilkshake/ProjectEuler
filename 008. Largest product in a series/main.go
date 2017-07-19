/*Find the thirteen adjacent digits in the 1000-digit number
that have the greatest product. What is the value of this product?
*/
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

//Error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Reads l digits from file f and converts them into a slice
func getnumslice(f string, l int) []int {
	var nums = make([]int, l)
	dat, err := ioutil.ReadFile(f)
	check(err)
	for i := 0; i < l; i++ {
		nums[i], _ = strconv.Atoi(string(dat[i]))
	}
	return nums
}

// Returns product of num digits in a slice sl starting with position pos
func getproduct(sl []int, pos int, num int) int {
	var prod = 1
	for i := pos; i < pos+num; i++ {
		prod *= sl[i]
	}
	return prod
}

func main() {
	n := getnumslice("number.txt", 1000)
	var maxprod, pos int
	for i := 0; i < len(n)-13; i++ {
		prod := getproduct(n, i, 13)
		if prod > maxprod {
			maxprod = prod
			pos = i
		}
	}
	fmt.Println("The maximum product is ", maxprod, "from numbers", n[pos:pos+13])
}
