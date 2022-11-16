def mind_diff_pairs(arr , n):

    ordered = arr
    ordered.sort()

    minimum_distance = ordered[1] - ordered[0] 
    closest_pairs = []
    for i in range(2, n):
        minimum_distance = min(minimum_distance, ordered[i] - ordered[i-1])

    for i in range(1, n):
        if (ordered[i] - ordered[i-1]) == minimum_distance:
            closest_pairs.append(ordered[i-1])
            closest_pairs.append(ordered[i])

    return closest_pairs


if __name__ == "__main__":
    # arr = [-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854]
    arr = [5, 2, 3, 4, 1]
    print(mind_diff_pairs(arr, len(arr)))
    