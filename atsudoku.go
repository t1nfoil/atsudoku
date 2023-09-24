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

	// save cursor position, clear screen
	fmt.Printf("\0337\033[2J")

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	// clear screen, restore cursor
	fmt.Printf("\033[2J\0338")
}
