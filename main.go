package main

/*
	=== infos ===

	Gilles Ferrero
	https://github.com/gife2907/sugoku

	=== to do ===

	v	solves a real grid
	v	github
	v	byte to uint8 conv
	v	some more test samples
	v	time measurement
	v	gopath issues & package : go help gopath, https://www.golang-book.com/books/intro/11
	v	dble click on message opens editor
	v	uuint8
	v	struct
	v 	resolution speed and memory optimization
	v	ut try
	v	clean solve with a function passed to return next 0
	v	bench
	launch.js and popup
	builtin : https://gobyexample.com/sorting-by-functions - the builtin []string type -
	more go like code w code review
	more samples with failing cases
	log appropriately
	ut
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

func main() {

	var grids [][9][9]uint8
	grids = append(grids, sugoutils.Grid00())
	grids = append(grids, sugoutils.Grid01())
	grids = append(grids, sugoutils.Grid02())
	grids = append(grids, sugoutils.Grid03())
	grids = append(grids, sugoutils.Grid99())

	startTime := time.Now()
	elapsed := time.Since(startTime)

	for i := range grids {
		startTime = time.Now()
		sugoutils.Calls = 0
		sugoutils.SolveGrid(grids[i])
		elapsed = time.Since(startTime)
		fmt.Println("Optimized - Resolved in ", sugoutils.Calls, " calls. That's ", elapsed)
		startTime = time.Now()
		sugoutils.Calls = 0
		sugoutils.Solve(grids[i], sugoutils.ReadEmptyLocation(grids[i]))
		elapsed = time.Since(startTime)
		fmt.Println("Reading   - Resolved in ", sugoutils.Calls, " calls. That's ", elapsed)
		fmt.Println()
	}

}
