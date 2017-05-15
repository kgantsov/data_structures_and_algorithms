package selection

func Sort(list []int) {
	for pos_to_insert := len(list) - 1; pos_to_insert >= 0; pos_to_insert-- {
		post_of_max := pos_to_insert

		for i := 0; i < pos_to_insert; i++ {
			if list[i] > list[post_of_max] {
				post_of_max = i
			}
		}

		list[pos_to_insert], list[post_of_max] = list[post_of_max], list[pos_to_insert]
	}
}
