def bigSorting(unsorted):
    for v in sorted(unsorted, key = lambda v: int(v)):
        print(v)

if __name__ == '__main__':
        u = ['31415926535897932384626433832795','1','3','10','3','5']
        bigSorting(u)