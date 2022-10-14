def insertion_sort(l):
    for i in range(0, len(l)):
      
        j = i-1
        key = l[i]
        counter = 0
        while (j >= 0) and (l[j] > key):
           print(l[i],l[j])
           l[j+1] = l[j]
           j -= 1
        l[j+1] = key

ar = [7 ,4, 3, 5, 6, 2]
insertion_sort(ar)
print(" ".join(map(str,ar)))