n = int(input("Enter an integer: "))
bin = 0
multiplier = 1
while n != 0:
    bin += (n % 2) * multiplier
    n //= 2
    multiplier *= 10
print(bin)
