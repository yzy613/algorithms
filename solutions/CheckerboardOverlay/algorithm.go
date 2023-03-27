package CheckerboardOverlay

var (
	layer = 2
)

func overlay(checkerboard [][]int, size, rowStart, colStart, row, col int) {
	currLayer := layer
	switch size {
	case 0:
		return
	case 1:
		if checkerboard[rowStart][colStart] == 0 {
			checkerboard[rowStart][colStart] = currLayer
		}
		return
	}
	nextSize := size / 2
	position := 0
	switch {
	case row < rowStart+nextSize && col < colStart+nextSize:
		// top left
		position = 1
	case row < rowStart+nextSize && col >= colStart+nextSize:
		// top right
		position = 2
	case row >= rowStart+nextSize && col < colStart+nextSize:
		// bottom left
		position = 3
	case row >= rowStart+nextSize && col >= colStart+nextSize:
		// bottom right
		position = 4
	}
	// 填充
	var (
		blockRow1, blockCol1 int
		blockRow2, blockCol2 int
		blockRow3, blockCol3 int
		blockRow4, blockCol4 int
	)
	if position != 1 {
		blockRow1 = rowStart + nextSize - 1
		blockCol1 = colStart + nextSize - 1
		if checkerboard[blockRow1][blockCol1] == 0 {
			checkerboard[blockRow1][blockCol1] = currLayer
		}
	}
	if position != 2 {
		blockRow2 = rowStart + nextSize - 1
		blockCol2 = colStart + nextSize
		if checkerboard[blockRow2][blockCol2] == 0 {
			checkerboard[blockRow2][blockCol2] = currLayer
		}
	}
	if position != 3 {
		blockRow3 = rowStart + nextSize
		blockCol3 = colStart + nextSize - 1
		if checkerboard[blockRow3][blockCol3] == 0 {
			checkerboard[blockRow3][blockCol3] = currLayer
		}
	}
	if position != 4 {
		blockRow4 = rowStart + nextSize
		blockCol4 = colStart + nextSize
		if checkerboard[blockRow4][blockCol4] == 0 {
			checkerboard[blockRow4][blockCol4] = currLayer
		}
	}
	// nextLayer
	layer++
	if position != 1 {
		overlay(checkerboard, nextSize, rowStart, colStart, blockRow1, blockCol1)
	} else {
		overlay(checkerboard, nextSize, rowStart, colStart, row, col)
	}
	if position != 2 {
		overlay(checkerboard, nextSize, rowStart, colStart+nextSize, blockRow2, blockCol2)
	} else {
		overlay(checkerboard, nextSize, rowStart, colStart+nextSize, row, col)
	}
	if position != 3 {
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart, blockRow3, blockCol3)
	} else {
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart, row, col)
	}
	if position != 4 {
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart+nextSize, blockRow4, blockCol4)
	} else {
		overlay(checkerboard, nextSize, rowStart+nextSize, colStart+nextSize, row, col)
	}
	return
}
