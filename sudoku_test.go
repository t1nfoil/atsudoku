package main

import (
	"testing"
)

func TestDoesNumExistInRowSol(t *testing.T) {
	tests := []struct {
		board  SudokuBoard
		row    int
		value  int
		expect bool
	}{
		{SudokuBoard{solution: [9][9]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, 1, true},
		{SudokuBoard{solution: [9][9]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, 10, false},
		{SudokuBoard{solution: [9][9]int{{0, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, 0, true},
		{SudokuBoard{solution: [9][9]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, -1, false},
	}

	for _, test := range tests {
		got := test.board.doesNumExistInRowSol(test.row, test.value)
		if got != test.expect {
			t.Errorf("Expected %v, got %v", test.expect, got)
		}
	}
}

func TestDoesNumExistInColumnSud(t *testing.T) {
	tests := []struct {
		board  SudokuBoard
		column int
		value  int
		expect bool
	}{
		{SudokuBoard{sudoku: [9][9]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, 1, true},
		{SudokuBoard{sudoku: [9][9]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, 10, false},
		{SudokuBoard{sudoku: [9][9]int{{0}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, 0, true},
		{SudokuBoard{sudoku: [9][9]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, -1, false},
	}

	for _, test := range tests {
		got := test.board.doesNumExistInColumnSud(test.column, test.value)
		if got != test.expect {
			t.Errorf("Expected %v, got %v", test.expect, got)
		}
	}
}

func TestDoesNumExistInRowSud(t *testing.T) {
	tests := []struct {
		board  SudokuBoard
		row    int
		value  int
		expect bool
	}{
		{SudokuBoard{sudoku: [9][9]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, 1, true},
		{SudokuBoard{sudoku: [9][9]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, 10, false},
		{SudokuBoard{sudoku: [9][9]int{{0, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, 0, true},
		{SudokuBoard{sudoku: [9][9]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}}, 0, -1, false},
	}

	for _, test := range tests {
		got := test.board.doesNumExistInRowSud(test.row, test.value)
		if got != test.expect {
			t.Errorf("Expected %v, got %v", test.expect, got)
		}
	}
}

func TestDoesNumExistInColumnSol(t *testing.T) {
	tests := []struct {
		board  SudokuBoard
		column int
		value  int
		expect bool
	}{
		{SudokuBoard{solution: [9][9]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, 1, true},
		{SudokuBoard{solution: [9][9]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, 10, false},
		{SudokuBoard{solution: [9][9]int{{0}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, 0, true},
		{SudokuBoard{solution: [9][9]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, 0, -1, false},
	}

	for _, test := range tests {
		got := test.board.doesNumExistInColumnSol(test.column, test.value)
		if got != test.expect {
			t.Errorf("Expected %v, got %v", test.expect, got)
		}
	}
}

func TestValidateSolution(t *testing.T) {
	tests := []struct {
		name   string
		board  SudokuBoard
		expect bool
	}{
		{name: "valid solution",
			board: SudokuBoard{solution: [9][9]int{
				{8, 3, 7, 5, 4, 9, 2, 1, 6},
				{9, 6, 4, 2, 1, 7, 3, 5, 8},
				{5, 1, 2, 8, 6, 3, 9, 7, 4},
				{7, 9, 6, 3, 5, 2, 8, 4, 1},
				{3, 8, 1, 6, 7, 4, 5, 9, 2},
				{4, 2, 5, 9, 8, 1, 7, 6, 3},
				{6, 4, 8, 7, 2, 5, 1, 3, 9},
				{1, 5, 9, 4, 3, 8, 6, 2, 7},
				{2, 7, 3, 1, 9, 6, 4, 8, 5},
			}},
			expect: true,
		},
		{name: "valid solution",
			board: SudokuBoard{solution: [9][9]int{
				{2, 7, 3, 8, 9, 6, 5, 4, 1},
				{1, 6, 8, 5, 7, 4, 2, 3, 9},
				{5, 4, 9, 2, 1, 3, 8, 7, 6},
				{7, 1, 5, 6, 4, 2, 9, 8, 3},
				{9, 8, 4, 3, 5, 1, 6, 2, 7},
				{6, 3, 2, 7, 8, 9, 1, 5, 4},
				{3, 9, 1, 4, 2, 5, 7, 6, 8},
				{8, 5, 6, 1, 3, 7, 4, 9, 2},
				{4, 2, 7, 9, 6, 8, 3, 1, 5},
			}},
			expect: true,
		},
		{name: "invalid solution (duplicate in row)",
			board: SudokuBoard{solution: [9][9]int{
				{2, 7, 3, 8, 9, 6, 5, 4, 1},
				{1, 6, 8, 5, 7, 4, 2, 3, 9},
				{5, 4, 9, 2, 1, 3, 8, 7, 6},
				{7, 1, 5, 6, 4, 2, 9, 8, 3},
				{9, 8, 4, 3, 5, 1, 6, 2, 7},
				{6, 3, 2, 7, 5, 9, 1, 5, 4}, // <-- Duplicate 5 in column/row
				{3, 9, 1, 4, 2, 5, 7, 6, 8},
				{8, 5, 6, 1, 3, 7, 4, 9, 2},
				{4, 2, 7, 9, 6, 8, 3, 1, 5},
			}},
			expect: false,
		},
		{
			name: "invalid solution (duplicate in column)",
			board: SudokuBoard{solution: [9][9]int{
				{8, 3, 7, 5, 4, 9, 2, 1, 6},
				{9, 6, 4, 2, 1, 7, 3, 5, 8},
				{5, 1, 2, 8, 6, 3, 9, 7, 4},
				{7, 9, 6, 3, 5, 2, 8, 4, 1},
				{3, 8, 1, 6, 7, 4, 5, 9, 2},
				{4, 2, 5, 9, 8, 1, 7, 6, 3},
				{6, 4, 8, 7, 2, 7, 1, 3, 9}, // <-- Duplicate 7 in colunm/row
				{1, 5, 9, 4, 3, 8, 6, 2, 7},
				{2, 7, 3, 1, 9, 6, 4, 8, 5},
			}},
			expect: false,
		},
	}

	for _, test := range tests {
		got := test.board.validateSolution()
		if got != test.expect {
			t.Errorf("Test: %v, expected %v, got %v", test.name, test.expect, got)
		}
	}
}

func TestValidateSudoku(t *testing.T) {
	tests := []struct {
		name   string
		board  SudokuBoard
		expect bool
	}{
		{
			name: "valid alternate solution",
			board: SudokuBoard{solution: [9][9]int{
				{4, 6, 2, 9, 8, 7, 1, 5, 3},
				{5, 3, 7, 4, 6, 1, 8, 9, 2},
				{1, 8, 9, 3, 5, 2, 4, 6, 7},
				{2, 1, 4, 8, 9, 6, 3, 7, 5},
				{3, 5, 6, 7, 1, 4, 9, 2, 8},
				{7, 9, 8, 2, 3, 5, 6, 1, 4},
				{8, 2, 1, 5, 4, 9, 7, 3, 6},
				{9, 4, 5, 6, 7, 3, 2, 8, 1},
				{6, 7, 3, 1, 2, 8, 5, 4, 9},
			},
				sudoku: [9][9]int{
					{4, 6, 2, 9, 8, 7, 1, 5, 3},
					{5, 3, 7, 4, 6, 1, 8, 9, 2},
					{1, 8, 9, 3, 5, 2, 4, 6, 7},
					{2, 1, 6, 8, 4, 9, 3, 7, 5},
					{3, 5, 4, 7, 1, 6, 9, 2, 8},
					{7, 9, 8, 2, 3, 5, 6, 1, 4},
					{8, 2, 1, 5, 9, 4, 7, 3, 6},
					{9, 4, 5, 6, 7, 3, 2, 8, 1},
					{6, 7, 3, 1, 2, 8, 5, 4, 9},
				}},
			expect: true,
		},
		{
			name: "invalid alternate solution",
			board: SudokuBoard{solution: [9][9]int{
				{4, 6, 2, 9, 8, 7, 1, 5, 3},
				{5, 3, 7, 4, 6, 1, 8, 9, 2},
				{1, 8, 9, 3, 5, 2, 4, 6, 7},
				{2, 1, 4, 8, 9, 6, 3, 7, 5},
				{3, 5, 6, 7, 1, 4, 9, 2, 8},
				{7, 9, 8, 2, 3, 5, 6, 1, 4},
				{8, 2, 1, 5, 4, 9, 7, 3, 6},
				{9, 4, 5, 6, 7, 3, 2, 8, 1},
				{6, 7, 3, 1, 2, 8, 5, 4, 9},
			},
				sudoku: [9][9]int{
					{4, 6, 2, 9, 8, 7, 1, 5, 3},
					{5, 3, 7, 4, 6, 1, 8, 9, 2},
					{1, 8, 9, 3, 5, 2, 4, 6, 7},
					{2, 1, 6, 8, 4, 9, 3, 7, 5},
					{3, 5, 4, 7, 1, 6, 9, 2, 8},
					{7, 9, 8, 2, 3, 5, 6, 1, 4},
					{8, 2, 1, 5, 4, 9, 7, 3, 6},
					{9, 4, 5, 6, 7, 3, 2, 8, 1},
					{6, 7, 3, 1, 2, 8, 5, 4, 9},
				}},
			expect: false,
		},
	}

	for _, test := range tests {
		got := test.board.validateSudoku()
		if got != test.expect {
			t.Errorf("Test: %v, expected %v, got %v", test.name, test.expect, got)
		}
	}
}
