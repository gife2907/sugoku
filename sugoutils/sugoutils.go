package sugoutils

func Grid00() [9][9]int8 {

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
	return stringToGrid(gs)
}

func Grid01() [9][9]int8 {

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
	return stringToGrid(gs)
}

func Grid99() [9][9]int8 {

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
	return stringToGrid(gs)
}

func GridFullAndValid() [9][9]int8 {

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
	return stringToGrid(gs)
}

func GridManyHoles() [9][9]int8 {

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
	return stringToGrid(gs)
}

func GridFewHoles() [9][9]int8 {

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
	return stringToGrid(gs)
}

func stringToGrid(s []string) [9][9]int8 {

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
