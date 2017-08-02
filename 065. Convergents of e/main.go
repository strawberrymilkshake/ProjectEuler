/*description:
https://projecteuler.net/problem=65
*/
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func mkfrac(i, limit int64) *big.Rat {
	num := new(big.Rat)
	if i%3 != 1 {
		num.SetInt64(1)
	} else {
		num.SetInt64((i - 1) / 3 * 2)
	}
	if i >= limit {
		return num
	}
	return num.Add(num, new(big.Rat).Inv(mkfrac(i+1, limit)))
}

func main() {
	var sum int
	r := mkfrac(0, 101).Num().String()
	for i := 0; i < len(r); i++ {
		digit, _ := strconv.Atoi(fmt.Sprintf("%c", r[i]))
		sum += digit
	}
	fmt.Println(sum)
}
