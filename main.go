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

	fmt.Println()
	fmt.Println("Start merge sort for slice")
	z := []int{5, 11, 56, 1, 0, 58, 5, 1, 2, 4, 2, 9}
	fmt.Println(z)
	z = algorithms.Merge_sort(z)
	fmt.Printf("Sorted slice - %v\n", z)

	fmt.Println()
	fmt.Println("Start quick sort for slice")
	z = []int{5, 11, 56, 1, 0, 58, 5, 1, 2, 4, 2, 9}
	fmt.Println(z)
	z = algorithms.Quick_sort(z)
	fmt.Printf("Sorted slice - %v\n", z)
}
