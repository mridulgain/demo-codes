# top down, right to left
lines = int(input("number of rows? "))
for i in range(1, lines + 1): # 1 to lines
    # single line
    for j in range(lines - i):
        print(" ", end = "")
    for k in range(i):
        print("+", end = "")
    print(" ") #new line

'''
number of rows? 4
   + 
  ++ 
 +++ 
++++ 
'''
