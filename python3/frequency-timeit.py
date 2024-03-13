import timeit
t= timeit.timeit(
'''
n = [1,1,2,3,4,5,16,6,0,-99,0,1,2,6,3]

def remove_duplicates(l):
    unique = []
    for i in l:
      if i not in unique:
        unique.append(i)
    return unique
    
unique = remove_duplicates(n)
#print("Member\tFrequency")
for i in unique:
    #print(i, "\t", n.count(i))
    n.count(i)
''', number=10)
print(t)
