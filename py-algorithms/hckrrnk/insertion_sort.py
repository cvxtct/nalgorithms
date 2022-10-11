from audioop import reverse


def insertionSort1(n, arr):
    # Write your code here
    target = arr[-1]
    idx = n-2
    while (target < arr[idx] and idx >= 0):
        arr[idx+1] = arr[idx]
        print(' '.join(map(str, arr)))
        idx -= 1

    arr[idx+1] = target
    print(' '.join(map(str, arr)))
    # small = arr[n-1]
    # while n != 0:
    #     print(arr[n-1])
    #     if arr[n-1] >= small:
    #         arr[n-1] = arr[n-2] 
    #     print(' '.join(map(str, arr)))
    #     n = n-1

    # for i in reversed((range(n))):
    #     arr[n-1] = arr[n-2]
    #     if arr[i] > small:
    #         arr[i] = arr[i-1]
    #         print(' '.join(map(str, arr)))
    #     if arr[i] < small:
  
    #         arr[i] = small
    #         print(' '.join(map(str, arr)))
    #     if small in arr:
    #         break
        

if __name__ == '__main__':
    #insertionSort1(14, [1, 3, 5, 9, 13, 22, 27, 35, 46, 51, 55, 83, 87, 23])
    insertionSort1(5, [1,2,4,5,3])