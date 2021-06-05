keys = ['Ten', 'Twenty', 'Thirty']
values = [10, 20, 30]
# d = {keys[0]: values[0]...}
d = {}
print("-----loop1----")
for i, j in zip(keys, values):
    # d[i] = j
    print(i, j)

print("-----loop1----")
for i in range(len(keys)):
    d[keys[i]] = values[i]
    #print(keys[i], values[i])

print(keys)
print(values)
print(d)
