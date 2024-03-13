import os
import random
import re

def add_gaps(folderName, prefix, gapSize):
    dir_contents = os.listdir(folderName)
    print("Looking for files starting with ", prefix)
    pattern = "^" + prefix + "(\d+).*"
    d = {}
    start = end = 0
    for fileName in sorted(dir_contents):
        search_result = re.search(pattern, fileName)
        if search_result != None:
            print(search_result.group(1), ": ", fileName)
            d[int(search_result.group(1))] = fileName
    # choosing randomly
    print(d.values())
    files_to_be_deleted = random.sample(list(d.values()), gapSize)
    for f in files_to_be_deleted:
        print("Deleting ", f)
        os.unlink(os.path.join(folderName, f))


add_gaps('.', 'logs_', 3)
