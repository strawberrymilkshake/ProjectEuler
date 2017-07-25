//Work out the first ten digits of the sum of the one-hundred 50-digit numbers from numbers.txt

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Reads x numbers of y digits each from file f and calculate their sum
//sum is returned as []int in reverse order
func getsumfromfile(f string, x int, y int) []int {
	var nums = make([]int, y+2)
	var sum, num, k int
	dat, err := ioutil.ReadFile(f)
	check(err)
	row := strings.Split(string(dat), "\n")
	for i := y - 1; i >= 0; i-- {
		for j := 0; j < x; j++ {
			num, _ = strconv.Atoi(string(row[j][i]))
			sum += num
		}
		nums[k] = sum % 10
		sum = sum / 10
		k++
	}
	for sum > 0 {
		nums[k] = sum % 10
		sum = sum / 10
		k++
	}
	return nums
}

func main() {
	dat := getsumfromfile("numbers.txt", 100, 50)
	fmt.Println("First 10 digits of sum is:")
	i := len(dat) - 1
	for i >= len(dat)-10 {
		fmt.Print(dat[i])
		i--
	}
	fmt.Println("")
}
