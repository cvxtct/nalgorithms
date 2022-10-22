def insertion_sort(l):
    shifts = 0
    for i in range(0, len(l)):
      
        j = i-1
        key = l[i]
        counter = 0
        while (j >= 0) and (l[j] > key):
           shifts += 1
           l[j+1] = l[j]
           j -= 1
        l[j+1] = key
    # print(shifts)
# arr = [7 ,4, 3, 5, 6, 2]
arr = [2, 1 ,3 ,1 ,2]
insertion_sort(arr)
print(" ".join(map(str,arr)))
print(map(str, arr).__dir__())