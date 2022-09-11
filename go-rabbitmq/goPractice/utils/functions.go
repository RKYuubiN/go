package utils

func ReturnsTwoValues(x int, y int) (int, int) {
	a := x + y
	b := x * y
	return a, b
}

func ProAdder(values ...int) (int, string) {
	total := 0
	for _, values := range values {
		total += values
	}
	return total, "Hello"
}
