from typing import Any

def bouble(arr: list[Any]) ->  list[Any]:
    """
    >>> bouble([-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10])
    [-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10]
    """
    swapped = True
    while swapped:
        swapped = False
        for i in range(0, len(arr)-1):
            if arr[i+1] < arr[i]:
                arr[i+1], arr[i] = arr[i], arr[i+1]
                swapped = True
    
    return arr

if __name__ == '__main__':
    import doctest
    doctest.testmod()

                
