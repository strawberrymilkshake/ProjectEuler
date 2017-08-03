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
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func runetoint(r byte) int {
	return int(r - '0')
}

func isin(list []string, element string) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}

type euler79 struct {
	numlist  []string
	verses   map[int][]string
	passcode string
}

func (e *euler79) init() {
	e.numlist = *new([]string)
	e.verses = *new(map[int][]string)
	e.passcode = ""
}

func (e *euler79) fill(f string) {
	dat, err := ioutil.ReadFile(f)
	check(err)
	nums := strings.Split(string(dat), "\n")
	e.numlist = nums[:len(nums)-1]
	verses := make(map[int][]string)
	for _, val := range e.numlist {
		vers := val[:2]
		pos := runetoint(vers[0])
		if !isin(verses[pos], vers) {
			verses[pos] = append(verses[pos], vers)
		}
		vers = val[1:]
		pos = runetoint(vers[0])
		if !isin(verses[pos], vers) {
			verses[pos] = append(verses[pos], vers)
		}
	}
	e.verses = verses
}

//recursive search for longest passcode
func (e *euler79) buildup(str string) {
	next := runetoint(str[len(str)-1])
	if e.verses[next] != nil {
		for _, j := range e.verses[next] {
			if strings.Index(str, j[1:]) == -1 {
				var buffer bytes.Buffer
				buffer.WriteString(str)
				buffer.WriteString(j[len(j)-1:])
				e.buildup(buffer.String())
			}
		}
	} else {
		if len(str) > len(e.passcode) {
			e.passcode = str
		}
	}
}

func (e *euler79) process() {
	//parsing each possible first digit, constantly building up the key string
	for i := 0; i <= 9; i++ {
		if !(e.verses[i] == nil) {
			for _, vers := range e.verses[i] {
				e.buildup(vers)
			}
		}
	}
}

func main() {
	var e = new(euler79)
	e.init()
	e.fill("p079_keylog.txt")
	e.process()
	fmt.Println("The passcode is", e.passcode)
}
