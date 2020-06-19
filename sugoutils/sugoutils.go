package sugoutils

import (
	"fmt"
	"sort"
)

var Calls int32

type Location struct {
	I        uint8
	J        uint8
	ValCount uint8
}

type byMostValues []Location

func (s byMostValues) Len() int {
	return len(s)
}
func (s byMostValues) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byMostValues) Less(i, j int) bool {
	return s[i].ValCount > s[j].ValCount
}

func StringToGrid(s []string) [9][9]uint8 {

	var res [9][9]uint8
	var i8 uint8

	for i := range s {
		bytes := []byte(s[i])
		for j := range bytes {
			i8 = uint8(bytes[j])
			if i8 == 32 {
				res[i][j] = 0
			} else {
				res[i][j] = i8 - 32 - 16
			}
		}
	}
	return res
}

func GridToString(s [9][9]uint8) string {

	res := ""
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] != 0 {
				res = res + " " + string(byte(48+s[i][j]))
			} else {
				res = res + " ."
			}
		}
		if i < 8 {
			res = res + "\n"
		}
	}
	return res
}

func locInit(s [9][9]uint8) []Location {
	var locs []Location
	var lineCount [9]uint8
	var colCount [9]uint8
	var curVal uint8
	var cellCount uint8

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			curVal = s[i][j]
			if curVal != 0 {
				lineCount[i]++
				colCount[j]++
			}
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			curVal = s[i][j]
			if curVal == 0 {
				cellCount = lineCount[i] + colCount[j]
				locs = append(locs, Location{uint8(i), uint8(j), cellCount})
			}
		}
	}

	sort.Sort(byMostValues(locs))
	return locs
}

func OptimEmptyLocations(s [9][9]uint8) []Location {

	// Prepare the Deterministic Location List (where there are zeros)
	// Going through this list in deterministic order is a mandatory aspect of backtracking
	// This is because once a branch has failed, you need to tray with the next value and
	// and the same Location order to try with new values
	// In the standard version of backtracking the deterministic order is the reading order

	return locInit(s)

}

func ReadEmptyLocation(s [9][9]uint8) []Location {

	// returns the list of i,j location if any, where there is no value yet (0 is no value yet)
	// the meaning of next is 'next by reading order' and is deterministic as needed

	var locs []Location

	for u := 0; u < 9; u++ {
		for v := 0; v < 9; v++ {
			if s[u][v] == 0 {
				locs = append(locs, Location{uint8(u), uint8(v), 0})
			}
		}
	}
	return locs
}

func IsValidSolution(s [9][9]uint8) bool {

	// checks if s is a valid, fully filled sudoku grid

	var lineSum, colSum uint8
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

func listAcceptableValues(s [9][9]uint8, i uint8, j uint8) []uint8 {

	// An acceptable value in i,j is a value absent in column i and line j

	res := make([]uint8, 0)
	if s[i][j] != 0 {
		fmt.Println("There should not be any value here")
	} else {
		var nextVal uint8 = 1
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

func SolveGrid(s [9][9]uint8) ([9][9]uint8, bool) {

	return Solve(s, OptimEmptyLocations(s))
}

func Solve(s [9][9]uint8, locs []Location) ([9][9]uint8, bool) {

	// This is a classic backtracking algorithm (algoritme des flots en Francais)

	var u, v uint8
	u = locs[0].I
	v = locs[0].J

	// This function is to be called on an available 0 location
	if s[u][v] == 0 {

		Calls++

		// If on a 0 let's list acceptable values
		options := listAcceptableValues(s, u, v)

		for _, option := range options {

			// for each acceptable value we create a new grid with the value ...
			s2 := s
			s2[u][v] = option

			// ... and for this grid we find the next 0 ...
			if len(locs) != 1 {
				// ... and we calculate the solution with the new grid and the next 0 location ...
				attempt, status := Solve(s2, locs[1:])
				if status {
					// ... and if it is a solution then we return this solution ...
					return attempt, status
				} else {
					// ... otherwise we continue to the next acceptable value in the list
					continue
				}
			} else {
				// There is no more zeros so we need to test the situation
				if IsValidSolution(s2) {
					return s2, true
				} else {
					return s2, false
				}
			}
		}
	}
	return s, false
}
