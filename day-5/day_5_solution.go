package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"sort"
)

func main() {
	pantry:=ParsePantry(os.Stdin)
	// fmt.Println(pantry.fresh_ranges)

	count:=CountFreshIngredients(pantry)
	fmt.Println("Day 5 Part 1 answer:",count)

	sfrs:=MergeFreshRanges(pantry.fresh_ranges)
	// fmt.Println("sorted fresh ranges",sfrs)

	count =CountAllPossibleFresh(sfrs)
	fmt.Println("Day 5 Part 2 answer:",count)
}
type FreshRange struct {
	min	int
	max	int
}
type Pantry struct {
	fresh_ranges	[]FreshRange
	ingredients		[]int
}
func ParsePantry(in io.Reader) Pantry {
	scanner:=bufio.NewScanner(in)
	var fresh_ranges []FreshRange
	var ingredients []int
	for scanner.Scan() {
		line:=string(scanner.Text())
		if strings.Contains(line,"-") {
				mn,mx,_:=strings.Cut(line,"-")
				min,er:=strconv.Atoi(mn); if er!=nil { panic(er) }
				max,er:=strconv.Atoi(mx); if er!=nil { panic(er) }
				fresh_ranges = append(fresh_ranges,FreshRange{min,max})
		} else if line!="" {
				ig,er:=strconv.Atoi(line); if er!=nil { panic(er) }
				ingredients = append(ingredients,ig)
		} else {
				continue
		}
	}
	er:=scanner.Err(); if er!=nil { panic(er) }
	return Pantry{fresh_ranges,ingredients}
}
func MergeFreshRanges(fresh_ranges []FreshRange) []FreshRange {
	i:=0
	for {
		sort.Slice(fresh_ranges, func(i,j int) bool {
			return fresh_ranges[i].min < fresh_ranges[j].min
		})

		if i+1>=len(fresh_ranges) { break }

		first:=fresh_ranges[i  ]
		next :=fresh_ranges[i+1]
		if first.min<next.min && next.min<first.max {
			new_fr:=FreshRange{first.min,first.max}
			if first.max<next.max {
				new_fr = FreshRange{first.min,next.max}
			}
			fresh_ranges = slices.Delete(fresh_ranges,i,i+1)
			fresh_ranges = slices.Delete(fresh_ranges,i,i+1)
			fresh_ranges = append(fresh_ranges,new_fr)
		} else { 
			i++
		}
	}
	return fresh_ranges
}
func CountFreshIngredients(pantry Pantry) int {
	var count 	int
	for i:=range len(pantry.ingredients) {
		count += CheckFresh(pantry.fresh_ranges,pantry.ingredients[i])
	}
	return count
}
func CheckFresh(fresh_ranges []FreshRange, ingredient int) int {
	for _,fresh_range:=range fresh_ranges {
		if ingredient<=fresh_range.max && fresh_range.min<=ingredient {
				return 1
		}
	}
	return 0
}
func CountAllPossibleFresh(fresh_ranges []FreshRange) int {
	var count int
	for _,fresh_range:=range fresh_ranges {
		if fresh_range.min!=0 && fresh_range.max!=0 {
			count += fresh_range.max-fresh_range.min+1
		}
	}
	return count 
}
// func CountAllPossibleFresh(fresh_ranges []FreshRange) int {
// 	// var fresh   []int
// 	fresh := make([]int,100000)
// 	i := 0
// 	for _,fresh_range:=range fresh_ranges {
// 		id:=fresh_range.min
// 		for range fresh_range.max-fresh_range.min+1 {
// 			if !slices.Contains(fresh,id) {
// 				fresh[i] = id
// 				i++
// 			}
// 			id++
// 		}
// 	}
// 	fmt.Println(i)
// 	var count int
// 	for j:=range len(fresh) {
// 		if fresh[j]==0 {
// 			break
// 		} else {
// 			count++
// 		}
// 	}
// 	return count 
// }
