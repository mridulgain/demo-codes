def factorial_loop(a):
    ans = 1
    for i in range(1, a+1):
        ans *= i
    return ans

def factorial_rec(a):
    if a == 0:
        return 1
    else:
        return a * factorial_rec(a-1)
    
    
n = 10
ans = factorial_loop(n)
print(f"Factorial of {n}(calculated using loop) is {ans}")
ans = factorial_rec(n)
print(f"Factorial of {n}(with recursion) is {ans}")
