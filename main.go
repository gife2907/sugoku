package main

/*
	=== infos ===

	Gilles Ferrero
	https://github.com/gife2907/sugoku

	=== to do ===

	v	solves a real grid
	v	github
	v	byte to int8 conv
	v	some more test samples
	v	time measurement

	x	gopath issues
	x	launch.js and popup
	x	dble click on message opens editor
	x	package : go help gopath, https://www.golang-book.com/books/intro/11

	ut
	resolution speed and memory optimization
	more go like code w code review
	more samples with failing cases
	log appropriately
	api
	service

	vr read
	ar vid overlay

	as an online Pi service

*/

import (
	"fmt"
	"sugoutils"
	"time"
)

func prn(s [9][9]int8) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] != 0 {
				fmt.Print(" ", s[i][j])
			} else {
				fmt.Print(" .")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func isValidSolution(s [9][9]int8) bool {

	// checks if s is a valid, fully filled sudoku grid

	var lineSum, colSum int8
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == 0 {
				return false
			} else {
				lineSum += s[i][j]
				colSum += s[j][i]
			}
		}
		if lineSum != 45 || colSum != 45 {
			return false
		} else {
			lineSum = 0
			colSum = 0
		}
	}
	return true
}

func listAcceptableValues(s [9][9]int8, i int, j int) []int8 {

	// An acceptable value in i,j is a value absent in column i and line j

	res := make([]int8, 0)
	if s[i][j] != 0 {
		fmt.Println("There should not be any value here")
	} else {
		var nextVal int8 = 1
		for nextVal < 10 {
			for v := 0; v < 9; v++ {
				// inspect presence in line
				if s[i][v] == nextVal {
					nextVal++
					// -1 below instead of 0, so that 'continue' then v++ will restart on 0, quite unelegant ...
					v = -1
					continue
				}
				// inspect presence in line
				if s[v][j] == nextVal {
					nextVal++
					// -1 below instead of 0, so that 'continue' then v++ will restart on 0, quite unelegant ...
					v = -1
					continue
				}
			}
			if nextVal != 10 {
				res = append(res, nextVal)
				nextVal++
			} else {
				break
			}
		}
	}
	return res
}

func nextEmptyLocation(s [9][9]int8, i int, j int) (int, int, bool) {

	// returns the next i,j location if any, where there is no value yet (0 is no value yet)
	// the meaning of next is 'next by reading order' and is deterministic as needed

	for u := i; u < 9; u++ {
		for v := j; v < 9; v++ {
			if s[u][v] == 0 {
				return u, v, true
			}
		}
		j = 0
	}
	return -1, -1, false
}

func prepDLocList(s [9][9]int8) {

	// Prepare the Deterministic Location List (where there are zeros)
	// Going through this list in deterministic order is a mandatory aspect of backtracking
	// This is because once a branch has failed, you need to tray with the next value and
	// and the same location order to try with new values
	// In the standard version of backtracking the deterministic order is the reading order

}

func solve(s [9][9]int8, i int, j int) ([9][9]int8, bool) {

	// This is a classic backtracking algorithm (algoritme des flots en Francais)

	// This function is to be called on an available 0 location
	if s[i][j] == 0 {

		// If on a 0 let's list acceptable values
		options := listAcceptableValues(s, i, j)

		for _, option := range options {

			// for each acceptable value we create a new grid with the value ...
			s2 := s
			s2[i][j] = option

			// ... and for this grid we find the next 0 ...
			u, v, ok := nextEmptyLocation(s2, i, j)
			if ok {
				// ... and we calculate the solution with the new grid and the next 0 location ...
				attempt, status := solve(s2, u, v)
				if status {
					// ... and if it is a solution then we return this solution ...
					return attempt, status
				} else {
					// ... otherwise we continue to the next acceptable value in the list
					continue
				}
			} else {
				// There is no more zeros so we need to test the situation
				if isValidSolution(s2) {
					return s2, true
				} else {
					return s2, false
				}
			}
		}
	}
	return s, false

}

func main() {

	start := sugoutils.Grid01()

	prn(start)
	fmt.Println()
	startTime := time.Now()

	startx, starty, hasEmptyLocation := nextEmptyLocation(start, 0, 0)
	if hasEmptyLocation {
		end, _ := solve(start, startx, starty)
		prn(end)
	} else {
		fmt.Println("There are no zeros in this grid")
	}

	fmt.Println()

	elapsed := time.Since(startTime)
	fmt.Println("Resolved in: ", elapsed)

}
