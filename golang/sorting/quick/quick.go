package quick

func Sort(list []int) {
	SortHelper(list, 0, len(list) - 1)
}

func SortHelper(list []int, left int, right int) {
	if len(list) == 0 {
		return
	}

	index := partition(list, left, right)

	if left < index - 1 {
		SortHelper(list, left, index - 1)
	}

	if index < right {
		SortHelper(list, index, right)
	}

	return
}

func partition(list []int, left int, right int) int {
	pivot := list[(left + right) / 2]

	for left <= right {
		for list[left] < pivot {
			left++
		}

		for list[right] > pivot {
			right--
		}

		if left <= right {
			list[left], list[right] = list[right], list[left]
			left++
			right--
		}
	}
	return left
}
