
def sort(lst):
    for i in range(0, len(lst) - 1):
        for j in range(1, len(lst) - i):
            if lst[j - 1] > lst[j]:
                lst[j - 1], lst[j] = lst[j], lst[j - 1]
    return lst
