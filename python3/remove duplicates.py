def remove_duplicate(l):
    unique = []
    for i in l:
      if i not in unique:
        unique.append(i)
    return unique

list = []
print('''To stop just enter stop or 'stop' or "stop"''')
i = 0
while True:
    i += 1
    user_input = input(f"Enter integer {i}: ")
    if user_input == "stop" or user_input == "'stop'" or user_input == '''"stop"''':
        break
    list.append(int(user_input))
print("The list is", list)
print("After removing duplicates", remove_duplicate(list))

#list = [12,1,0,0,-1,12,0,5]



# l1 = [12, 33, 2, 2, 3, 5, 7, 8] 
# u1 = remove_duplicate(l1)
# print(l1)
# print(u1)

# l2 = [12,1,0,0,-1,12,0,5]
# u2 = remove_duplicate(l2)
# print(l2, ">", u2)
