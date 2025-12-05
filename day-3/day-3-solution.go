package main

import ("os"; "fmt"; "bufio"; "io"; "strconv")

func main() {
    lines,er:=ParseInput(os.Stdin); if er!=nil { panic(er) }
    // fmt.Println(lines)
    sum:=0
    for i:=range len(lines) {
	if lines[i]!="" {
	    sum += GetBatteriesOn(lines[i])
	}
    }
    fmt.Println("Day 1 Part 1:",sum)
}

func ParseInput(in io.Reader) ([]string,error) {
    scanner:=bufio.NewScanner(in)
    var lines 	[]string
    for scanner.Scan() {
	lines = append(lines, scanner.Text())
    }
    er:=scanner.Err()
    if er!=nil { panic(er) }
    return lines,nil
}

// func GetSum(lines []string) int {
//     // var (
//     //     bigs 	[]int
//     //     b1	int
//     //     b2	int
//     // )
//     // f9,l9,f8,l8,f7,l7,f6,l6,f5,l5,f4,l4
//     for _,line:=range lines {
// //	if len(line)<2 { panic("too short") }
// //	f9:=strings.Index(line,"9")
// //	l9:=strings.LastIndex(line,"9)
// //	if f9!=-1 && l9!=-1 && f9!=l9 {}
// 	// b1 := strconv.Atoi(line[0])
// 	// b2 := strconv.Atoi(line[1])
// 	// for i,s:=range line {
// 	//     if i>1 && strconv.Atoi(s)>b1 {
// 	// 	b1 = strconv.Atoi(line[0])
// 	//     	b2 = strconv.Atoi(line[1])
//  	//     }
// 	// }
// 	fmt.Println(FindGreatest(line))
//     }
//     return 0
// }

type Big struct {
    index	int
    num		int
}

func GetBatteriesOn(s string) int {
    // f9:=strings.Index(line,"9")
    // f8:=strings.Index(line,"8")
    // f7:=strings.Index(line,"7")
    // f6:=strings.Index(line,"6")
    // f5:=strings.Index(line,"5")
    // f4:=strings.Index(line,"4")
    // f3:=strings.Index(line,"3")
    // f2:=strings.Index(line,"2")
    // f1:=strings.Index(line,"1")
    // l9:=strings.LastIndex(line,"9")
    // l8:=strings.LastIndex(line,"8")
    // l7:=strings.LastIndex(line,"7")
    // l6:=strings.LastIndex(line,"6")
    // l5:=strings.LastIndex(line,"5")
    // l4:=strings.LastIndex(line,"4")
    // l3:=strings.LastIndex(line,"3")
    // l2:=strings.LastIndex(line,"2")
    // l1:=strings.LastIndex(line,"1")
    // if f9!=-1 && l9!=-1 && f9!=l9 {
    //     return 99
    // }
    // ns:=[]string{"9","8","7","6","5","4","3","2","1"}
    // fmt.Println(ns)
    var b1, b2 Big
    max_i:=len(s)-1
    for i:=range max_i {
	curr_num,er:=strconv.Atoi(string(s[i]));   if er!=nil { panic(er) }
	next_num,er:=strconv.Atoi(string(s[i+1])); if er!=nil { panic(er) }
    	if b1.num<curr_num {
	    b1.num = curr_num; b1.index = i
	    b2.num = next_num; b2.index = i+1
	}
	if b2.num<curr_num && b1.index!=i {
	    b2.num = curr_num; b2.index = i
	}
	if b2.num<next_num && i+1==max_i {
	    b2.num = next_num; b2.index = i+1
	}
    }
    batteries_on,er:=strconv.Atoi(string(s[b1.index])+string(s[b2.index])); if er!=nil { panic(er) }
    // fmt.Println(batteries_on)
    return batteries_on
}
