package unittests

func isPositive(i int) string {
	if i == 0 {
		return "i is 0"
	}
	if i < 0 {
		return "is less than 0"
	}
	return "i is greater than 0"

}
