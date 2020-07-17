package section1

import (
	"fmt"
	"strings"
)

func Fibonacci(depth int, result ...interface{}) {
	// first initiation of the result
	// due to Fibonacci would always start from 1
	// as starting point
	if len(result) < 1 {
		result = []interface{}{1}
	}
	// less than equal to 1 due to initials of [1] in the result array
	if depth <= 1 {
		// print out the result to screen
		fmt.Printf(strings.Repeat("%v ", len(result)), result...)
		return
	}
	// append value to the result
	var prev, current = 0, 0
	if len(result) > 0 {
		current = result[len(result)-1].(int)
	}
	if len(result) > 1 {
		prev = result[len(result)-2].(int)
	}
	result = append(
		result,
		prev+current,
	)
	// recursive call for the next iteration -- decremental
	// including the last result to be appended with the next one
	Fibonacci(depth-1, result...)
}

func Fibonacci2(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci2(n-1) + Fibonacci2(n-2)
}
