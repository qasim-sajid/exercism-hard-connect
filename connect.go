package connect

import (
	"fmt"
)

var boardIndexMap map[string]bool

func ResultOf(board []string) (string, error) {
	if Is1X1Board(board) {
		if board[0] == "." {
			return "", nil
		} else {
			return board[0], nil
		}
	}

	winner := ""

	if XWinCheck(board) {
		winner = "X"
	} else if OWinCheck(board) {
		winner = "O"
	}

	return winner, nil
}

func XWinCheck(board []string) bool {
	boardIndexMap = make(map[string]bool)
	for i := 0; i < len(board); i++ {
		if RecursiveWinnerCheck(board, 'X', i, 0) {
			return true
		}
	}
	return false
}

func OWinCheck(board []string) bool {
	boardIndexMap = make(map[string]bool)
	for j := 0; j < len(board[0]); j++ {
		if RecursiveWinnerCheck(board, 'O', 0, j) {
			return true
		}
	}
	return false
}

func RecursiveWinnerCheck(board []string, player rune, currentRow int, currentCol int) bool {
	if []rune(board[currentRow])[currentCol] == player {
		if player == 'X' {
			if currentCol == len(board[currentRow])-1 {
				return true
			}
		} else {
			if currentRow == len(board)-1 {
				return true
			}
		}

		connectedNeighbouringRows, connectedNeighbouringCols := GetConnectedNeighbouringIndexes(board, player, currentRow, currentCol)
		for k := 0; k < len(connectedNeighbouringRows); k++ {
			isIndexAlreadyChecked := false
			mapKey := fmt.Sprintf("%d,%d", connectedNeighbouringRows[k], connectedNeighbouringCols[k])

			if _, ok := boardIndexMap[mapKey]; ok {
				isIndexAlreadyChecked = true
			} else {
				boardIndexMap[mapKey] = true
			}

			if !isIndexAlreadyChecked {
				if RecursiveWinnerCheck(board, player, connectedNeighbouringRows[k], connectedNeighbouringCols[k]) {
					return true
				}
			}
		}
	}

	return false
}

func GetConnectedNeighbouringIndexes(board []string, player rune, currentRow int, currentCol int) ([]int, []int) {
	neighbouringRows, neighbouringCols := GetNeighbouringIndexes(currentRow, currentCol)

	connectedNeighbouringRows := make([]int, 0, 6)
	connectedNeighbouringCols := make([]int, 0, 6)

	for i := 0; i < len(neighbouringRows); i++ {
		if neighbouringRows[i] >= 0 && neighbouringRows[i] < len(board) && neighbouringCols[i] >= 0 && neighbouringCols[i] < len(board[currentRow]) {
			if []rune(board[neighbouringRows[i]])[neighbouringCols[i]] == player {
				connectedNeighbouringRows = append(connectedNeighbouringRows, neighbouringRows[i])
				connectedNeighbouringCols = append(connectedNeighbouringCols, neighbouringCols[i])
			}
		}
	}

	return connectedNeighbouringRows, connectedNeighbouringCols
}

func GetNeighbouringIndexes(row, col int) ([]int, []int) {
	neighbouringRows := make([]int, 6)
	neighbouringCols := make([]int, 6)

	neighbouringRows[0] = row - 1
	neighbouringCols[0] = col

	neighbouringRows[1] = row - 1
	neighbouringCols[1] = col + 1

	neighbouringRows[2] = row
	neighbouringCols[2] = col - 1

	neighbouringRows[3] = row
	neighbouringCols[3] = col + 1

	neighbouringRows[4] = row + 1
	neighbouringCols[4] = col

	neighbouringRows[5] = row + 1
	neighbouringCols[5] = col - 1

	return neighbouringRows, neighbouringCols
}

func Is1X1Board(board []string) bool {
	if len(board) == 1 && len(board[0]) == 1 {
		return true
	} else {
		return false
	}
}
