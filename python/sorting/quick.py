
def sort(lst, left=None, right=None):
    if not lst:
        return lst

    left = left or 0
    right = right or len(lst) - 1

    index = partition(lst, left, right)

    if left < index - 1:
        sort(lst, left, index - 1)

    if index < right:
        sort(lst, index, right)

    return lst


def partition(lst, left, right):
    pivot = lst[(left + right) // 2]

    while left <= right:
        while lst[left] < pivot:
            left += 1

        while lst[right] > pivot:
            right -= 1

        if left <= right:
            lst[left], lst[right] = lst[right], lst[left]
            left += 1
            right -= 1

    return left
