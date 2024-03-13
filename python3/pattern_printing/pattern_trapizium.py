def print_pattern(rows):
    for i in range(1, rows + 1):
        # Print spaces before the stars
        print(" " * (rows - i), end="")

        # Print stars
        print("* " * (2 * i - 1))

    for i in range(rows - 1, 0, -1):
        # Print spaces before the stars
        print(" " * (rows - i), end="")

        # Print stars
        print("* " * (2 * i - 1))

# Example usage:
num_rows = 5
print_pattern(num_rows)


'''
    * 
   * * * 
  * * * * * 
 * * * * * * * 
* * * * * * * * * 
 * * * * * * * 
  * * * * * 
   * * * 
    * 
'''
