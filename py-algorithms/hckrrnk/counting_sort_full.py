def counting_sort(n, arr):
    arr_len = len(arr) # this is actualy given
    arr_max = int(max([i[0] for i in arr]))
    arr_min = int(min([i[0] for i in arr]))
    
    count_arr_len = arr_max + 1 - arr_min

    count = [0] * count_arr_len
    seen = set()

    for number in arr:
        count[int(number[0])-arr_min] += 1
    
    print(count)

    # replace character with '-' in arr[i][1]
    for i in range(0, arr_len//2):
        arr[i][1] = '-'


    for i in range(1, count_arr_len):
        # this identify the slot where the actual elem belongs to
        count[i] = count[i] + count[i-1]

    print(count)

    ordered = [0] * arr_len
    for i in reversed(range(0, arr_len)):
        ordered[count[int(arr[i][0]) - arr_min] -1] = arr[i]
        count[int(arr[i][0]) - arr_min] -= 1
    
    to_print = []
    print(" ".join(map(str, [i[1] for i in ordered])))
    # return ordered


if __name__ == '__main__':
    n = 20
    arr = [[0, 'ab'], [6, 'cd'], [0, 'ef'], [6, 'gh'], [4, 'ij'], [0, 'ab'], [6, 'cd'], [0, 'ef'], [6, 'gh'], [0, 'ij'], [4, 'that'], [3, 'be'], [0, 'to'], [1, 'be'], [5, 'question'], [1, 'or'], [2, 'not'], [4, 'is'], [2, 'to'], [4, 'the']]
    inp = """20
            0 ab
            6 cd
            0 ef
            6 gh
            4 ij
            0 ab
            6 cd
            0 ef
            6 gh
            0 ij
            4 that
            3 be
            0 to
            1 be
            5 question
            1 or
            2 not
            4 is
            2 to
            4 the"""
    # arr2 = []
    # for _ in range(inp):
    #     arr.append(input().rstrip().split())
    # print(arr2)
    print(counting_sort(n, arr))