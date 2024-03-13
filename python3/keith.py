def break_digits(n):
    dig = []
    while(n > 0):
        d = n % 10
        n //= 10
        dig.insert(0, d)
    return dig

def isKeith(n):
    # breaking the digits
    digits = break_digits(n)
    # print(digits)
    # sum of 'd' digits, where 'd' is no of digits in 'n'
    d = len(digits)
    while True:
        # generate next term in the series
        total = 0
        for i in range(1, d+1):
            total += digits[-i]
        # print(digits, total)
        if total == n:
            return True
        elif total > n:
            return False
        else:
            # include next term in the series
            digits.append(total)
            #digits.pop(0)

def showKeith(d):
    if d < 2:
        print("Calculate for 2 digits or more..")
        return
    start = 10 ** (d-1) 
    # start = int("1" + "0" * (d-1))
    end = 10 ** d
    # end = int("1" + "0" * d)
    for i in range(start, end):
        if isKeith(i):
            print(i)

if __name__ == '__main__':
    # print(isKeith(456))
    # print(isKeith(47))
    # print(isKeith(197))
    # print(isKeith(1104))
    # print(isKeith(159))
    showKeith(2)

