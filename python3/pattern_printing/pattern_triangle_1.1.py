'''output
*
**
***
****
'''
row = 5
for i in range(row):
    for j in range(i+1):
        print("*", end="")
    print()