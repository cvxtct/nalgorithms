

def quick_sort(n, arr):
    pivot = arr[0]
    e = []
    l = []
    r = []

    for i in range(n):
        if arr[i] < pivot:
            l.append(arr[i])
            print(l)
        if arr[i] > pivot:
            r.append(arr[i])
            print(r)
        if arr[i] == pivot:
            e.append(arr[i])
   
    res = l + e + r
    
    return res

if __name__ == '__main__':
    print(quick_sort(5, [4,5,3,7,2]))




