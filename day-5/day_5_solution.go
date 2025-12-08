package main

import ("os"; "io"; "fmt"; "strconv"; "bufio"; "strings")

func main() {
    pantry:=ParsePantry(os.Stdin)
    fmt.Println(pantry)

    count:=CountFreshIngredients(pantry)
    fmt.Println("Day 5 Part 1 answer:",count)

    count =CountAllPossibleFresh(pantry.fresh_ranges)
    fmt.Println("Day 5 Part 2 answer:",count)
}
type Fresh struct {
    min	int
    max	int
}
type Pantry struct {
    fresh_ranges	[]Fresh
    ingredients		[]int
}
func ParsePantry(in io.Reader) Pantry {
    scanner:=bufio.NewScanner(in)
    var fresh_ranges []Fresh
    var ingredients []int
    for scanner.Scan() {
	line:=string(scanner.Text())
	if strings.Contains(line,"-") {
	    mn,mx,_:=strings.Cut(line,"-")
	    min,er:=strconv.Atoi(mn); if er!=nil { panic(er) }
	    max,er:=strconv.Atoi(mx); if er!=nil { panic(er) }
	    fresh_ranges = append(fresh_ranges,Fresh{min,max})
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
func CountFreshIngredients(pantry Pantry) int {
    var count 	int
    for i:=range len(pantry.ingredients) {
	count += CheckFresh(pantry.fresh_ranges,pantry.ingredients[i])
    }
    return count
}
func CheckFresh(fresh_ranges []Fresh, ingredient int) int {
    for _,fresh_range:=range fresh_ranges {
	if ingredient<=fresh_range.max && fresh_range.min<=ingredient {
	    return 1
	}
    }
    return 0
}
func CountAllPossibleFresh(fresh_ranges []Fresh) int {
    var count 	int
    for _,fresh_range:=range fresh_ranges {
	ids_in_range:=(fresh_range.max - fresh_range.min) + 1
	fmt.Println(ids_in_range)
	count += ids_in_range
    }
    return count
}
