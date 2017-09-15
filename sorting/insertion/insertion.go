package insertion

func Sort(list []int) {
	for i := 1; i < len(list); i++ {
		value := list[i]
		j := i - 1

		for j >= 0 && list[j] > value {
			list[j+1] = list[j]
			j--
		}

		list[j+1] = value
	}
}
