package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	pts:=ParsePoints(os.Stdin)
	fmt.Println("points:",pts)

	fmt.Println("Day 9 Part 1 answer:")
}

type Point struct {
	row	int
	col int
}

type Tiles struct {
	red_tiles []Point
	max_row   int
	max_col   int
	// grid      []string
	ul        Point
	ur        Point
	dl        Point
	dr        Point

}

func ParsePoints(in io.Reader) []Point {
	scanner:=bufio.NewScanner(in)
	var pts 		[]Point
	var row,col int
	var er 			error
	for scanner.Scan() {
		line:=string(scanner.Text())
		s   :=strings.Split(line,",")
		col,er = strconv.Atoi(s[0]); if er!=nil { panic(er) }
		row,er = strconv.Atoi(s[1]); if er!=nil { panic(er) }
		pts = append(pts,Point{row,col})
	}
	er = scanner.Err(); if er!=nil { panic(er) }
	sort.Slice(pts, func(i,j int) bool {
		return pts[i].row < pts[j].row
	})
	return pts
}

func GetCornerTiles(pts []Point) Tiles {

	sort.Slice(pts, func(i,j int) bool {
		return pts[i].row < pts[j].row
	})
	ul:=pts[0]

}
