# write in a file
f = open("new.txt", "w")
# for i in range(100):
#     f.write(f"line {i*10} :\n")
name = input("enter your name please: ")
standard = input("Enter your class: ")
school = input("What's the name of your school? ") 
# print(name, standard, school)
tmp = f"Name: {name}\n\
Class: {standard}\n\
School: {school}"
f.writelines(tmp)
f.close()