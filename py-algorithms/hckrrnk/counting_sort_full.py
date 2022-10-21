def counting_sort(n, arr):
    # count = [0 for _ in range(max(arr)+1)]
    # prep the input
    # b = a.split(' ')
    # [[b[i],b[i+1]] for i in range(0, len(b), 2)]
    # cast first elem to int instead
    # [[int(b[i]),b[i+1]] for i in range(0, len(b), 2)]
    count = [0 for _ in range(20)]
    # count = [0] * 100
    seen = set()

    for i in range(len(arr)):
        if arr[i][0] not in seen:
            seen.add(arr[i][0])
        if arr[i][0] in seen:
            count[arr[i][0]] += 1

    res = []
    for i in range(len(count)):
        if count[i] > 0:
            cnt = count[i]
            while cnt > 0:
                res.append(i)
                cnt -= 1


    return count, res

            


if __name__ == '__main__':
    n = 20
    arr = [[0, 'ab'], [6, 'cd'], [0, 'ef'], [6, 'gh'], [4, 'ij'], [0, 'ab'], [6, 'cd'], [0, 'ef'], [6, 'gh'], [0, 'ij'], [4, 'that'], [3, 'be'], [0, 'to'], [1, 'be'], [5, 'question'], [1, 'or'], [2, 'not'], [4, 'is'], [2, 'to'], [4, 'the']]
    print(counting_sort(n, arr))