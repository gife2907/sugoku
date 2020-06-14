package main

/*
	=== to do ===

	github
	launch.js and popup
	dble click on message opens editor
	ut
	package
	speed orptimization
	more go like w code review
	samples
	log appropriately
	api
	service
	vr
	ar
	online Pi service

*/

import (
	"fmt"
	"math/rand"
)

func getWinning() [9][9]int8 {

	var gs []string
	/*
			gs = append(gs, "123456789")
			gs = append(gs, "234567891")
			gs = append(gs, "345678912")
			gs = append(gs, "456789123")
			gs = append(gs, "567891234")
			gs = append(gs, "678912345")
			gs = append(gs, "789123456")
			gs = append(gs, "891234567")
			gs = append(gs, "912345678")

			gs = append(gs, "123456789")
		gs = append(gs, "230067891")
		gs = append(gs, "345678912")
		gs = append(gs, "456709023")
		gs = append(gs, "567891234")
		gs = append(gs, "678912345")
		gs = append(gs, "789123456")
		gs = append(gs, "891034500")
		gs = append(gs, "910305678")
	*/

	gs = append(gs, "000000000")
	gs = append(gs, "000000000")
	gs = append(gs, "345678912")
	gs = append(gs, "000000000")
	gs = append(gs, "000000000")
	gs = append(gs, "678912345")
	gs = append(gs, "000000000")
	gs = append(gs, "891034500")
	gs = append(gs, "910305678")

	var a [9][9]int8
	var line []int8
	var b int8

	for i := range gs {
		bytes := []byte(gs[i])

		//line = []int8(bline)
		for j := range bytes {
			b = bytes[j]
			if b == 32 {
				a[i][j] = 0
			} else {
				a[i][j] = b - 32 - 16
			}
		}
	}

	return a
}

func prn(s [9][9]int8) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Println(" ", s[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func getSudo00() [9][9]int8 {

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

	var a [9][9]int8
	var line []int8
	var b int8

	for i := range gs {
		line = []int8(gs[i])
		for j := range line {
			b = line[j]
			if b == 32 {
				a[i][j] = 0
			} else {
				a[i][j] = b - 32 - 16
			}
		}
	}

	return a
}

func proofLine(l [9]int8) bool {

	var sum int8
	for i := 0; i < 9; i++ {
		sum += l[i]
	}
	return (sum == 45)
}

func sumLine(l [9]int8) int8 {

	var sum int8
	for i := 0; i < 9; i++ {
		sum += l[i]
	}
	return sum
}

func winningSudo(s [9][9]int8) bool {
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

func nextOkValue(s [9][9]int8, i int, j int) int8 {
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

func listOkValues(s [9][9]int8, i int, j int) []int8 {
	res := make([]int8, 0)
	var nextVal int8 = 1
	for nextVal < 10 {
		for v := 0; v < 9; v++ {
			if s[i][v] == nextVal {
				nextVal++
				v = 0
				continue
			}
			if s[v][j] == nextVal {
				nextVal++
				v = 0
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
	return res
}

func nextZero(s [9][9]int8, i int, j int) (int, int, bool) {
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

/*
func nextSudo(s int8 [9][9], location int, iterateThisLocation bool ) int8 [9][9] {

}

func nextSudo() func int8[9][9] {
	s := start
	location := 0
	iterateThisLocation := false

	return func() int8[9][9] {

		for {
			u, v := divmod( location, 9 )
			if iterateThisLocation {
				nextVal := s[u][v]+1
				for nextVal <= 10 {
					for i :=0; i<9; i++ {
						if ( s[u][i] == nextVal || s[i][v] == nextVal ) {
							nextVal++
						}
					}
				}
			}
			if nextVal == 10 {
				location++
			} else {
				s[u][v] = nextVal
				return s
			}
		}
	}
}

*/

func solve(s [9][9]int8, i int, j int) ([9][9]int8, bool) {

	if rand.Intn(5000000) == 10 {
		prn(s)
	}

	// Do we have the winning grid ?
	if winningSudo(s) {
		fmt.Println("Winner found up")
		return s, true
	}

	// If not we are in i, j and there is a 0, so let's try with the next available option
	if s[i][j] == 0 {
		options := listOkValues(s, i, j)

		for _, option := range options {
			s2 := s
			s2[i][j] = option
			u, v, ok := nextZero(s2, i, j)
			if ok {
				// There is a next zero
				attempt, status := solve(s2, u, v)
				if status {
					return attempt, status
				} else {
					continue
				}
			} else {
				// There is no more zeros so we need to fail ????
				if winningSudo(s2) {
					fmt.Println("Winner found down")
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

	//start := getSudo00()
	start := getWinning()

	fmt.Println(start)
	fmt.Println(nextZero(start, 0, 0))

	startx, starty, hasZeros := nextZero(start, 0, 0)
	if hasZeros {
		end, _ := solve(start, startx, starty)
		prn(end)
	} else {
		fmt.Println("There are no zeros in this grid")
	}

}
