import re as r
phoneRegx = r.compile(r'\+\d{2}-\d{10}')
searchString = \
'''
First Name: Mridul
Last Name: Gain
Phone: +91-9051760152

'''
# find first occurance
#searchResult = phoneRegx.search(searchString)
#print(searchResult.group())

# find all occurance
#searchResult = phoneRegx.findall(searchString)
#print(searchResult)

# find groups
# phoneRegx = r.compile(r'(\+\d{2})-(\d{10})')
# searchResult = phoneRegx.findall(searchString)
# print(searchResult)

# pipe
# regx = r.compile(r'(Bat|Cat)man')
# searchString = \
# '''
# Catman
# Batman

# Iceman
# '''
# searchResult = regx.search(searchString)
# print(searchResult.group())

# optional matching
# phoneRegx = r.compile(r'(\+\d{2}-)?\d{10}')
# searchString = \
# '''

# abhijai: 9876543201
# '''
# searchResult = phoneRegx.search(searchString)
# print(searchResult.group())

# repeatitions
regx = r.compile(r'(ha){3}')
