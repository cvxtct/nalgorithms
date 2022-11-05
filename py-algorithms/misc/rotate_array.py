def rotate_array__apr_one(array_, d):
    # to the left
    # k = array_.index(d)
    temp_ = array_[d:]
    temp_ = temp_ + array_[:d]
    for a,t in zip(array_, temp_):
        print(a,t)
        array_[a-1] = t
    print(array_)

def rotate_array__apr_two(array_, d, n):
    p = 1

    while p <= d:
        last = array_[0]
        for i in range(n-1):
            array_[i] = array_[i+1]
        array_[n - 1] = last
        p+=1
    print(array_)



if __name__ == '__main__':
    array_ = [1,2,3,4,5,6,7]
    d = 3
    n = len(array_)
    #rotate_array__apr_one(array_, d)
    rotate_array__apr_two(array_, d, n)
