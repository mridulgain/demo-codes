txt = input("Enter a text to count vowel: ")
if txt.isalpha():
    vowels = "aeiou"
    vowel_count = 0

    for i in txt:
      if i in vowels:
        vowel_count += 1
        #print(vowel_count, ":", i)

    print("No of vowels", vowel_count)
    print("No of consonanats", len(txt) - vowel_count)
else:
    print("invalid input")
