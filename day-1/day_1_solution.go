package main

import ("os"; "fmt"; "bufio"; "io"; "strconv"; "math")

func main() {
    lines,er:=ParseInput(os.Stdin); if er!=nil { panic(er) }
    // fmt.Println(lines)
    rotations:=CalcLockRotations(lines)
    // fmt.Println(rotations)
    combo1:=CalcLockCombo(rotations)
    fmt.Println("Day 1, Part 1 answer: ",combo1)
    combo2:=CalcLockComboAnyClicks(rotations)
    fmt.Println("Day 1, Part 2 answer: ",combo2)
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
    // fmt.Println("start", sum)
    for _,num:=range rotations {
	sum = (sum + num) % 100
 	if sum==0 {
 	    zeros += 1
 	}
	// fmt.Println(sum)
    }
    return zeros
}

func CalcLockComboAnyClicks(rotations []int) int {
    var (
	sum	   int
	zeros	   int
	clicks	   int
	pos	   int
	equo	   int
	remainder  int
    )
    sum	= 50
    // fmt.Println("Input:\tTotal zeros:\tRaw clicks:\tClicks:\tPosition:")
    // fmt.Println("num:\tsum:\tzeros:\tfcs:\trd:\tclicks:\tposition:")
    fmt.Println("num:\tsum:\tremainder:\tzeros:\tclicks:\tposition:")
    for _,num:=range rotations {
	// sum    	= sum + num
        // fcs,rd  :=math.Modf(float64(sum / 100))
	// clicks 	= int(math.Abs(fcs))
	// sum    	= sum % 100
	// pos 	= (sum + 100) % 100
	// zeros += clicks
	// fmt.Println(num,"\t",sum,"\t",zeros,"\t",fcs,"\t",rd,"\t",clicks,"\t",pos)
	equo 	  = int(math.Floor(float64((sum+num)/100)))
	remainder = (sum + int(math.Abs(float64(num)))) % 100
	// clicks	  = (sum - int(math.Abs(float64(remainder)))) / 99
	sum	  = (sum + num) % 100
	pos 	  = (sum + 100) % 100
	clicks 	  = (pos + remainder) / 100
	if remainder==0 && equo==0 { zeros += 1 }
	zeros 	  += clicks
	fmt.Println(num,"\t",sum,"\t",remainder,"\t\t",zeros,"\t",clicks,"\t",pos)
    }
    return zeros
}
