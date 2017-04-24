
def sort(lst):
    for pos_to_insert in range(len(lst) - 1, 0, -1):
        position_of_max = pos_to_insert

        for pos in range(0, pos_to_insert, 1):
            if lst[pos] > lst[position_of_max]:
                position_of_max = pos

        lst[pos_to_insert], lst[position_of_max] = lst[position_of_max], lst[pos_to_insert]

    return lst
