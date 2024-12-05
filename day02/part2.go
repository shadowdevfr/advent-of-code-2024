package day02

func isListSafeWithDampener(tab []int) bool {
	// We go through all the elements of the list and remove them one by one, then we check if the list is safe
	for i := 0; i < len(tab); i++ {
		var newslice []int
		newslice = append(newslice, tab[:i]...)
		newslice = append(newslice, tab[i+1:]...)

		if isListSafe(newslice) {
			return true
		}
	}

	return false
}
