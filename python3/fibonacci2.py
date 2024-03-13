def fib(n):
    global e
    if len(e) < n:
        e += [e[-1] + e[-2]]
        fib(n)


n = int(input("How many numbers ? "))
e = [1, 1]
while len(e) < n:
    e += [e[-1] + e[-2]]
#fib(n)
print(e)
