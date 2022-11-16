def counting_sort(collection):

    # if the collection is empty, returns empty
    if collection == []:
        return []

    # get some information about the collection
    coll_len = len(collection)
    coll_max = max(collection)
    coll_min = min(collection)

    # create the counting array
    counting_arr_length = coll_max + 1 - coll_min
    counting_arr = [0] * counting_arr_length

    # count how much a number appears in the collection
    for number in collection:
        counting_arr[number - coll_min] += 1

    # sum each position with it's predecessors. now, counting_arr[i] tells
    # us how many elements <= i has in the collection
    for i in range(1, counting_arr_length):
        counting_arr[i] = counting_arr[i] + counting_arr[i - 1]

    
    # create the output collection
    ordered = [0] * coll_len

    # place the elements in the output, respecting the original order (stable
    # sort) from end to begin, updating counting_arr
    for i in reversed(range(0, coll_len)):
        ordered[counting_arr[collection[i] - coll_min] - 1] = collection[i]
        counting_arr[collection[i] - coll_min] -= 1

    # get closest pairs
    
    minimum_distance = max(ordered)
    closest_pairs = []
    for i in range(0, len(ordered) - 1):
        calc_distance = ordered[i+1] - ordered[i]
        if minimum_distance > calc_distance:
            minimum_distance = calc_distance

    print(minimum_distance)
    print(ordered[(ordered.index(max(ordered)))])
    print(ordered)
    return ordered


def counting_sort_string(string):
    """
    >>> counting_sort_string("thisisthestring")
    'eghhiiinrsssttt'
    """
    return "".join([chr(i) for i in counting_sort([ord(c) for c in string])])


if __name__ == "__main__":
    # arr = [-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854]
    arr = [5, 2, 3, 4, 1]
    print(counting_sort(arr))
    