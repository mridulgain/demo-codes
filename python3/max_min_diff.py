# input
n = []
tmp = None
while True:
    tmp = input("? ")
    if tmp == 'stop':
        break
    n += [int(tmp)]
print(n)
# process
max_ = -99999
min_ = +99999
max_pair = None
min_pair = None
for i in range(len(n)):
    for j in range(i+1, len(n)):
        diff = abs(n[i] - n[j])
        pair = (n[i], n[j])
        if diff > max_:
            max_ = diff
            max_pair = pair
        if diff < min_:
            min_ = diff
            min_pair = pair
# output
print(f"Max diffrence is {max_} for {max_pair}")
print(f"Min diffrence is {min_} for {min_pair}")
