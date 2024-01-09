package simplealgos

func GCD(i, j int) int {
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}
	if i == j {
		return i
	}
	if i > j {
		return GCD(i - j, j)
	}
	return GCD(i, j - i)
}
