import filter_dir_contents as fdc
import os
import random
import re

def fill_gaps(folderName, prefix):
    pattern = "^" + prefix + "(\d+)(.*)"
    filtered = fdc.filter_dir_contents(folderName, prefix)
    print(filtered)
    # rename
    current_num = 1
    for f in sorted(filtered):
        search_result = re.search(pattern, filtered[f])
        suffix = search_result.group(2)
        new_file_name = re.sub(pattern, prefix + str(current_num) + suffix, filtered[f])
        current_num += 1
        print("Renaming", filtered[f], "into", new_file_name)


fill_gaps(".", "logs_")

