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
    fmt.Println("Day 3 Part 1:",sum)

    sum_12:=TurnOnBig12(lines)
    fmt.Println("Day 3 Part 2:",sum_12)
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

type Big struct {
    index	int
    num		int
}

func GetBatteriesOn(s string) int {
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

func TurnOnBig12(lines []string) int {
    energy_sum:=0
    for _,line:=range lines {
	fmt.Println(line)
	var er error
	battery:=0
	if len(line) <  12 {
	    panic("too few batteries in bank!")
	} else if len(line) == 12 {
	    battery,er=strconv.Atoi(line); if er!=nil { panic(er) }
	} else {
	    fmt.Println(line[:12])
	    battery_str:=string(line[:12])
	    battery,er=strconv.Atoi(battery_str); if er!=nil { panic(er) }
	    max_i:=len(line)-12
	    for i:=range max_i {
		battery_i_str:=string(line[i:i+12])
		battery_i,er:=strconv.Atoi(battery_i_str); if er!=nil { panic(er) }
		if battery<battery_i {
		    battery = battery_i
		}
		for j:=range 12 {
		    // fmt.Println(i+j, max_i)
		    if i+j<=12 {
    		        battery_j_str:=string(line[i+j:i+12])
		        curr_num_str:=string(battery_i_str[i:i+j])+battery_j_str
		        curr_num,er:=strconv.Atoi(curr_num_str); if er!=nil { panic(er) }
		        // fmt.Println(curr_num, len(curr_num_str))
		        if battery<curr_num {
		            battery = curr_num 
	    	        }
		    }
		}
	    }
	}
        fmt.Println(battery)
	energy_sum += battery
    }
    return energy_sum
}
