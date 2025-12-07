package main

import ("os"; "io"; "fmt"; "bufio"; "strings")

func main() {
    grid:=ParseGridMatrix(os.Stdin)
    PrintGridMatrix(grid.m)
    
    // total:=CountAccessableRolls(grid.m,4)
    // fmt.Println("\nDay 4 Part 1 answer:",total)
    // PrintGridMatrix(grid.m)

    // grid.m = ResetGrid(grid.m)
    // PrintGridMatrix(grid.m)

    // t2:=CountAccessableRolls(grid.m,4)
    // fmt.Println("\nt2:",t2)
    // PrintGridMatrix(grid.m)

    // grid.m = ResetGrid(grid.m)
    // PrintGridMatrix(grid.m)
    total:=CountRemovableRolls(grid.m,4)
    fmt.Println("Day 4 Part 2 answer:",total)
}
type Roll struct {
    value	string
    can_access  bool
}
type GridMatrix struct {
    m		[][]Roll
    max_rows	int
    max_cols	int
}
func ParseGridMatrix(in io.Reader) GridMatrix {
    scanner:=bufio.NewScanner(in)
    var grid [][]Roll
    var max_rows, max_cols int
    for scanner.Scan() {
	row_data:=scanner.Text()
	max_cols = len(row_data)
	var row_rolls []Roll
	for _,s:=range row_data {
	    value:=string(s)
	    row_rolls = append(row_rolls, Roll{value,false})
	}
	grid = append(grid, row_rolls)
	max_rows++
    }
    return GridMatrix{grid,max_rows,max_cols}
}
func CountRemovableRolls(grid [][]Roll, max_around int) int {
    var total int
    remaining_accessable:=1
    for remaining_accessable!=0 {
	remaining_accessable = CountAccessableRolls(grid,max_around)
	total = total + remaining_accessable
	PrintGridMatrix(grid)
	grid = ResetGrid(grid)
	PrintGridMatrix(grid)
    }
    return total
}
func CountAccessableRolls(grid [][]Roll, max_around int) int {
    var access_total	int
    for i:=range len(grid) {
	for j:=range len(grid[i]) {
	    if grid[i][j].value=="@" {
		can_access:=CanAccessRoll(i,j,grid,max_around) 
		if can_access {
		    access_total++
		    grid[i][j].value = "x"
		    grid[i][j].can_access = can_access
		}
	    }
	}
    }
    return access_total
}
func CanAccessRoll(i int, j int, grid [][]Roll, max_around int) bool {
    ul:="."
    u :="."
    ur:="."
    l :="."
    r :="."
    dl:="."
    d :="."
    dr:="."
    max_i:=len(grid)-1
    max_j:=len(grid[i])-1
    if (i==0 && j==0) || (i==0 && j==max_j) || (i==max_i && j==0) || (i==max_i && j==max_j) {
	return true
    } else if i==0 {
	l  = grid[i  ][j-1].value
        r  = grid[i  ][j+1].value
	dl = grid[i+1][j-1].value
        d  = grid[i+1][j  ].value
        dr = grid[i+1][j+1].value 
    } else if j==0 {
	u  = grid[i-1][j  ].value
	ur = grid[i-1][j+1].value
        r  = grid[i  ][j+1].value
        d  = grid[i+1][j  ].value
        dr = grid[i+1][j+1].value 
    } else if i==max_i {
	ul = grid[i-1][j-1].value
	u  = grid[i-1][j  ].value
	ur = grid[i-1][j+1].value
	l  = grid[i  ][j-1].value
        r  = grid[i  ][j+1].value
    } else if j==max_j {
	ul = grid[i-1][j-1].value
	u  = grid[i-1][j  ].value
	l  = grid[i  ][j-1].value
	dl = grid[i+1][j-1].value
        d  = grid[i+1][j  ].value
    } else {
	ul = grid[i-1][j-1].value
	u  = grid[i-1][j  ].value
	ur = grid[i-1][j+1].value
	l  = grid[i  ][j-1].value
        r  = grid[i  ][j+1].value
	dl = grid[i+1][j-1].value
        d  = grid[i+1][j  ].value
        dr = grid[i+1][j+1].value 
    }
    a_count:=strings.Count(ul+u+ur+l+r+dl+d+dr,"@") 
    x_count:=strings.Count(ul+u+ur+l+r+dl+d+dr,"x") 
    if a_count+x_count < max_around {
        return true
    }
    return false
}
func ResetGrid(grid [][]Roll) [][]Roll {
    for i:=range len(grid) {
	for j:=range len(grid[i]) {
	    if grid[i][j].value=="x" {
		grid[i][j].value = "."
	    }
	}
    }
    return grid
}
func PrintGridMatrix(grid [][]Roll) {
    fmt.Println()
    var i int
    for x:=range len(grid) {
	for y:=range len(grid[x]) {
	    if i==len(grid) {
		i = 1
		fmt.Println()
	    } else {
		i++
	    }
	    fmt.Print(grid[x][y].value)
	}
    }
    fmt.Println()
}
