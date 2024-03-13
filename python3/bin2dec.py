bin = int(input("Enter a binary number: "))
dec = 0
multiplier = 1
while bin > 0:
    digit = bin % 10
    dec += digit * multiplier
    multiplier *= 2
    # dec = dec * 2 + digit
    bin //= 10
print(dec) 