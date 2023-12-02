package main

import (
	"fmt"

	"github.com/mstee1/grok_algorithms/algorithms"
)

func main() {

	fmt.Println("Start binary search")
	algorithms.Binaty_search()

	fmt.Println()
	fmt.Println("Start selection sort")
	algorithms.Selection_sort()

	fmt.Println()
	fmt.Println("Start recursion for num 5")
	x := algorithms.Recursion_factorial(5)
	fmt.Printf("10! = %d\n", x)

	fmt.Println()
	fmt.Println("Start recursion for 5 elements slice")
	y := algorithms.Recursion_count([]int{1, 2, 3, 4, 5})
	fmt.Printf("Count elements = %d\n", y)
}
