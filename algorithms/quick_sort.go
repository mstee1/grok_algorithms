package algorithms

func Quick_sort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}

	var less, greater, res []int

	mid := len(slice) / 2

	piv := slice[mid]

	for i := 0; i < len(slice); i++ {

		if i == mid {
			continue
		}

		if slice[i] <= piv {
			less = append(less, slice[i])
			continue
		}

		greater = append(greater, slice[i])

	}

	res = append(res, Quick_sort(less)...)
	res = append(res, piv)
	res = append(res, Quick_sort(greater)...)

	return res

}
