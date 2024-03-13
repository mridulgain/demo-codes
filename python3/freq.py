import timeit
t = timeit.timeit(
'''
n = [1,1,2,3,4,5,16,6,0,-99,0,1,2,6,3]
freq = {}
for i in n:
    if i not in freq:
        freq[i] = 1
    else:
        freq[i] += 1
''', number=10)
print(t)
