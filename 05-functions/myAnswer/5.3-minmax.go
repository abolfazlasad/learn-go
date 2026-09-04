// Question 5.3 — Multiple return values
//
// minmax(nums []int) (min, max int, err error)
// Empty slice is an error. Otherwise return min and max.
//
// Extra: div(a, b int) (q, r int, err error) — error on division by zero.
//
// You will need: multiple results, error, errors.New or fmt.Errorf.
//
// Run:
//   go run 5.3-minmax.go

package main

import (
	"errors"
	"fmt"
)

func minmax(num []int) (int, int, error) {
	if len(num) == 0 {
		return 0, 0, errors.New("empty slice")
	}
	min, max := num[0], num[0]
	for _, n := range num {
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}
	return min, max, nil

}

func main() {
	nums := []int{3, 1, 4, 1, 5}
	min, max, err := minmax(nums)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(min, max)

	fmt.Println("----------------------------")

	nums = []int{}
	min, max, err = minmax(nums)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(min, max)
}
