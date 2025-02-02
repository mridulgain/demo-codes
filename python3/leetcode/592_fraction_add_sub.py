import re
'''
Given a string expression representing an expression of fraction addition and subtraction, return the calculation result in string format.

The final result should be an irreducible fraction. If your final result is an integer, change it to the format of a fraction that has a denominator 1. So in this case, 2 should be converted to 2/1.

Constraints:

    The input string only contains '0' to '9', '/', '+' and '-'. So does the output.
    Each fraction (input and output) has the format ±numerator/denominator. If the first input fraction or the output is positive, then '+' will be omitted.
    The input only contains valid irreducible fractions, where the numerator and denominator of each fraction will always be in the range [1, 10]. If the denominator is 1, it means this fraction is actually an integer in a fraction format defined above.
    The number of given fractions will be in the range [1, 10].
    The numerator and denominator of the final result are guaranteed to be valid and in the range of 32-bit int.

'''

class Fraction:
    def __init__(self, sign, numerator, denominator):
        multiplier = 1
        if sign == '-':
            multiplier = -1
        self.numerator = multiplier*int(numerator)
        self.denominator = int(denominator)
        
    def reduce(self):
        a = abs(self.denominator)
        b = abs(self.numerator)
        while b > 0:
            a, b = b, a % b
        gcf = a
        self.denominator //= gcf
        self.numerator //= gcf

    def __str__(self) -> str:
        s = f"{self.numerator}/{self.denominator}"
        return s

    def __add__(x, y):
        z_denominator = x.denominator * y.denominator
        z_numerator = y.denominator * x.numerator + x.denominator * y.numerator
        z = Fraction('', z_numerator, z_denominator)
        z.reduce()
        return z

def compute(input):
    regx = re.compile(r'([+-]?)(\d+)/(\d+)')
    results = regx.findall(input)
    print(results)

    sum = Fraction('', '0', '1')
    for i in results:
        sum += Fraction(*i)
    return sum

if __name__ == "__main__":
    # tc1
    input = "-1/2+1/2"
    result = compute(input)
    print(result)
    assert str(result) == '0/1'
    # tc2
    input = "1/2+1/2"
    result = compute(input)
    print(result)
    assert str(result) == '1/1'
    # tc3
    input = "-1/2+1/2+1/3"
    result = compute(input)
    print(result)
    assert str(result) == '1/3'
