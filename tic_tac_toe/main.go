package tic_tac_toe

import "fmt"

type Field map[int]map[int]string

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
			rowString += getSymbol(value)
		}
		fmt.Println(rowString)
	}
}

func getSymbol(value string) string {
	switch value {
	case "cross":
		return "X"
	case "naught":
		return "0"
	default:
		return "□"
	}
}

func StartGame() {
	field := createEmptyField()
	field.Print()

	player := "cross"
	gameLoop(player, field)
}

func gameLoop(player string, field Field) {
	for {
		fmt.Println("It's", player, "turn")
		isTurnDone := playerTurn(player, field)

		if !isTurnDone {
			continue
		}

		isGameFinished, isDraw := field.checkWinner(player)

		if isGameFinished {
			printCongrats(player, isDraw)
			break
		}

		player = changePlayer(player)
	}
}

func printCongrats(player string, isDraw bool) {
	if isDraw {
		fmt.Println("It's a draw")
		return
	}

	fmt.Printf("Congratulations %s! You are winner!", player)
}

func changePlayer(player string) string {
	if player == "cross" {
		return "naught"
	}

	return "cross"
}

func playerTurn(player string, field Field) (isTurnDone bool) {
	var rowNumber int

	fmt.Println("Please select a row between 1, 2, 3")
	fmt.Scan(&rowNumber)

	var colNumber int
	fmt.Println("Please select a column between 1, 2, 3")
	fmt.Scan(&colNumber)

	if field.setSymbol(rowNumber, colNumber, player) {
		field.Print()
		isTurnDone = true
	}

	return isTurnDone
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

func (f Field) checkWinner(player string) (isGameFinished bool, isDraw bool) {
	isGameFinished = false
	isDraw = false

	for rowNumber := 1; rowNumber <= 3; rowNumber++ {
		row := f[rowNumber]
		if row[1] == player && row[2] == player && row[3] == player {
			isGameFinished = true
		}
	}

	for colNumber := 1; colNumber <= 3; colNumber++ {
		if f[1][colNumber] == player && f[2][colNumber] == player && f[3][colNumber] == player {
			isGameFinished = true
		}
	}

	if f[1][1] == player && f[2][2] == player && f[3][3] == player {
		isGameFinished = true
	}

	if f[1][3] == player && f[2][2] == player && f[3][1] == player {
		isGameFinished = true
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if f[i][j] == "" {
				isGameFinished = true
				isDraw = true
			}
		}
	}

	return isGameFinished, isDraw
}
