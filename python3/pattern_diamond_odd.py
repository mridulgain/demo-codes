def print_star_pattern(rows):
    # Print the upper part of the pattern
    for i in range(1, rows + 1, 2):
        spaces = (rows - i) // 2
        stars = '*' * i
        print(' ' * spaces + stars)

    # Print the lower part of the pattern
    for i in range(rows - 2, 0, -2):
        spaces = (rows - i) // 2
        stars = '*' * i
        print(' ' * spaces + stars)

# Example usage:
num_rows = 5
print_star_pattern(num_rows)

'''
  *  
 ***
*****
 ***
  * 
'''
