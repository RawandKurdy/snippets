# sudoku2 Interview Question Asked by Apple & Uber
# in Python
def sudoku2(grid):
    trackColNumbers = {} # Numbers in each column list
    trackRowNumbers = {} # Numbers in each row list
    currentSubMatrix = {} # Numbers in Each Sub Matrix (3x3)
    sub_Size = len(grid)//3 # Sub group sizes, here its 9/3 = 3
    
    col_add = 0 # Added to the sub matrix col to get the real column
    row_add = 0 # Added to the sub matrix row to get the real row
        
    while row_add != 9:
        for sub_col in range(sub_Size):
            real_Col = sub_col+col_add # real column in the grid
            if trackColNumbers.get(real_Col) is None:
                trackColNumbers[real_Col] = {}

            for sub_row in range(sub_Size):
                real_Row = sub_row+row_add # real row in the grid
                if trackRowNumbers.get(real_Row) is None:
                    trackRowNumbers[real_Row] = {}

                item = grid[real_Row][real_Col] # each item in the cells
                    
                if (item is not '.'):
                    if currentSubMatrix.get(item) is None:
                        currentSubMatrix[item] = 0
                    else : return False
                        
                    if trackRowNumbers[real_Row].get(item) is None:
                        trackRowNumbers[real_Row][item] = 0
                    else : return False
                        
                    if trackColNumbers[real_Col].get(item) is None:
                        trackColNumbers[real_Col][item] = 0
                    else : return False
                
        currentSubMatrix = {} # Reset Sub Matrix each time
        col_add+=3 # moves to the next 3 columns
        if (col_add==9):
            col_add = 0 # reset back columns to start
            row_add += 3 # moves to the next 3 rows
                    
    return True

grid = [['.', '.', '.', '.', '2', '.', '.', '9', '.'],
        ['.', '.', '.', '.', '6', '.', '.', '.', '.'],
        ['7', '1', '.', '.', '7', '5', '.', '.', '.'],
        ['.', '7', '.', '.', '.', '.', '.', '.', '.'],
        ['.', '.', '.', '.', '8', '3', '.', '.', '.'],
        ['.', '.', '8', '.', '.', '7', '.', '6', '.'],
        ['.', '.', '.', '.', '.', '2', '.', '.', '.'],
        ['.', '1', '.', '2', '.', '.', '.', '.', '.'],
        ['.', '2', '.', '.', '3', '.', '.', '.', '.']]

print("is Grid Valid?", sudoku2(grid)) # answer :> False