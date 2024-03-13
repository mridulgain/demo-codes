from pgm_reader import Reader
import matplotlib.pyplot as plt
import numpy as np
f = "out.pgm"
reader = Reader()
image = reader.read_pgm(f)
print(image)
print("width = ", reader.width)
plt.imshow(image, cmap='gray')
image = image.transpose()

# reader.show_img()
def write_image(mat, fnam):
    f = open(fnam, "w")
    f.write("P2\n")
    f.write(f'{len(mat[0])} {len(mat)}\n')
    f.write(str(255) + "\n")
    for row in mat:
        for pixel in row:
            f.write(str(pixel) + "\n")
    f.close()
    print("Output: ", fnam)

# write_image(result, "out1.pgm")