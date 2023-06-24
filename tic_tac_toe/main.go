package tic_tac_toe

import "fmt"

type Field map[int]map[int]string

type Game struct {
	field         Field
	currentPlayer string
}

const (
	PlayerCross  = "cross"
	PlayerNaught = "naught"
)

func createEmptyField() Field {
	field := make(Field, 3)

	for i := 1; i <= 3; i++ {
		field[i] = make(map[int]string, 3)
		for j := 1; j <= 3; j++ {
			field[i][j] = ""
		}
	}

	return field
}

func (f Field) Print() {
	fmt.Println("▣ 1|2|3")
	for rowNumber := 1; rowNumber <= 3; rowNumber++ {
		row := f[rowNumber]
		rowString := fmt.Sprintf("%d: ", rowNumber)
		for colNumber := 1; colNumber <= 3; colNumber++ {
			value := row[colNumber]
			rowString += f.printSymbol(value)
		}
		fmt.Println(rowString)
	}
}

func (f Field) printSymbol(value string) string {
	switch value {
	case PlayerCross:
		return "X"
	case PlayerNaught:
		return "0"
	default:
		return "□"
	}
}

func (f Field) setSymbol(rowNumber int, colNumber int, symbol string) bool {
	if colNumber < 1 || colNumber > 3 || rowNumber < 1 || rowNumber > 3 {
		fmt.Println("Wrong column or row number, please set valid values between 1 and 3")
		return false
	}

	if f[rowNumber][colNumber] != "" {
		fmt.Println("This cell is already occupied, please select another one")
		return false
	}

	f[rowNumber][colNumber] = symbol
	return true
}

func StartGame() {
	field := createEmptyField()
	field.Print()
	player := PlayerCross

	game := Game{
		field:         field,
		currentPlayer: player,
	}

	game.run()
}

func (g *Game) run() {
	for {
		fmt.Println("It's", g.currentPlayer, "turn")
		isTurnDone := g.playerTurn()

		if !isTurnDone {
			continue
		}

		isGameFinished, isDraw := g.checkWinner()

		if isGameFinished {
			g.printCongrats(isDraw)
			break
		}

		g.changeCurrentPlayer()
	}
}

func (g *Game) printCongrats(isDraw bool) {
	if isDraw {
		fmt.Println("It's a draw")
		return
	}

	fmt.Printf("Congratulations %s! You are winner!", g.currentPlayer)
}

func (g *Game) changeCurrentPlayer() {
	if g.currentPlayer == PlayerCross {
		g.currentPlayer = PlayerNaught
	} else {
		g.currentPlayer = PlayerCross
	}
}

func (g *Game) playerTurn() (isTurnDone bool) {
	var rowNumber int

	fmt.Println("Please select a row between 1, 2, 3")
	fmt.Scan(&rowNumber)

	var colNumber int
	fmt.Println("Please select a column between 1, 2, 3")
	fmt.Scan(&colNumber)

	if g.field.setSymbol(rowNumber, colNumber, g.currentPlayer) {
		g.field.Print()
		isTurnDone = true
	}

	return isTurnDone
}

func (g *Game) checkWinner() (isGameFinished bool, isDraw bool) {
	isGameFinished = false
	isDraw = false

	for rowNumber := 1; rowNumber <= 3; rowNumber++ {
		row := g.field[rowNumber]
		if row[1] == g.currentPlayer && row[2] == g.currentPlayer && row[3] == g.currentPlayer {
			isGameFinished = true
		}
	}

	for colNumber := 1; colNumber <= 3; colNumber++ {
		if g.field[1][colNumber] == g.currentPlayer && g.field[2][colNumber] == g.currentPlayer && g.field[3][colNumber] == g.currentPlayer {
			isGameFinished = true
		}
	}

	if g.field[1][1] == g.currentPlayer && g.field[2][2] == g.currentPlayer && g.field[3][3] == g.currentPlayer {
		isGameFinished = true
	}

	if g.field[1][3] == g.currentPlayer && g.field[2][2] == g.currentPlayer && g.field[3][1] == g.currentPlayer {
		isGameFinished = true
	}

	count := 0
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if g.field[i][j] != "" {
				count++
			}
		}
	}

	if count >= 9 && !isGameFinished {
		isGameFinished = true
		isDraw = true
	}

	return isGameFinished, isDraw
}
