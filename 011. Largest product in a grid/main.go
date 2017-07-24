package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
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

var mut sync.Mutex

func main() {
	var wg sync.WaitGroup

	dat := getnumslice("grid.txt", 20, 20)
	var product, maxproduct int

	wg.Add(4)
	//search down
	go func() {
		for i := 0; i <= 16; i++ {
			for j := 0; j <= 19; j++ {
				mut.Lock()
				product = dat[i][j] * dat[i+1][j] * dat[i+2][j] * dat[i+3][j]
				if product > maxproduct {
					maxproduct = product
				}
				mut.Unlock()
			}
		}
		wg.Done()
	}()

	//search right
	go func() {
		for i := 0; i <= 19; i++ {
			for j := 0; j <= 16; j++ {
				mut.Lock()
				product = dat[i][j] * dat[i][j+1] * dat[i][j+2] * dat[i][j+3]
				if product > maxproduct {
					maxproduct = product
				}
				mut.Unlock()
			}
		}
		wg.Done()
	}()

	//search diagonally down-right
	go func() {
		for i := 0; i < 17; i++ {
			for j := 0; j < 17; j++ {
				mut.Lock()
				product = dat[i][j] * dat[i+1][j+1] * dat[i+2][j+2] * dat[i+3][j+3]
				if product > maxproduct {
					maxproduct = product
				}
				mut.Unlock()
			}
		}
		wg.Done()
	}()

	//search diagonally down-right and up-right
	go func() {
		for i := 19; i >= 3; i-- {
			for j := 0; j <= 16; j++ {
				mut.Lock()
				product = dat[i][j] * dat[i-1][j+1] * dat[i-2][j+2] * dat[i-3][j+3]
				if product > maxproduct {
					maxproduct = product
				}
				mut.Unlock()
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(maxproduct)
}
