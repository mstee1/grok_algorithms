package algorithms

import "fmt"

func Selection_sort() {

	items := []int{4, 5, 2, 6, 3, 7, 3, 7, 11}
	sliceLength := len(items)
	fmt.Println("Source slice - ", items)

	sortItems := make([]int, len(items))

	for i := 0; i < sliceLength; i++ {

		smallest := items[0]
		smallestIndex := 0

		for j := 1; j < len(items); j++ {

			if items[j] < smallest {
				smallest = items[j]
				smallestIndex = j
			}
		}

		items = append(items[:smallestIndex], items[smallestIndex+1:]...)

		sortItems[i] = smallest
	}

	fmt.Println("Sorted slice - ", sortItems)

}
