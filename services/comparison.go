package services

// Min returns the minimum of 2 integers
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the minimum of 2 integers
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
