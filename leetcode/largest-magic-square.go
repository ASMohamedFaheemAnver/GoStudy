// https://leetcode.com/problems/largest-magic-square/?envType=daily-question&envId=2026-01-18
package leetcode

func largestMagicSquare(grid [][]int) int {
	columnSize := len(grid)
	rowSize := len(grid[0])
	if columnSize == 1 || rowSize == 1 {
		return 1
	}
	possibleMagicSquare := min(rowSize, columnSize)
	for possibleMagicSquare > 1 {
		for xShift := range rowSize - possibleMagicSquare + 1 {
		yShiftLoop:
			for yShift := range columnSize - possibleMagicSquare + 1 {
				var xMin = xShift
				var xMax = xShift + possibleMagicSquare - 1
				var yMin = yShift
				var yMax = yShift + possibleMagicSquare - 1
				// Row sum check
				var sumRef = 0
				for y := yMin; y <= yMax; y++ {
					tempRowSum := 0
					for x := xMin; x <= xMax; x++ {
						tempRowSum += grid[y][x]
					}
					if sumRef == 0 {
						sumRef = tempRowSum
					} else if sumRef != tempRowSum {
						continue yShiftLoop
					}
				}

				// Column check
				for x := xMin; x <= xMax; x++ {
					tempColSum := 0
					for y := yMin; y <= yMax; y++ {
						tempColSum += grid[y][x]
					}
					if sumRef != tempColSum {
						continue yShiftLoop
					}
				}

				// Top left to bottom right diagonal check
				tempTopLeftToBottomRightSum := 0
				for x, y := xMin, yMin; x <= xMax && y <= yMax; x, y = x+1, y+1 {
					tempTopLeftToBottomRightSum += grid[y][x]
				}
				if sumRef != tempTopLeftToBottomRightSum {
					continue yShiftLoop
				}

				// Top to bottom diagonal check
				tempBottomLeftToTopRightSum := 0
				for x, y := xMin, yMax; x <= xMax && y >= yMin; x, y = x+1, y-1 {
					tempBottomLeftToTopRightSum += grid[y][x]
				}
				if sumRef != tempBottomLeftToTopRightSum {
					continue yShiftLoop
				}

				return possibleMagicSquare
			}
		}
		possibleMagicSquare--
	}
	return 1
}

func LargestMagicSquare(grid [][]int) int {
	return largestMagicSquare(grid)
}
