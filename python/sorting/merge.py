
def merge(left, right):
    result = []

    while left and right:
        if left[0] <= right[0]:
            result.append(left[0])
            left = left[1:]
        else:
            result.append(right[0])
            right = right[1:]

    if left:
        result += left
    if right:
        result += right

    return result


def sort(lst):
    if len(lst) <= 1:
        return lst

    middle = int(len(lst) / 2)
    left, right = sort(lst[:middle]), sort(lst[middle:])

    return merge(left, right)
