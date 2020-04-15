package easy

var rock byte = 'R'
var space byte = '.'
var bishop byte = 'B'
var pawn byte = 'p'
var boardLength = 8

func numRookCaptures(board [][]byte) int {
	var rockX, rockY int
	for i := 0; i < boardLength; i++ {
		for j := 0; j < boardLength; j++ {
			x := board[i][j]

			if x == rock {
				rockX = i
				rockY = j
			}
		}
	}

	var total int
	total += catchPawn(board, rockX, rockY, func(x, y int) (int, int) {
		x--
		return x, y
	})
	total += catchPawn(board, rockX, rockY, func(x, y int) (int, int) {
		y--
		return x, y
	})

	total += catchPawn(board, rockX, rockY, func(x, y int) (int, int) {
		x++
		return x, y
	})
	total += catchPawn(board, rockX, rockY, func(x, y int) (int, int) {
		y++
		return x, y
	})
	return total
}

func catchPawn(board [][]byte, x, y int, move func(x, y int) (int, int)) int {
	for {
		x, y = move(x, y)

		if !(x >= 0 && x < boardLength &&
			y >= 0 && y < boardLength) {
			return 0
		}

		switch board[x][y] {
		case bishop:
			return 0
		case space:
			continue
		case pawn:
			return 1
		}
	}

	return 0
}
