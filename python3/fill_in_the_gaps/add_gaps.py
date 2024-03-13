import os
import random
import re
import filter_dir_contents as fdc

def add_gaps(folderName, prefix, gapSize):
    pattern = "^" + prefix + "(\d+)(.*)"
    filtered = fdc.filter_dir_contents(folderName, prefix)
    print(filtered)
    # choose some files randomly to delete
    files_to_del = random.sample(list(filtered.values()), gapSize)
    print("Delete ", files_to_del)
    for f in files_to_del:
        os.unlink(os.path.join(folderName,f))



add_gaps('.', 'logs_', 3)