/*
A common security method used for online banking is to ask the user for three
random characters from a passcode. For example, if the passcode was 531278,
they may ask for the 2nd, 3rd, and 5th characters; the expected reply would be: 317.

The text file, keylog.txt, contains fifty successful login attempts.

Given that the three characters are always asked for in order, analyse the file
so as to determine the shortest possible secret passcode of unknown length.
*/

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

//dummy
func main() {
	dat, err := ioutil.ReadFile("p079_keylog.txt")
	check(err)
	p := [10][10]int{}
	nums := strings.Split(string(dat), "\n")
	for _, v := range nums {
		num, _ := strconv.Atoi(v)
		if num == 0 {
			break
		}
		third := num % 10
		num /= 10
		second := num % 10
		num /= 10
		first := num
		p[first][second] = 1
		p[second][third] = 1
	}
	for i := 0; i <= 9; i++ {
		fmt.Println(p[i])
	}

}
