def fib(n):
    if n <= 0:
        print("Enter a valid input")
    elif n == 1 or n == 2:
        e[n] = 1
        return 1
    else:
        if not e[n-1]:
            e[n-1] = fib(n-1)
        if not e[n-2]:
            e[n-2] = fib(n-2)
        e[n] = e[n-1] + e[n-2]
        return e[n]


n = int(input("How many numbers ? "))
e = [None] * (n+1)
print(fib(n))
print(e[1:])
