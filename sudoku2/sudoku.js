// sudoku2 Interview Question Asked by Apple & Uber
// in Javascript
function sudoku2(grid) {
    let trackColNumbers = {}; // Numbers in each column list
    let trackRowNumbers = {}; // Numbers in each row list
    let currentSubMatrix = {}; // Numbers in Each Sub Matrix (3x3)
    const sub_Size = grid.length / 3; // Sub group sizes, here its 9/3 = 3

    let col_add = 0; // Added to the sub matrix col to get the real column
    let row_add = 0; // Added to the sub matrix row to get the real row

    while (row_add != 9) {
        for (let sub_col = 0; sub_col < sub_Size; sub_col++) {
            let real_Col = sub_col + col_add; // real column in the grid
            if (trackColNumbers[real_Col] == null)
                trackColNumbers[real_Col] = {};

            for (let sub_row = 0; sub_row < sub_Size; sub_row++) {
                let real_Row = sub_row + row_add; // real row in the grid
                if (trackRowNumbers[real_Row] == null)
                    trackRowNumbers[real_Row] = {};

                let item = grid[real_Row][real_Col] // each item in the cells

                if (item != '.') {
                    if (currentSubMatrix[item] == null)
                        currentSubMatrix[item] = 0;
                    else return false;

                    if (trackRowNumbers[real_Row][item] == null)
                        trackRowNumbers[real_Row][item] = 0;
                    else return false;

                    if (trackColNumbers[real_Col][item] == null)
                        trackColNumbers[real_Col][item] = 0;
                    else return false;
                }
            }
        }
        currentSubMatrix = {}; // Reset Sub Matrix each time
        col_add += 3;  // moves to the next 3 columns
        if (col_add == 9) {
            col_add = 0; // reset back columns to start
            row_add += 3; // moves to the next 3 rows
        }
    }
    return true;
}

let grid = [['.', '.', '.', '.', '2', '.', '.', '9', '.'],
['.', '.', '.', '.', '6', '.', '.', '.', '.'],
['7', '1', '.', '.', '7', '5', '.', '.', '.'],
['.', '7', '.', '.', '.', '.', '.', '.', '.'],
['.', '.', '.', '.', '8', '3', '.', '.', '.'],
['.', '.', '8', '.', '.', '7', '.', '6', '.'],
['.', '.', '.', '.', '.', '2', '.', '.', '.'],
['.', '1', '.', '2', '.', '.', '.', '.', '.'],
['.', '2', '.', '.', '3', '.', '.', '.', '.']];

console.log("is Grid Valid?", sudoku2(grid)) // answer :> False