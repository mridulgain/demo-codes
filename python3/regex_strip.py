import re

inp = "   hello world   "
# regex = re.compile(r'\s*(.*\S)\s*')
# output = regex.search(inp)
# print(output.group(1))

inp2 = "~~~~hello~~~world>~~~~~"
# regex = re.compile(r'~*(.*[^~])~*')
# output = regex.search(inp2)
# print(output.group(1))

def regex_strip(inp, rem = None):
    if rem is not None:
        tmp = rem + '*(.*[^' + rem + '])' + rem + '*'
    else:
        tmp = r'\s*(.*\S)\s*'
    regex = re.compile(tmp)
    output = regex.search(inp)
    return output.group(1)

print(regex_strip(inp))
print(regex_strip(inp2, "~"))
inp3 = "----Hi remove all '-' except these------"
out = regex_strip(inp3, "-")
print(out)