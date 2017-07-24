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

//Reads the grid of numbers from file f and converts them into a slice [x][y]
func getnumslice(f string, x int, y int) [][]int {
	var nums = make([][]int, x)
	var num = make([]int, y)
	var column []string
	dat, err := ioutil.ReadFile(f)
	check(err)
	row := strings.Split(string(dat), "\n")

	for i := 0; i < x; i++ {
		column = strings.Split(row[i], " ")
		for j := 0; j < y; j++ {
			num[j], _ = strconv.Atoi(column[j])
		}
		nums[i] = num
	}
	return nums
}

func main() {
	test := getnumslice("grid.txt", 20, 20)
	fmt.Println(test)
}
