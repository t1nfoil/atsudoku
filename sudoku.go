package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SudokuBoard struct {
	difficulty   string
	hintsEnabled bool
	sudokuSolved bool
	solution     [9][9]int
	sudoku       [9][9]int
	userFillable [9][9]bool
}

func (s *SudokuBoard) doesNumExistInRow(row, value int) bool {
	if row >= 0 && row <= 8 {
		for column := 0; column <= 8; column++ {
			if s.getValue(row, column) == value {
				return true
			}
		}
	}
	return false
}

func (s *SudokuBoard) doesNumExistInColumn(column, value int) bool {
	if column >= 0 && column <= 8 {
		for row := 0; row <= 8; row++ {
			if s.getValue(row, column) == value {
				return true
			}
		}
	}
	return false
}

func (s *SudokuBoard) getValue(row, column int) int {
	return s.solution[row][column]
}

func (s *SudokuBoard) setValue(row, column, value int) {
	s.solution[row][column] = value
}

func (s *SudokuBoard) backTrack(value int) {
	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			if s.solution[row][column] == value {
				s.solution[row][column] = 0
			}
		}
	}
}

func (s *SudokuBoard) setBoardInitValues() {
	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			s.solution[row][column] = 0
		}
	}
}

func (s *SudokuBoard) validateBoard() bool {
	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			if s.solution[row][column] == 0 {
				return false
			}
		}
	}

	for row := 0; row <= 8; row++ {
		for column := 1; column <= 9; column++ {
			if !s.doesNumExistInRow(row, column) {
				return false
			}
		}
	}

	for row := 0; row <= 8; row++ {
		for column := 1; column <= 9; column++ {
			if !s.doesNumExistInColumn(row, column) {
				return false
			}
		}
	}

	var gc gridCell
	gridCells := gc.getGrid()
	for row := 0; row <= 8; row++ {
		for number := 1; number <= 9; number++ {
			for column := 0; column <= 8; column++ {
				if s.getValue(gridCells[row].row+column/3, gridCells[row].column+column%3) == number {
					break
				}
				if column == 8 {
					return false
				}
			}
		}
	}

	return true
}

type gridCell struct {
	row, column int
}

func (g *gridCell) getGrid() [9]gridCell {
	return ([9]gridCell{
		{row: 0, column: 0},
		{row: 0, column: 3},
		{row: 0, column: 6},

		{row: 3, column: 0},
		{row: 3, column: 3},
		{row: 3, column: 6},

		{row: 6, column: 0},
		{row: 6, column: 3},
		{row: 6, column: 6},
	})
}

func (s *SudokuBoard) generateBoard() {
	var gc gridCell
	grid := gc.getGrid()
	numberOfIterations := 0
	numberOfAttempts := 0
	sudokuNumber := 1
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		gridNumber := -1
	next:

		if gridNumber+1 > 8 {
			sudokuNumber++
			gridNumber = 0
			if sudokuNumber == 10 {
				break
			}
		} else {
			gridNumber++
		}

		for {
			numberOfIterations++
			randomRowOffset := rnd.Intn(3)
			randomColumnOffset := rnd.Intn(3)

			if s.getValue(grid[gridNumber].row+randomRowOffset, grid[gridNumber].column+randomColumnOffset) == 0 {
				if !s.doesNumExistInColumn(grid[gridNumber].column+randomColumnOffset, sudokuNumber) && !s.doesNumExistInRow(grid[gridNumber].row+randomRowOffset, sudokuNumber) {
					s.setValue(grid[gridNumber].row+randomRowOffset, grid[gridNumber].column+randomColumnOffset, sudokuNumber)
					goto next
				}
			}

			if numberOfAttempts > 200 {
				numberOfAttempts = 0
				gridNumber = 0
				numberOfIterations = 0
				s.setBoardInitValues()
			}

			if numberOfIterations > 1000 {
				numberOfAttempts++
				gridNumber = 0
				numberOfIterations = 0
				s.backTrack(sudokuNumber)
			}

		}
	}
}

func (s *SudokuBoard) generatePuzzle() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var removeElements int
	switch s.difficulty {
	case "hard":
		removeElements = 51
	case "medium":
		removeElements = 46
	case "easy":
		removeElements = 40
	default:
		removeElements = 46
	}

	var indicesToRemove []int
	for i := 0; i < removeElements; i++ {
		for {
			index := rnd.Intn(81)
			isUnique := true
			for _, v := range indicesToRemove {
				if v == index {
					isUnique = false
				}
			}
			if isUnique {
				indicesToRemove = append(indicesToRemove, index)
				break
			}
		}
	}

	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			s.userFillable[row][column] = false
		}
	}

	s.sudoku = s.solution
	for _, v := range indicesToRemove {
		s.sudoku[v/9][v%9] = 0
		s.userFillable[v/9][v%9] = true
	}
}

