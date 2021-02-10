package basic

func For1(param int) int {
	var sum int
	for i := 0; i < param; i++ {
		sum = i
	}

	return sum
}

func For2(param int) int {
	i := 0
	for i < param {
		i++
	}

	return i
}
