# input: list, row, col
# output: matrix[row][col]
n = [i for i in range(9)]
row = 3
col = 3
print(n)
print("")
if len(n) != row * col:
    print("Invalid input")
else:
    m = []
    # method 1
    for i in range(row):
        tmp = []
        for j in range(col):
            tmp += [ n[i * col + j] ]
        m += [tmp]
    # method 2
    m1 = []
    tmp = []
    for i, x in enumerate(n):
        tmp += [x]
        if len(tmp) == col:
            m1 += [tmp]
            tmp = []
    # method 3
    m2 = []
    x = 0
    for i in range(row):
        m2.append(n[x:x+col])
        x += col
    # output
    print("Method1: ")
    for i in m:
        print(i)
    #
    print("Method2: ")
    for i in m1:
        print(i)
    #
    print("Method3: ")
    for i in m2:
        print(i)
