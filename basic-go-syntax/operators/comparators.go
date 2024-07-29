package operators

func TwoIntsAreTheSame(a, b int) bool {
	return a == b
}

func TwoIntsAreNotTheSame(a, b int) bool {
	return a != b
}

func FirstIntIsGreater(a, b int) bool {
	return a > b
}

func FirstIntIsGreaterThanOrEqualTo(a, b int) bool {
	return a >= b
}

func OneIsTrue(a, b bool) bool {
	return a || b
}

func BothAreTrue(a, b bool) bool {
	return a && b
}

func NotTrue(a bool) bool {
	return !a
}
