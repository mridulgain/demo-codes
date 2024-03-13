import re

inp = '''
password # invalid
pass0rD # length < 8
passw0rd # no caps
PASSW0RD # no small
PassworD # no digits
Passw0rD # valid
passW0rd100 #valid
'''

passw = re.compile(r'\w{8,}')
caps = re.compile(r'[A-Z]')
small = re.compile(r'[a-z]')
digit = re.compile(r'\d')

results = passw.findall(inp)
#print(result)
for p in results:
    hasCaps = caps.search(p)
    hasSmall = small.search(p)
    hasDigit = digit.search(p)
    if hasCaps and hasSmall and hasDigit:
        print(p)
