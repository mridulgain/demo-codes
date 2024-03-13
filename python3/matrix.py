# a = [
#     [1,2,3],
#     [4,5,6],
#     [7,8,9]
# ]
def input_matrix(row, col):
    a = [[0 for i in range(col)] for j in range(row)]
    for i in range(row):
        for j in range(col):
            a[i][j] = input(f"{i},{j}? ")
    return a

def print_matrix(mat):
    row = len(mat)
    col = len(mat[0])
    for i in range(row):
        for j in range(col):
            print(mat[i][j], end=" ")
        print("")

row = int(input("row? "))
col = int(input("col? " ))
print("--------------------")
a = input_matrix(row, col)
print("--------------------")
print_matrix(a)
 
#

# itr 1
# i = 0
# j = 0
# print(i, j)
# j = 1
# print(i, j)
# j = 2
# print(i, j)
# # itr 2
# i = 1
# j = 0
# print(i, j)
# j = 1
# print(i, j)
# j = 2
# print(i, j)
# # itr 2
# i = 2
# j = 0
# print(i, j)
# j = 1
# print(i, j)
# j = 2
# print(i, j)

