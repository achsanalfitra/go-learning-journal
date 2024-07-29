package controls

func BubbleSort(i *[]int) {
	var newSlice []int
	slice := *i

	for _, v := range slice {
		// add slice element to newSlice
		newSlice = append(newSlice, v)

		// check if newSlice only has one element
		if len(newSlice) == 1 {
			continue
		} else {
			// Run this code if newSlice has more than one element

			// Forward sorting
			for j, w := range newSlice {
				//Checking for the last element
				if j+1 == len(newSlice) {
					continue
				} else {
					// Checking whether the w element is bigger than the next
					if w > newSlice[j+1] {
						newSlice[j], newSlice[j+1] = newSlice[j+1], newSlice[j]
					}
				}
			}
		}

	}

	*i = newSlice
}
