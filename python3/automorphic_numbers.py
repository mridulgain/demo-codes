n = input("Enter an int: ")

sq = int(n) ** 2
print(n, "^ 2 = ", sq)
if n == str(sq)[-len(n):]:
    print("automorphic")
else:
    print("not automorphic")
