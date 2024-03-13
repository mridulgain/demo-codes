import sys

# def word_freq(string):
#    string = string.lower().split()
#    list = []
#    for i in string:
#       print(i)
#       if i not in list:
#          list.append(i)
#       if not i.isalpha():
#          i.replace(i, "")
            
#    for i in range(0, len(list)):
#       print('Frequency of', list[i], 'is :', string.count(list[i]))

def word_freq(sentence):
   words = sentence.lower().split()
   freq = {}
   for word in words:
      # remove non-letters from the end
      while len(word) > 0 and not word[-1].isalpha():
         word = word[:-1]
      # remove non-letters from the begining
      while len(word) > 0 and not word[0].isalpha():
         word = word[1:]
      if len(word) > 0:
         if word not in freq:
            freq[word] = 1
         else:
            freq[word] += 1
   
   print(freq)



def main():
   string = input("Enter a string: ")
   word_freq(string)

while True:
   choice = input("Do you want to continue or not? Yes or No: ")
   if choice == "y" or choice == "yes":
      main()
   elif choice == "n" or choice == "no":
      sys.exit()
   else:
      pass