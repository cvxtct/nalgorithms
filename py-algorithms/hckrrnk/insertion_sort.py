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

def insertionSort2(n, arr):
    # first is sorted
    # check 1 against 0
    if arr[1] > arr[0]:
        pass
    # else:
    #     arr[1] = arr[0]

    for i in range(len(arr)):
        if i == 0:
            continue
        #print(i)
       # while j != 0:
        counter = 0
        for j in reversed(range(i)):
            if arr[i-counter] < arr[j]: 
                tmp = arr[j]
                arr[j] = arr[i-counter]
                arr[i-counter] = tmp  
            counter += 1
     
        print(' '.join(map(str, arr)))
    
 
if __name__ == '__main__':
    # insertionSort1(14, [1, 3, 5, 9, 13, 22, 27, 35, 46, 51, 55, 83, 87, 23])
    insertionSort1(5, [1,2,4,5,3])
    insertionSort2(8, [8,7,6,5,4,3,2,1])