/*
Assumes you have a PDF with A6-sized pages. With those, it calculates the new order for these A6 pages in your PDF so that you can print it out using 4x A6 per A4, front and back. After printing the A6 pages in the new order, cut your A4 paper stack horizontally and put the top halves into the center of the bottom halves. Fold the new stack vertically and you have an A6 booklet.

Urheberrecht and Copyright 2018 Ernst Rohlicek.
License: GPLv3 or (at your option) any later version, see https://www.gnu.org/licenses/gpl-3.0.html
*/

package main

import (
	"fmt"
	"flag"
)

type SheetA4 struct {
	FrontPhysical [4]int
	BackPhysical [4]int
}

type SheetA5Horizontal struct {
	FrontPhysical [2]int
	BackPhysical [2]int
}

func main() {
	// parse flags
	var help, debug bool
	var inpath, outpath string
	var nPages int
	flag.BoolVar(&help, "h", false, "print usage information")
	flag.BoolVar(&debug, "debug", false, "give detailed event output")
	flag.StringVar(&inpath, "in", "", "input PDF with A6 pages")
	flag.StringVar(&outpath, "out", "", "output PDF path")
	flag.IntVar(&nPages, "pages", 0, "number of pages - must be multiple of 8")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}
	// checks
	if inpath == "" {
		fmt.Println("ERROR: need input PDF path")
		flag.PrintDefaults()
		return
	}
	if outpath == "" {
		fmt.Println("ERROR: need output PDF path")
		flag.PrintDefaults()
		return
	}
	if nPages == 0 {
		fmt.Println("ERROR: need number of pages")
		flag.PrintDefaults()
		return
	}
	if nPages % 8 != 0 {
		fmt.Println("ERROR: number of pages must be multiple of 8")
		flag.PrintDefaults()
		return
	}

	// create physical A4 sheets and number the physical A6 pages on them
	nSheets := nPages / 8;
	sheetsA4 := make([]SheetA4, nSheets)
	curPage := 1
	// number all A4 sheets
	for n := 0 ; n < nSheets; n++ {
		// front
		sheetsA4[n].FrontPhysical[0] = curPage
		curPage++
		sheetsA4[n].FrontPhysical[1] = curPage
		curPage++
		sheetsA4[n].FrontPhysical[2] = curPage
		curPage++
		sheetsA4[n].FrontPhysical[3] = curPage
		curPage++

		// back
		sheetsA4[n].BackPhysical[0] = curPage
		curPage++
		sheetsA4[n].BackPhysical[1] = curPage
		curPage++
		sheetsA4[n].BackPhysical[2] = curPage
		curPage++
		sheetsA4[n].BackPhysical[3] = curPage
		curPage++
	}

	// virtually rip in half horizontally
	sheetsA5 := make([]SheetA5Horizontal, nSheets*2)
	// top half
	for n := 0 ; n < nSheets; n++ {
		// front
		sheetsA5[n].FrontPhysical[0] = sheetsA4[n].FrontPhysical[0]
		sheetsA5[n].FrontPhysical[1] = sheetsA4[n].FrontPhysical[1]

		// back
		sheetsA5[n].BackPhysical[0] = sheetsA4[n].BackPhysical[0]
		sheetsA5[n].BackPhysical[1] = sheetsA4[n].BackPhysical[1]
	}
	// bottom half
	for n := 0 ; n < nSheets; n++ {
		// front
		sheetsA5[nSheets+n].FrontPhysical[0] = sheetsA4[n].FrontPhysical[2]
		sheetsA5[nSheets+n].FrontPhysical[1] = sheetsA4[n].FrontPhysical[3]

		// back
		sheetsA5[nSheets+n].BackPhysical[0] = sheetsA4[n].BackPhysical[2]
		sheetsA5[nSheets+n].BackPhysical[1] = sheetsA4[n].BackPhysical[3]
	}

	// debug print for virtual ripping
	if debug {
		fmt.Println("virtual horizontal A5 sheets after cutting horizontally:")
		for n := 0; n < nSheets*2; n++ {
			fmt.Printf("front %v back %v\n", sheetsA5[n].FrontPhysical, sheetsA5[n].BackPhysical)
		}
	}

	// number all physical pages in book order
	curPage = 0	// index actually
	// left pages of the stack going up = first half book pages
	bookOrder := make([]int, nPages)	// this is a translation table, i.e. book page 1 belongs to nth position
	for n := 0 ; n < len(sheetsA5); n++ {
		// back right
		bookOrder[curPage] = sheetsA5[n].BackPhysical[1]
		curPage++
		// front left
		bookOrder[curPage] = sheetsA5[n].FrontPhysical[0]
		curPage++
	}
	// right pages of the stack going down = second half book pages
	for n := len(sheetsA5)-1 ; n >= 0; n-- {
		// front right
		bookOrder[curPage] = sheetsA5[n].FrontPhysical[1]
		curPage++
		// back left
		bookOrder[curPage] = sheetsA5[n].BackPhysical[0]
		curPage++
	}

	// apply translation
	printOrder := make([]int, nPages)
	for index, pgNumber := range bookOrder {
		printOrder[pgNumber-1] = index+1
	}

	// print result
	if debug {
		fmt.Println("print order:", printOrder)
	}

	// print pdftk command-line
	fmt.Printf("pdftk A=%s cat", inpath)
	for _, pgNumber := range printOrder {
		fmt.Printf(" A%d", pgNumber)
	}
	fmt.Printf(" output %s\n", outpath)
}
