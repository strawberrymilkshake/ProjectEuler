package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fromroman(s string) (sum int) {
	sum = 0
	for _, v := range []string{"CM", "CD", "XC", "XL", "IX", "IV"} {
		if i := strings.Index(s, v); i != -1 {
			s = s[:i] + s[i+2:]
			sum += roman[v]
		}
	}
	for i := 0; i < len(s); i++ {
		sum += roman[s[i:i+1]]
	}
	return
}

func toroman(n int) string {
	var s string
	var CM, CD, XC, XL, IX, IV bool
	for n > 0 {
		if (n - 1000) >= 0 {
			s = s + decimal[1000]
			n -= 1000
		} else if !CM && (n-900) >= 0 {
			s = s + decimal[900]
			CM = true
			n -= 900
		} else if (n - 500) >= 0 {
			s = s + decimal[500]
			n -= 500
		} else if !CD && (n-400) >= 0 {
			s = s + decimal[400]
			CD = true
			n -= 400
		} else if (n - 100) >= 0 {
			s = s + decimal[100]
			n -= 100
		} else if !XC && (n-90) >= 0 {
			s = s + decimal[90]
			XC = true
			n -= 90
		} else if (n - 50) >= 0 {
			s = s + decimal[50]
			n -= 50
		} else if !XL && (n-40) >= 0 {
			s = s + decimal[40]
			XL = true
			n -= 40
		} else if (n - 10) >= 0 {
			s = s + decimal[10]
			n -= 10
		} else if !IX && (n-9) >= 0 {
			s = s + decimal[9]
			IX = true
			n -= 9
		} else if (n - 5) >= 0 {
			s = s + decimal[5]
			n -= 5
		} else if !IV && (n-4) >= 0 {
			s = s + decimal[4]
			IV = true
			n -= 4
		} else if (n - 1) >= 0 {
			s = s + decimal[1]
			n -= 1
		}
	}
	return s
}

func difference(s string) int {
	l1 := len(s)
	l2 := len(toroman(fromroman(s)))
	return l1 - l2
}

var roman = map[string]int{"I": 1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000}

var decimal = map[int]string{1: "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M"}

func main() {
	var sum int
	dat, err := ioutil.ReadFile("p089_roman.txt")
	check(err)
	row := strings.Split(string(dat), "\n")
	for _, v := range row {
		sum += difference(v)
	}
	fmt.Println(sum)
}
