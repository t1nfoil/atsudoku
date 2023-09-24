package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type fontCharacter struct {
	rows [5]string
}

var fcZero = fontCharacter{
	rows: [5]string{
		"         ",
		"         ",
		"    ▁    ",
		"         ",
		"         ",
	},
}

var fcOne = fontCharacter{
	rows: [5]string{
		"    ▃    ",
		"   ▀█    ",
		"    █    ",
		"    █    ",
		"  ▀▀▀▀▀  ",
	},
}

var fcTwo = fontCharacter{
	rows: [5]string{
		"   ▃▃▃   ",
		"  ▀   █  ",
		"    ▃▀   ",
		"  ▃▀     ",
		"  ▀▀▀▀▀  ",
	},
}

var fcThree = fontCharacter{
	rows: [5]string{
		"   ▃▃▃   ",
		"  ▀   █  ",
		"    ▀▀▃  ",
		"  ▃   █  ",
		"   ▀▀▀   ",
	},
}

var fcFour = fontCharacter{
	rows: [5]string{
		"     ▃   ",
		"  █  █   ",
		"  █▃▃█▃  ",
		"     █   ",
		"     ▀   ",
	},
}

var fcFive = fontCharacter{
	rows: [5]string{
		"  ▃▃▃▃▃  ",
		"  █      ",
		"  ▀▀▀▀▃  ",
		"      █  ",
		"  ▀▀▀▀   ",
	},
}

var fcSix = fontCharacter{
	rows: [5]string{
		"   ▃▃▃   ",
		"  █      ",
		"  █▀▀▀▃  ",
		"  █   █  ",
		"   ▀▀▀   ",
	},
}

var fcSeven = fontCharacter{
	rows: [5]string{
		"  ▃▃▃▃▃  ",
		"     ▃▀  ",
		"    ▃▀   ",
		"    █    ",
		"    ▀    ",
	},
}

var fcEight = fontCharacter{
	rows: [5]string{
		"   ▃▃▃   ",
		"  █   █  ",
		"  ▃▀▀▀▃  ",
		"  █   █  ",
		"   ▀▀▀   ",
	},
}

var fcNine = fontCharacter{
	rows: [5]string{
		"   ▃▃▃   ",
		"  █   █  ",
		"   ▀▀▀█  ",
		"      █  ",
		"   ▀▀▀   ",
	},
}

var font = map[int]fontCharacter{
	0: fcZero,
	1: fcOne,
	2: fcTwo,
	3: fcThree,
	4: fcFour,
	5: fcFive,
	6: fcSix,
	7: fcSeven,
	8: fcEight,
	9: fcNine,
}

func (f *fontCharacter) printFontCharacter(row, column, fgColor, bgColor int) {
	fmt.Print("\033[s")
	for i := 0; i < 5; i++ {
		fmt.Printf("\033[%dm\033[%dm\033[%d;%dH", fgColor, bgColor, row+i, column)
		fmt.Print(f.rows[i])
	}
	fmt.Print("\033[u")
}

type model struct {
	stopwatch    stopwatch.Model
	gridX, gridY int
}

func initialModel() model {
	return model{gridX: 4,
		gridY:     4,
		stopwatch: stopwatch.NewWithInterval(time.Second),
	}
}

func (m model) Init() tea.Cmd {
	s.setBoardInitValues()
	for {
		s.generateBoard()

		if s.validateSolution() {
			break
		}
		s.setBoardInitValues()
	}
	s.generatePuzzle()
	return m.stopwatch.Init()

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newSudoku := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.gridY-1 < 0 {
				m.gridY = 8
			} else {
				m.gridY--
			}
		case "down":
			if m.gridY+1 > 8 {
				m.gridY = 0
			} else {
				m.gridY++
			}
		case "left":
			if m.gridX-1 < 0 {
				m.gridX = 8
			} else {
				m.gridX--
			}
		case "right":
			if m.gridX+1 > 8 {
				m.gridX = 0
			} else {
				m.gridX++
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if s.userFillable[m.gridY][m.gridX] {
				s.sudoku[m.gridY][m.gridX] = int(msg.String()[0]) - 48
			}
		case "delete", "-", "backspace":
			if s.userFillable[m.gridY][m.gridX] {
				s.sudoku[m.gridY][m.gridX] = 0
			}
		case "h", "H":
			if s.hintsEnabled {
				s.hintsEnabled = false
			} else {
				s.hintsEnabled = true
			}
		case "n":

			if s.difficulty == "easy" {
				s.difficulty = "medium"
				goto generateNewSudoku
			}
			if s.difficulty == "medium" {
				s.difficulty = "hard"
				goto generateNewSudoku
			}
			if s.difficulty == "hard" {
				s.difficulty = "easy"
				goto generateNewSudoku
			}
		generateNewSudoku:
			s.setBoardInitValues()
			for {
				s.generateBoard()
				if s.validateSolution() {
					break
				}
				s.setBoardInitValues()
			}
			s.generatePuzzle()
			newSudoku = true
		}
	}
	var cmd tea.Cmd
	if newSudoku {
		cmd = m.stopwatch.Reset()
		m.stopwatch, _ = m.stopwatch.Update(msg)
		return m, cmd
	}
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s.displayPuzzle(m.gridX, m.gridY)
	s.displayInfo(m.stopwatch.Elapsed())
	return ""
}
