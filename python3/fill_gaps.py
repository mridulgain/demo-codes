import os
import random
import re

def fill_gaps(folderName, prefix):
    dir_contents = os.listdir(folderName)
    print("Looking for files starting with ", prefix)
    pattern = "^" + prefix + "(\d+)(.*)"
    d = {}
    start = 1000
    for fileName in sorted(dir_contents):
        search_result = re.search(pattern, fileName)
        if search_result != None:
            num = int(search_result.group(1))
            if num < start:
                start = num
            print(num, ": ", fileName)
            d[num] = fileName
    print(d)
    for i in sorted(d):
        search_result = re.search(pattern, fileName)
        suffix = search_result.group(2)
        new_name = re.sub(pattern, prefix + str(start) + suffix, d[i])
        start += 1
        print("Renaming", d[i], "to", new_name)
    
    
fill_gaps(".", "logs_")
