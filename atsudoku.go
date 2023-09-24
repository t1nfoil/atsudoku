package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var s SudokuBoard

func main() {

	flag.StringVar(&s.difficulty, "level", "medium", "the difficulty level, one of 'easy', 'medium' or 'hard' (default 'medium')")
	flag.Parse()

	if s.difficulty != "easy" && s.difficulty != "medium" && s.difficulty != "hard" {
		fmt.Println("invalid difficulty level, (--level) must be one of 'easy', 'medium' or 'hard'")
		return
	}

	fmt.Printf("\0337\033[2J")

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("\033[2J\0338")

	// s.generatePuzzle("easy")

	// s.printPuzzle()

	// fmt.Print("\033[H\033[2J")

	// // print the puzzle board using font characters, increment columns by 11 and rows by 7

	// for row := 0; row <= 8; row++ {
	// 	for column := 0; column <= 8; column++ {
	// 		number := font[s.puzzle[row][column]]
	// 		number.printFontCharacter((row*5)+2, column*9+55, 36, 40)
	// 	}
	// }

	//s.printPuzzle()
}
