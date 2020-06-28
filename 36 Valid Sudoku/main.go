package main

import (
	"fmt"
)

var example1 = [][]byte{
	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

var example2 = [][]byte{
	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

func isValidSudoku1(board [][]byte) bool {

	var senseArray [27][9]bool

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			val := int8(board[x][y] - 49)
			subBox := 3*(x/3) + (y / 3)
			if val == -3 {
				continue
			}
			if senseArray[x][val] || senseArray[y+9][val] || senseArray[subBox+18][val] {
				return false
			}
			senseArray[x][val] = true
			senseArray[y+9][val] = true
			senseArray[subBox+18][val] = true
		}
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {

	// The 9 Sudoku values ('1' to '9') times nine rows, nine columns, and 9 boxes yields 243 possible
	// bits needed if a set bit indicates that a particualar Sudoku value has been noted at a particular
	// position in one of the rows, one of the columns, and one of the boxes.  The largest bit type in
	// Go is a uint64 with 64 bits so four uint64s are needed to capture the 243 possible slots.
	//
	// We generate a number from 0 to 27 where 0 to 8 represent the nine rows, 9 to 15 represent the nine
	// columns, and 16 to 27 to represent the boxes ('pos').  We then multiply that number by 9 and then add the
	// Sudoku value ('val') to it.  This give us the bit offset ('bitOffset') from 0 to 242 We then use integer
	// div and mod on this number.  A mod by 4 gives us a 2-bit index selector from 0 to 3 ('index'). And, a the
	// integer div by 4 gives us a 64-bit offset in the selected word ('wordOffset').
	//
	// We can then use bit set and test manipulations on the uint64 array value at index 'index' to determine
	// if a Sudoku value has been seen.
	//
	// This approach uses 243 out of 256 bits to track Sudoku values (95% effeciency)

	var sense [4]uint64

	for row := uint8(0); row < 9; row++ {
		for col := uint8(0); col < 9; col++ {

			// Convert the UTF-8 byte to an integer and shift downward based upon the ASCII code set so that
			// the '1' glyph has the value of 0.  This will make the '.' glyph have a value of -3.  If we see
			// that value, then move on to the next location in the 9x9 Sudoku board.
			val := int(board[row][col]) - 49
			if val == -3 {
				continue
			}

			// Convert the current row, column, and value into bit offsets as described above
			rowBitOffset := uint8(9*row + uint8(val))                        // [0-80]
			colBitOffset := uint8(81 + 9*col + uint8(val))                   // [81-161]
			boxBitOffset := uint8(162 + 27*(row/3) + 9*(col/3) + uint8(val)) // [162-242]

			// Check these bits to see if they are currently set.  If they are, then we have seen the number before
			// in either a row, column, or box and therefore the Sudoku board is invalid
			if sense[rowBitOffset%4]&(0x01<<(rowBitOffset/4)) > 0 ||
				sense[colBitOffset%4]&(0x01<<(colBitOffset/4)) > 0 ||
				sense[boxBitOffset%4]&(0x01<<(boxBitOffset/4)) > 0 {
				return false
			}

			// Otherwise, we have not previously seen them so set the bits to true indicating that we now have
			sense[rowBitOffset%4] |= 0x01 << (rowBitOffset / 4)
			sense[colBitOffset%4] |= 0x01 << (colBitOffset / 4)
			sense[boxBitOffset%4] |= 0x01 << (boxBitOffset / 4)
		}
	}
	return true
}

// And, here is the best one I could find on LeetCode:

/***
Sol Approach v2: Bitmap
Concepts:
    1. Use three array to store the record. (row, col, grid)
    2. For each candidate, we do mask = 1 << int(v), then use row,col,grid to OR that value
    3. Whenever encounter a row,col,grid & mask is true, this mean that we have previously OR this value (duplicate now),
        return false
Time: O(n^2)
Space: O(n)
***/
func isValidSudoku3(board [][]byte) bool {

	var row, col, grid [10]int
	n := len(board)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			v := board[r][c]
			if v != '.' {
				mask := 1 << int(v-'0')
				// if has duplicate
				if (row[r]&mask)|(col[c]&mask)|(grid[r/3*3+c/3]&mask) > 0 {
					return false
				}
				row[r] |= mask
				col[c] |= mask
				grid[r/3*3+c/3] |= mask
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValidSudoku1(example1))
	fmt.Println(isValidSudoku1(example2))

	fmt.Println(isValidSudoku2(example1))
	fmt.Println(isValidSudoku2(example2))

	fmt.Println(isValidSudoku3(example1))
	fmt.Println(isValidSudoku3(example2))

}
