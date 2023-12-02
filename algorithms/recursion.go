package algorithms

func Recursion_factorial(num uint) uint {

	if num == 1 {
		return num
	}

	return num * Recursion_factorial(num-1)

}

func Recursion_count(slice []int) int {

	if len(slice) == 1 {
		return 1
	}

	return 1 + Recursion_count(slice[1:])
}
