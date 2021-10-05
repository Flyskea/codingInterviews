package main

import (
	"fmt"
	"math"
)

/* func divide(a int, b int) int {
	if a == -0x80000000 && b == -1 {
		return math.MaxInt32
	}
	var (
		result, value, quotient int
		negative                = 2
	)
	if a > 0 {
		negative--
		a = -a
	}
	if b > 0 {
		negative--
		b = -b
	}
	for a <= b {
		value = b
		quotient = 1
		for value >= -0xc0000000 && a <= value+value {
			quotient += quotient
			value += value
		}
		result += quotient
		a -= value
	}
	if negative == 1 {
		return -result
	}
	return result
} */

func divide(a int, b int) int {
	sign := 1
	if a^b < 0 {
		sign = -1
	}
	if a < 0 {
		a = ^a + 1
	}
	if b < 0 {
		b = ^b + 1
	}
	res := 0
	for i := 31; i >= 0; i-- {
		if a>>i >= b {
			if i == 31 && sign == 1 {
				return math.MaxInt32
			}
			res += 1 << i
			a -= b << i
		}
	}
	if sign == -1 {
		return ^res + 1
	}
	return res
}

func main() {
	fmt.Println(0xc0000000)
	fmt.Println(-2147483648 == -0x80000000)
	fmt.Println(divide(-2147483648, -1))
}
