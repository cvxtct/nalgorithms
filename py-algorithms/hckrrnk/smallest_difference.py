def mind_diff_pairs(arr , n):

    ordered = arr
    ordered.sort()
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


if __name__ == "__main__":
    # arr = [-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854]
    arr = [5, 2, 3, 4, 1]
    print(mind_diff_pairs(arr, len(arr)))
    