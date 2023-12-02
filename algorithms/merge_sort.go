package algorithms

func Merge_sort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := Merge_sort(slice[:mid])
	right := Merge_sort(slice[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {

	res := make([]int, 0, len(left)+len(right))

	for len(left) > 0 && len(right) > 0 {

		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
			continue
		}
		res = append(res, right[0])
		right = right[1:]
	}

	res = append(res, left...)
	res = append(res, right...)

	return res
}
