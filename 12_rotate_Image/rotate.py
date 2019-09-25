# Rotate the n x n 2D image matrix by 90 degrees (clockwise)
# Asked By Apple, Microsoft, Amazon
# in Python
def rotateImage(array):
	n = len(array) # n x n Matrix so ROW == COLUMN
	rotatedImg = []

	for row in range(n):
		# Init. the sub Array for each cell
		rotatedImg.append([])

		for col in range(n):
			# Enter the value to their new corresponding indices
			oldCol = row         # old Col to grab from
			oldRow = n - 1 - col # old Row to grab from
			rotatedImg[row].append(array[oldRow][oldCol])

	return rotatedImg

matrix = [[1, 2, 3],
     [4, 5, 6],
     [7, 8, 9]]

print(rotateImage(matrix))
# Output
#     [[7, 4, 1],
#      [8, 5, 2],
#      [9, 6, 3]]