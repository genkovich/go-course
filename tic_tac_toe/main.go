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

	symbol := "cross"

	for {
		fmt.Println("It's", symbol, "turn")
		win, remakeMove := turn(symbol, field)

		if win {
			break
		}

		if !remakeMove {
			continue
		}

		symbol = changeSymbol(symbol)

	}
}

func changeSymbol(symbol string) string {
	if symbol == "cross" {
		return "naught"
	}

	return "cross"
}

func turn(symbol string, field Field) (bool, bool) {
	var rowNumber int
	fmt.Println("Please select a row between 1, 2, 3")
	fmt.Scan(&rowNumber)

	var colNumber int
	fmt.Println("Please select a column between 1, 2, 3")
	fmt.Scan(&colNumber)

	if !field.setSymbol(rowNumber, colNumber, symbol) {
		return false, false
	}

	field.Print()
	return field.checkWinner(symbol), true
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

func (f Field) checkWinner(symbol string) bool {
	for rowNumber := 1; rowNumber <= 3; rowNumber++ {
		row := f[rowNumber]
		if row[1] == symbol && row[2] == symbol && row[3] == symbol {
			fmt.Println(symbol, "is the winner")
			return true

		}
	}

	for colNumber := 1; colNumber <= 3; colNumber++ {
		if f[1][colNumber] == symbol && f[2][colNumber] == symbol && f[3][colNumber] == symbol {
			fmt.Println(symbol, "is the winner")
			return true
		}
	}

	if f[1][1] == symbol && f[2][2] == symbol && f[3][3] == symbol {
		fmt.Println(symbol, "is the winner")
		return true
	}

	if f[1][3] == symbol && f[2][2] == symbol && f[3][1] == symbol {
		fmt.Println(symbol, "is the winner")
		return true
	}

	if f[1][1] != "" && f[1][2] != "" && f[1][3] != "" && f[2][1] != "" && f[2][2] != "" && f[2][3] != "" && f[3][1] != "" && f[3][2] != "" && f[3][3] != "" {
		fmt.Println("It's a draw")
		return true
	}

	return false
}
