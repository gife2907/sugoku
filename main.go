package main

/*
	=== to do ===

	v	github
	v	byte to int8
	v	some more samples
	v	time measurement
	gopath issues
	launch.js and popup
	dble click on message opens editor
	ut
	package
	speed optimization
	more go like w code review
	more samples with failing cases
	log appropriately
	api
	service
	vr
	ar
	online Pi service

*/

import (
	"fmt"
	"time"
)

func getGrid00() [9][9]int8 {

	var gs []string
	gs = append(gs, "53  7    ")
	gs = append(gs, "6  195   ")
	gs = append(gs, " 98    6 ")
	gs = append(gs, "8   6   3")
	gs = append(gs, "4  8 3  1")
	gs = append(gs, "7   2   6")
	gs = append(gs, " 6    28 ")
	gs = append(gs, "   419  5")
	gs = append(gs, "    8  79")
	return strToGrid(gs)
}

func getGrid01() [9][9]int8 {

	var gs []string
	gs = append(gs, "8  4  91 ")
	gs = append(gs, "  3      ")
	gs = append(gs, "     3  4")
	gs = append(gs, "     1 4 ")
	gs = append(gs, " 58   7  ")
	gs = append(gs, " 7   68  ")
	gs = append(gs, "     2   ")
	gs = append(gs, "      16 ")
	gs = append(gs, "91  6 5  ")
	return strToGrid(gs)
}

func getGrid99() [9][9]int8 {

	// a grid designed against brute forcing
	// https://en.wikipedia.org/wiki/Sudoku_solving_algorithms
	// note: it takes 75ms to solve instead of 1 or 2ms
	// it would deserve an optimization algorithm some day to resolve to 1 or 2 ms as well
	// optimisation factor could be to divide time by 40 to 70
	var gs []string
	gs = append(gs, "         ")
	gs = append(gs, "     3 85")
	gs = append(gs, "  1 2    ")
	gs = append(gs, "   5 7   ")
	gs = append(gs, "  4   1  ")
	gs = append(gs, " 9       ")
	gs = append(gs, "5      73")
	gs = append(gs, "  2 1    ")
	gs = append(gs, "    4   9")
	return strToGrid(gs)
}

func getGridFullAndValid() [9][9]int8 {

	// This grid is full and valid
	var gs []string
	gs = append(gs, "123456789")
	gs = append(gs, "234567891")
	gs = append(gs, "345678912")
	gs = append(gs, "456789123")
	gs = append(gs, "567891234")
	gs = append(gs, "678912345")
	gs = append(gs, "789123456")
	gs = append(gs, "891234567")
	gs = append(gs, "912345678")
	return strToGrid(gs)
}

func getGridManyHoles() [9][9]int8 {

	// This grid has many holes and is based on the Full Valid grid
	var gs []string
	gs = append(gs, "         ")
	gs = append(gs, "         ")
	gs = append(gs, "345678912")
	gs = append(gs, "         ")
	gs = append(gs, "         ")
	gs = append(gs, "678912345")
	gs = append(gs, "         ")
	gs = append(gs, "891 345  ")
	gs = append(gs, "91 3 5678")
	return strToGrid(gs)
}

func getGridSimple() [9][9]int8 {

	// This grid has only a few holes and is based on the Full Valid grid
	var gs []string
	gs = append(gs, "123456789")
	gs = append(gs, "23  67891")
	gs = append(gs, "345678912")
	gs = append(gs, "4567 9 23")
	gs = append(gs, "567891234")
	gs = append(gs, "678912345")
	gs = append(gs, "789123456")
	gs = append(gs, "891 345  ")
	gs = append(gs, "91 3 5678")
	return strToGrid(gs)
}

func strToGrid(s []string) [9][9]int8 {

	var res [9][9]int8
	var i8 int8

	for i := range s {
		bytes := []byte(s[i])
		for j := range bytes {
			i8 = int8(bytes[j])
			if i8 == 32 {
				res[i][j] = 0
			} else {
				res[i][j] = i8 - 32 - 16
			}
		}
	}

	return res
}

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
	var sl, sc int8
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == 0 {
				return false
			} else {
				sl += s[i][j]
				sc += s[j][i]
			}
		}
		if sl != 45 || sc != 45 {
			return false
		} else {
			sl = 0
			sc = 0
		}
	}
	return true
}

func nextAcceptableValue(s [9][9]int8, i int, j int) int8 {
	var nextVal int8 = 1
	for nextVal < 10 {
		for v := 0; v < 9; v++ {
			if s[i][v] == nextVal {
				fmt.Println("found ", nextVal, " in line ", i)
				// if a line or a column already has a value for nextVal, we move to try next one
				nextVal++
				v = 0
				continue
			}
			if s[v][j] == nextVal {
				fmt.Println("found ", nextVal, " in column ", j)
				// if a line or a column already has a value for nextVal, we move to try next one
				nextVal++
				v = 0
				continue
			}
		}
		if nextVal != 10 {
			break
		}
	}
	return nextVal
}

func listAcceptableValues(s [9][9]int8, i int, j int) []int8 {

	res := make([]int8, 0)
	if s[i][j] != 0 {
		fmt.Println("There should not be any value here")
	} else {
		var nextVal int8 = 1
		for nextVal < 10 {
			for v := 0; v < 9; v++ {
				if s[i][v] == nextVal {
					nextVal++
					v = -1
					continue
				}
				if s[v][j] == nextVal {
					nextVal++
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
					// ... and if it is the solution then we return this solution ...
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

	start := getGrid01()

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
