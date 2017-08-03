// https://projecteuler.net/problem=81

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//read matrix from file
func getnumslice(f string, x int, y int) [][]int {
	var nums = make([][]int, x)
	var column []string
	dat, err := ioutil.ReadFile(f)
	check(err)
	row := strings.Split(string(dat), "\n")
	for i := 0; i < x; i++ {
		column = strings.Split(row[i], ",")
		nums[i] = make([]int, y)
		for j := 0; j < y; j++ {
			nums[i][j], _ = strconv.Atoi(column[j])
		}
	}
	return nums
}

func main() {
	var matrix = getnumslice("p081_matrix.txt", 80, 80)
	l := len(matrix) - 1
	//preparing last row and column
	for i := l - 1; i >= 0; i-- {
		matrix[i][l] += matrix[i+1][l]
		matrix[l][i] += matrix[l][i+1]
	}

	//calculating the path
	for i := l - 1; i >= 0; i-- {
		for j := l - 1; j >= 0; j-- {
			right := matrix[i][j] + matrix[i][j+1]
			down := matrix[i][j] + matrix[i+1][j]
			if right < down {
				matrix[i][j] = right
			} else {
				matrix[i][j] = down
			}
		}
	}
	fmt.Println("Minimal path sum is", matrix[0][0])
}
