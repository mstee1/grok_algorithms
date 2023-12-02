package algorithms

func Recursion_factorial(num uint) uint {

	if num == 1 {
		return num
	}

	return num * Recursion_factorial(num-1)

}
