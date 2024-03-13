#input
n1 = int(input("Enter a numberr "))
n2 = int(input("Enter another number "))
if n2 < n1:
    n1 , n2 = n2, n1

# calculation
def gcd_definition(n1, n2):
    # GCD by definition
    gcd = 1
    for i in range(1, n1+1):
        if (n1 % i == 0) and (n2 % i == 0):
            gcd = i
    return gcd

def gcd_division(n1, n2):
    while(n1 != 0 and n2 % n1 != 0):
        n1, n2 = n2 % n1, n1
        print(n1, n2)
    return n1

def primeFactors(n):
    factors = []
    while n % 2 == 0:
        factors.append(2)
        n = n / 2

    for i in range(3, int(n) + 1, 2):
        while n % i == 0:
            factors.append(i)
            n = n / i
    
    if n > 2:
        factors.append(n)
    return factors

def gcd_factor(n1, n2):
    pf_1 = primeFactors(n1)
    pf_2 = primeFactors(n2)

    # for i in pf_1:
    #     if i in pf_2:
    #         unique.append(i)

    gcd = 1
    for i in pf_1:
        if i in pf_2:
            gcd *= i
            pf_2.remove(i)
    print(f"GCD of {n1} and {n2} is: {gcd}")

# output
print(f"GCD of {n1} & {n2} is..")
gcd = gcd_definition(n1, n2)
print(f"Using definition: {gcd}")
gcd = gcd_division(n1, n2)
print(f"Using division method: {gcd}")
gcd = gcd_factor(n1, n2)
print(f"Using factor method: {gcd}")


