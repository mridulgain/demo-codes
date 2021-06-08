tableData = [['apples', 'oranges', 'cherries', 'banana'],
             ['Alice', 'Bob', 'Carol', 'David'],
             ['dogs', 'cats', 'moose', 'goose']]

def printTable(table):
    col_widths = getLongestWordLength(table)
    # print(col_widths)
    for i in col_widths:
        print(str(i).rjust(i), end=" ")
    print()
    for j in range(len(table)):
        for i in range(len(table[0])):
            print(table[j][i].rjust(col_widths[i]), end=" ")
        print()


def getLongestWordLength(table):
    return[max([len(row[i]) for row in table]) for i in range(len(table[0]))]
    # max_width = []
    # for i in range(len(table[0])):
    #     tmp = []
    #     for row in table:
    #         tmp.append(len(row[i]))
    #     max_width.append(max(tmp))
    # return max_width


printTable(tableData)