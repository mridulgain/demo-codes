f = open("new.txt", "r")

for line in f.readlines():
    print(line, end="")