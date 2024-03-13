import pickle
list_student = []
f = open("exams.csv", "r")
for line in f:
    student, maths, reading, writing = line.split(",")
    d = {
        "Student": student,
         "Maths": maths,
         "Reading": reading,
         "Writing": writing
        }
    '''
    d["Student"] = student
    d["Maths"] = maths
    d["Reading"] = reading
    d["Writing"] = writing
    '''
    list_student.append(d)

#print(list_student)
fp = open("exams.bin", "wb")
pickle.dump(list_student, fp)
print("Data stored in exams.bin")
