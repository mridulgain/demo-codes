# Matrix operations
def input_matrix(row, col):
    # matrix initialization
    # m = [[None for i in range(col)] for j in range(row)]
    m = []
    for i in range(row):
        m.append([None] * col)
    # input the elements
    print("Enter members of the array: ")
    for i in range(row):
        for j in range(col):
            m[i][j] = int(input(f"Enter value at [{i}][{j}]: "))
    return m

def print_matrix(m):
    row = len(m)
    col = len(m[0])
    # print the elements
    for i in range(row):
        print(" | ", end="")
        for j in range(col):
            print(m[i][j], end=" | ")
        print("")

def add_matrix(p, q):
    row = len(p)
    col = len(p[0])
    result = []
    for i in range(row):
        result.append([None] * col)
    # addition
    for i in range(row):
        for j in range(col):
            result[i][j] = p[i][j] + q[i][j]
    return result

def subtract_matrix(p, q):
    row = len(p)
    col = len(p[0])
    result = []
    for i in range(row):
        result.append([None] * col)
    # addition
    for i in range(row):
        for j in range(col):
            result[i][j] = p[i][j] - q[i][j]
    return result

# driver
row = int(input("Enter the no. of row: "))
col = int(input("Enter the no. of col: "))

print("Enter 1st matrix: ")
x = input_matrix(row, col)
print("Enter 2nd matrix: ")
y = input_matrix(row, col)
z = add_matrix(x, y)
print("Matrix1--------")
print_matrix(x)
print("Matrix2---------")
print_matrix(y)
print("Add-------")
print_matrix(z)
print("Subtract-------")
w = subtract_matrix(x, y)
print_matrix(w)