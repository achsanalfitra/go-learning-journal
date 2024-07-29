package controls

func BubbleSortSolution(initial *[]int) {
	slice := *initial
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				slice[i], slice[i-1] = slice[i-1], slice[i]
				swapped = true
			}
		}
	}

	*initial = slice
}
