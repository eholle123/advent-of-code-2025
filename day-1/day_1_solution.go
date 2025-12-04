package main

import ("os"; "fmt"; "bufio"; "io"; "strconv")

func main() {
    lines,er:=ParseInput(os.Stdin); if er!=nil { panic(er) }
    // fmt.Println(lines)
    rotations:=CalcLockRotations(lines)
    // fmt.Println(rotations)
    combo:=CalcLockCombo(rotations)
    fmt.Println(combo)
}

func ParseInput(in io.Reader) ([]string,error) {
    scanner:=bufio.NewScanner(in)
    var line  string
    var lines []string
    for scanner.Scan() {
	line = scanner.Text()
	lines = append(lines, line)
    }
    if err:=scanner.Err(); err!=nil { fmt.Fprintln(os.Stderr, "reading standard input:", err) }
    return lines,nil
}

func CalcLockRotations(lines []string) []int {
    var rotations []int
    for _,line:=range lines {
	if line!="" {
	    num,er:=strconv.Atoi(line[1:]); if er!=nil { panic(er) }
	    if string(line[0])=="L" {
	        num = num * -1
	    } else if string(line[0])=="R" {
	        num = num * 1
	    } else {}
	    rotations = append(rotations, num)
        }
    }
    return rotations
}

func CalcLockCombo(rotations []int) int {
    var zeros int
    start:=50
    sum:=start
    fmt.Println("start", sum)
    for _,num:=range rotations {
//  	new_sum:=sum+num
//  	if new_sum==0 || new_sum==100 {
//  	    zeros += 1
//  	    sum = 0
//  	}
//  	if new_sum < 0 {
//  	    sum = 100 + new_sum
//  	} else if 100 <= new_sum {
//  	    sum = new_sum - 100
// 	} else if new_sum==100 || new_sum==0 {
// 	    zeros += 1
// 	    sum = 0
//  	} else {
//  	    sum = new_sum
//  	}
	sum = (sum + num) % 100
 	if sum==0 {
 	    zeros += 1
 	}
	fmt.Println(sum)
    }
    return zeros
}
