def find_median(arr):
    arr.sort()
    # even
    if len(arr) % 2 == 0:
        return arr[int(len(arr)/2) - 1] + arr[int(len(arr) / 2)]
    # odd
    if len(arr) % 2 == 1:
        return arr[int(len(arr)/2)]

if __name__ == '__main__':
    arr = [5, 3, 1, 2, 4, 6, 7, 8]
    print(find_median(arr))