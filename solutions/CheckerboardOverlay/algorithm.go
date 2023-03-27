package CheckerboardOverlay

var (
	layer = 2
)

func overlay(checkerboard [][]int, size, rowStart, colStart, row, col int) {
	if size <= 1 {
		return
	}
	currLayer := layer
	layer++
	nextSize := size / 2
	// top left
	if row < rowStart+nextSize && col < colStart+nextSize {
		overlay(checkerboard, nextSize, rowStart, colStart, row, col)
	} else {
		blockRow1 := rowStart + nextSize - 1
		blockCol1 := colStart + nextSize - 1
		checkerboard[blockRow1][blockCol1] = currLayer
		overlay(checkerboard, nextSize, rowStart, colStart, blockRow1, blockCol1)
	}
	// top right
	if row < rowStart+nextSize && col >= colStart+nextSize {
		overlay(checkerboard, nextSize, rowStart, colStart+nextSize, row, col)
	} else {
		blockRow2 := rowStart + nextSize - 1
		blockCol2 := colStart + nextSize
		checkerboard[blockRow2][blockCol2] = currLayer
		overlay(checkerboard, nextSize, rowStart, colStart+nextSize, blockRow2, blockCol2)
	}
	// bottom left
	if row >= rowStart+nextSize && col < colStart+nextSize {
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart, row, col)
	} else {
		blockRow3 := rowStart + nextSize
		blockCol3 := colStart + nextSize - 1
		checkerboard[blockRow3][blockCol3] = currLayer
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart, blockRow3, blockCol3)
	}
	// bottom right
	if row >= rowStart+nextSize && col >= colStart+nextSize {
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart+nextSize, row, col)
	} else {
		blockRow4 := rowStart + nextSize
		blockCol4 := colStart + nextSize
		checkerboard[blockRow4][blockCol4] = currLayer
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart+nextSize, blockRow4, blockCol4)
	}
	return
}
