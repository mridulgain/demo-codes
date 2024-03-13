# left to right top down
lines = int(input("rows? "))

for i in range(1,lines+1):
    # for each line
    for j in range(i):
        print("*", end="")
    print("")

'''
rows? 4
*
**
***
****
'''
