package algorithms

import "fmt"

func Binaty_search() {

	items := []int{3, 7, 11}

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, item := range items {
		var found bool

		steps := 0
		low := 0
		high := len(a) - 1

		for low <= high {
			steps++
			mid := (low + high) / 2
			guess := a[mid]
			if guess == item {
				fmt.Printf("Found item - %d, index - %d, "+
					"steps - %d\n", item, mid, steps)
				found = true
				break
			}
			if guess > item {
				high = mid - 1
				continue
			}
			low = mid + 1
		}

		if !found {
			fmt.Printf("Item not found - %d\n", item)
		}
	}

}
