package sugoutils

func Grid00() [9][9]uint8 {

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
	return StringToGrid(gs)
}

func Grid01() [9][9]uint8 {

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
	return StringToGrid(gs)
}

func Grid02() [9][9]uint8 {

	var gs []string
	gs = append(gs, "5    2   ")
	gs = append(gs, "     3 9 ")
	gs = append(gs, "2 84 6   ")
	gs = append(gs, " 9 53    ")
	gs = append(gs, " 65     8")
	gs = append(gs, " 8    9 3")
	gs = append(gs, "6 4   7  ")
	gs = append(gs, "      21 ")
	gs = append(gs, "   274   ")
	return StringToGrid(gs)
}

func Grid03() [9][9]uint8 {

	var gs []string
	gs = append(gs, "85   24  ")
	gs = append(gs, "72      9")
	gs = append(gs, "  4      ")
	gs = append(gs, "   1 7  2")
	gs = append(gs, "3 5   9  ")
	gs = append(gs, " 4       ")
	gs = append(gs, "    8  7 ")
	gs = append(gs, " 17      ")
	gs = append(gs, "    36 4 ")
	return StringToGrid(gs)
}

func GridNoIntegrity01() [9][9]uint8 {

	var gs []string
	gs = append(gs, " 923     ")
	gs = append(gs, "6  3 2   ")
	gs = append(gs, " 6 5 1 9 ")
	gs = append(gs, "  5   987")
	gs = append(gs, "3 6 7    ")
	gs = append(gs, "1     3 8")
	gs = append(gs, "6  3 2   ")
	gs = append(gs, "    3  9 ")
	gs = append(gs, "45     3 ")
	return StringToGrid(gs)
}

func GridNoIntegrity02() [9][9]uint8 {

	var gs []string
	gs = append(gs, " 237    6")
	gs = append(gs, "  84   3 ")
	gs = append(gs, " 98 1    ")
	gs = append(gs, "  247  58")
	gs = append(gs, "4     8 5")
	gs = append(gs, " 2 3     ")
	gs = append(gs, "1     7 9")
	gs = append(gs, "4     3  ")
	gs = append(gs, "       71")
	return StringToGrid(gs)
}

func GridNoIntegrity03() [9][9]uint8 {

	var gs []string
	gs = append(gs, "6  3 2   ")
	gs = append(gs, " 47 8   1")
	gs = append(gs, "      8 1")
	gs = append(gs, "38 6     ")
	gs = append(gs, "   5     ")
	gs = append(gs, " 2       ")
	gs = append(gs, " 8  4    ")
	gs = append(gs, "  8 9 1  ")
	gs = append(gs, "4     5 8")
	return StringToGrid(gs)
}

func Grid07() [9][9]uint8 {

	var gs []string
	gs = append(gs, "1     3 8")
	gs = append(gs, "1    6 8 ")
	gs = append(gs, "249 6   3")
	gs = append(gs, "   8    9")
	gs = append(gs, "   5 1   ")
	gs = append(gs, "      8 1")
	gs = append(gs, " 476   5 ")
	gs = append(gs, "     7 95")
	gs = append(gs, " 4 5     ")
	return StringToGrid(gs)
}

func Grid08() [9][9]uint8 {

	var gs []string
	gs = append(gs, " 834     ")
	gs = append(gs, "  9     3")
	gs = append(gs, " 26 39   ")
	gs = append(gs, "2 3 8    ")
	gs = append(gs, "6  3 2   ")
	gs = append(gs, "1     9  ")
	gs = append(gs, "         ")
	gs = append(gs, " 2    593")
	gs = append(gs, "9 4  5   ")
	return StringToGrid(gs)
}

func Grid09() [9][9]uint8 {

	var gs []string
	gs = append(gs, "   52    ")
	gs = append(gs, "53  2 9  ")
	gs = append(gs, "1    786 ")
	gs = append(gs, "    5   1")
	gs = append(gs, " 47 2    ")
	gs = append(gs, "      94 ")
	gs = append(gs, " 2      6")
	gs = append(gs, "1     3 8")
	gs = append(gs, "2    1 9 ")
	return StringToGrid(gs)
}

func Grid10() [9][9]uint8 {

	var gs []string
	gs = append(gs, "  7  8   ")
	gs = append(gs, "   36    ")
	gs = append(gs, "34 6     ")
	gs = append(gs, "      4 1")
	gs = append(gs, " 4  5  67")
	gs = append(gs, "       4 ")
	gs = append(gs, "8  7    4")
	gs = append(gs, " 8   4 5 ")
	gs = append(gs, "    7  8 ")
	return StringToGrid(gs)
}

func Grid11() [9][9]uint8 {

	var gs []string
	gs = append(gs, "      8 1")
	gs = append(gs, " 2       ")
	gs = append(gs, " 52  68  ")
	gs = append(gs, "    1 78 ")
	gs = append(gs, "1       3")
	gs = append(gs, "4   7 1  ")
	gs = append(gs, "      8 1")
	gs = append(gs, "963      ")
	gs = append(gs, "15 3     ")
	return StringToGrid(gs)
}

func Grid12() [9][9]uint8 {

	var gs []string
	gs = append(gs, "         ")
	gs = append(gs, "    75   ")
	gs = append(gs, "6     7 3")
	gs = append(gs, "    6   4")
	gs = append(gs, " 32     5")
	gs = append(gs, "   5 3   ")
	gs = append(gs, " 5 3 7 4 ")
	gs = append(gs, "  5  8  1")
	gs = append(gs, "         ")
	return StringToGrid(gs)
}

func Grid99() [9][9]uint8 {

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
	return StringToGrid(gs)
}

// GridFullAndValid() returns a grid that is full and valid
func GridFullAndValid() [9][9]uint8 {

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
	return StringToGrid(gs)
}

func GridManyHoles() [9][9]uint8 {

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
	return StringToGrid(gs)
}

func GridFewHoles() [9][9]uint8 {

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
	return StringToGrid(gs)
}