func (s *SudokuBoard) getBgColorForRowCol(row, column int) int {
	if row >= 0 && row <= 2 && column >= 0 && column <= 2 {
		return 45
	}
	if row >= 0 && row <= 2 && column >= 3 && column <= 5 {
		return 40
	}
	if row >= 0 && row <= 2 && column >= 6 && column <= 8 {
		return 45
	}
	if row >= 3 && row <= 5 && column >= 0 && column <= 2 {
		return 40
	}
	if row >= 3 && row <= 5 && column >= 3 && column <= 5 {
		return 45
	}
	if row >= 3 && row <= 5 && column >= 6 && column <= 8 {
		return 40
	}
	if row >= 6 && row <= 8 && column >= 0 && column <= 2 {
		return 45
	}
	if row >= 6 && row <= 8 && column >= 3 && column <= 5 {
		return 40
	}
	if row >= 6 && row <= 8 && column >= 6 && column <= 8 {
		return 45
	}
	return 45
}

func (s *SudokuBoard) displayPuzzle(gridX, gridY int) {
	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			bgColor := s.getBgColorForRowCol(row, column)
			number := font[s.sudoku[row][column]]
			if row == gridY && column == gridX && s.userFillable[row][column] {
				number := font[s.sudoku[row][column]]
				number.printFontCharacter((row*5)+2, column*9+55, 93, 44)
				continue
			}
			if row == gridY && column == gridX && !s.userFillable[row][column] {
				number := font[s.sudoku[row][column]]
				number.printFontCharacter((row*5)+2, column*9+55, 97, 44)
				continue
			}
			if s.userFillable[row][column] {
				number.printFontCharacter((row*5)+2, column*9+55, 93, bgColor)
				continue
			} else {
				number.printFontCharacter((row*5)+2, column*9+55, 97, bgColor)
				continue
			}
		}
	}
}

func (s *SudokuBoard) displayInfo(elapsedTime time.Duration) {

	fmt.Printf("\033[3;8HSudoku Difficulty: %s   ", s.difficulty)
	s.sudokuSolved = true
	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			if s.sudoku[row][column] != s.solution[row][column] {
				s.sudokuSolved = false
			}
		}
	}
	sudokuStatus := "Solved"
	if !s.sudokuSolved {
		sudokuStatus = "Unsolved"
	}
	fmt.Printf("\033[4;8HSudoku Status: %s   ", sudokuStatus)

	fmt.Printf("\033[5;8HElapsed Time: %s   ", elapsedTime)

	if s.hintsEnabled {
		fmt.Printf("\033[7;8HHints: Enabled ")
		fmt.Print("\033[s")
		for row := 0; row <= 8; row++ {
			for column := 0; column <= 8; column++ {
				bgColor := s.getBgColorForRowCol(row, column)
				fmt.Printf("\033[%dm\033[%dm\033[%d;%dH%d", 97, bgColor, 7+row, 25+column, s.solution[row][column])
				//fmt.Printf("\033[%d;%dH%d ", 10+row, 10+column, s.solution[row][column])
			}
		}
		fmt.Print("\033[u")
	} else {
		fmt.Printf("\033[7;8HHints: Disabled")
		for row := 0; row <= 8; row++ {
			for column := 0; column <= 8; column++ {
				fmt.Printf("\033[%d;%dH ", 7+row, 25+column)
			}
		}
	}

	fmt.Printf("\033[17;8HNavigate with the arrow keys:")
	fmt.Printf("\033[18;16H       ┏━━━━━┓")
	fmt.Printf("\033[19;16H       ┃  ↑  ┃")
	fmt.Printf("\033[20;16H       ┗━━━━━┛")
	fmt.Printf("\033[21;16H┏━━━━━┓┏━━━━━┓┏━━━━━┓")
	fmt.Printf("\033[22;16H┃  ←  ┃┃  ↓  ┃┃  →  ┃")
	fmt.Printf("\033[23;16H┗━━━━━┛┗━━━━━┛┗━━━━━┛")
	fmt.Printf("\033[25;8HInsert numbers by pressing the number keys")
	fmt.Printf("\033[26;16H┏━━━━━┓       ┏━━━━━┓")
	fmt.Printf("\033[27;16H┃  1  ┃  ...  ┃  9  ┃")
	fmt.Printf("\033[28;16H┗━━━━━┛       ┗━━━━━┛")
	fmt.Printf("\033[30;8HDelete numbers by pressing")
	fmt.Printf("\033[31;16H┏━━━━━┓       ┏━━━━━┓")
	fmt.Printf("\033[32;16H┃ del ┃   or  ┃  -  ┃")
	fmt.Printf("\033[33;16H┗━━━━━┛       ┗━━━━━┛")
	fmt.Printf("\033[35;8HPress 'q' to quit, press 'h' for sudoku hints")
	fmt.Printf("\033[36;16H┏━━━━━┓       ┏━━━━━┓")
	fmt.Printf("\033[37;16H┃  q  ┃  ,    ┃  h  ┃")
	fmt.Printf("\033[38;16H┗━━━━━┛       ┗━━━━━┛")
	fmt.Printf("\033[40;8HPress 'n' to cycle new board / difficulty")
	fmt.Printf("\033[41;16H       ┏━━━━━┓")
	fmt.Printf("\033[42;16H       ┃  n  ┃")
	fmt.Printf("\033[43;16H       ┗━━━━━┛")

}
