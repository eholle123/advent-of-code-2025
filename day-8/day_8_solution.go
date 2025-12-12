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
	coords:=ParseCoords(os.Stdin)
	fmt.Println("Coords: (x, y, z):", coords)

	sorted_coords:=SortByMag(coords)
	fmt.Println(sorted_coords)
}
type Coord struct {
	x 	int
	y 	int
	z 	int
	mag int
}
func ParseCoords(in io.Reader) []Coord {
	scanner:=bufio.NewScanner(in)
	var coords 	[]Coord
	var x,y,z		int
	var er 			error
	for scanner.Scan() {
		line:=string(scanner.Text())
		s   :=strings.Split(line,",")
		x,er = strconv.Atoi(s[0]); if er!=nil { panic(er) }
		y,er = strconv.Atoi(s[1]); if er!=nil { panic(er) }
		z,er = strconv.Atoi(s[2]); if er!=nil { panic(er) }
		mag := x^2 + y^2 + z^2
		coords = append(coords,Coord{x,y,z,mag})
	}
	er = scanner.Err(); if er!=nil { panic(er) }
	return coords
}
func SortByMag(coords []Coord) []Coord {
	sort.Slice(coords, func(i,j int) bool {
		return coords[i].mag < coords[j].mag
	})
	m:=coords[0].mag
	for i:=range len(coords) {
		coords[i].mag = coords[i].mag - m
	}
	return coords
}
func Distance(c1,c2 Coord) int {
	return (c2.x-c1.x)^2 + (c2.y-c1.y)^2 + (c2.z-c1.z)^2
}
