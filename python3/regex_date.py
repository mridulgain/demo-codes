import re

str = ''' 
12/12/2021 #valid 
1990/12/12 # not valid 
32/09/2000 #day out of range 
12/13/1991 # month out of range 
02/11/3000 # year out of range 
1/1/1992 # not in dd/mm/yyyy format 
29/02/2020 # leap year validation 
29/03/999 # year out of range
'''
 
valid_date = re.compile(r'''(
([0-2]\d|3[0-1]) # day
/
(0\d|1[0-2])     # month
/
([1-2]\d{3})     # year
)''', re.VERBOSE)

result = valid_date.findall(str)
print(result)
