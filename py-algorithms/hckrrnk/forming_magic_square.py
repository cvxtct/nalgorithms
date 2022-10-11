def forming_magic_square(s):
    sum_cost = 0
    print(s)
    # make sure middle is 5
    # if s[1][1] != 5:
    #     s[1][1] = 5
    
    # for i in range(0, len(s)):
    #     line_cost = 0
    #     for j in range(0, len(s[i])):
    #         sum_line = sum(s[i])
    #         if sum_line != 15:
    #             if i in [0,2] and j in [0,2]:
    #                 if s[i][j] % 2 == 1:                        
    #                     s[i][j] = 15 - sum(s[i]) + s[i][j]    
    #             if (i == 0 and j == 1) or (i == 1 and j in [0,2]) or (i == 2 and j == 1):
    #                 if s[i][j] % 2 == 0:                        
    #                     s[i][j] = (15 - sum(s[i])) + s[i][j]
    
    n = [s[i][j] for i in range(3) for j in range(3)]
    print(n)
    all_n = [
        [8, 1, 6, 3, 5, 7, 4, 9, 2],
        [6, 1, 8, 7, 5, 3, 2, 9, 4],
        [4, 9, 2, 3, 5, 7, 8, 1, 6],
        [2, 9, 4, 7, 5, 3, 6, 1, 8],
        [8, 3, 4, 1, 5, 9, 6, 7, 2],
        [4, 3, 8, 9, 5, 1, 2, 7, 6],
        [6, 7, 2, 1, 5, 9, 8, 3, 4],
        [2, 7, 6, 9, 5, 1, 4, 3, 8]
    ]

    allsum = []
    for l in all_n:
        allsum.append(sum([abs(n[i]-l[i]) for i in range(9)]))    
    print(min(allsum))
                        
                
    # for _, a in enumerate(s):
    #     for _, b in enumerate(a):
    #         print(b)

if __name__ == '__main__':
    s = [[5, 3, 4], [1, 5, 8], [6, 4, 2]]
    forming_magic_square(s)