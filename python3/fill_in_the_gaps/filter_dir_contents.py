import os, re

def filter_dir_contents(folderName, prefix):
    # list dir contents
    dir_contents = os.listdir(folderName)
    # prepare regex for (logs_)(19)(.txt) 
    pattern = "^" + prefix + "(\d+)(.*)"
    # filter using the pattern
    filtered = {}
    for fileName in dir_contents:
        search_result = re.search(pattern, fileName)
        if search_result != None:
            num = int(search_result.group(1))
            print(num, ":", fileName)
            filtered[num] = fileName
    return filtered

if __name__ == '__main__':
    print(filter_dir_contents('.', 'logs_'))