'''
4321
321
21
1
'''
row = 4
for i in range(row):
    for j in range(row-i,0,-1):
        print(j, end="")
    print()