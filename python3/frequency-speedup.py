import time
start = time.time()
n = [1,1,2,3,4,5,16,6,0,-99,0,1,2,6,3]
freq = {}
for i in n:
    if i not in freq:
        freq[i] = 1
    else:
        freq[i] += 1
        
#print(freq)
print("Member\tFrequency")
for i in freq:
    print(i, "\t", freq[i])

end = time.time()
print("Time taken: ", end-start)
