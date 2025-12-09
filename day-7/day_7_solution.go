package main

import ("os"; "io"; "bufio"; "fmt"; "strings")

func main() {
	lines:=ParseInput(os.Stdin)
	beams:=Cheat(lines)
	fmt.Println("Day 7 Part 1 answer:",beams)
}
func ParseInput(in io.Reader) []string {
	scanner:=bufio.NewScanner(in)
	var lines []string
	for scanner.Scan(){
		line:=string(scanner.Text())
		fmt.Println(line)
		lines = append(lines, line)
	}
	er:=scanner.Err(); if er!=nil { panic(er) }
	return lines
}
func Cheat(lines []string) int{
	rows:=len(lines)
	cols:=len(lines[0])
	last_carrot_line:=lines[len(lines)-2]
	var min_beams, max_carrots int
	// Assuming cols is always odd for symmetry purposes...
	min_beams   = (cols-1)/2 + 1
	max_carrots = (cols-1)/2
	if string(last_carrot_line[0])=="^" {
		min_beams     = (cols-1)/2
		max_carrots   = (cols-1)/2 + 1
	}
	last_beams:=strings.Count(last_carrot_line,".")
	last_carrots:=strings.Count(last_carrot_line,"^")
	fmt.Println("rows:",rows,"cols:",cols)
	fmt.Println("min_beams:",min_beams,"max_carrots",max_carrots)
	fmt.Println(last_carrot_line)
	fmt.Println("last_beams:",last_beams,"last_carrots:",last_carrots)

	// for _,line:=range lines {
	// 	carrots:=strings.Count(line,"^")
	// 	if carrots!=0 { fmt.Println(carrots) }
	// }
	splits:=last_beams*2 + (last_beams-max_carrots)/2 + (last_beams-max_carrots)
	splits2:=min_beams*2 + (last_beams-min_beams)*4
	splits3:=min_beams*2 + (last_beams-min_beams)*(last_beams-last_carrots) + (last_beams-max_carrots)/2
	splits4:=0
	for i:=range rows {
		splits4 = splits4+i+1
	}
	fmt.Println("splits:",splits)
	fmt.Println("splits2:",splits2)
	fmt.Println("splits3:",splits3)
	fmt.Println("splits4:",splits4)
	var just_carrot_lines []string
	for _,line:=range lines {
		if strings.Contains(line,"^") {
			just_carrot_lines = append(just_carrot_lines,line)
		}
	}
	// fmt.Println(just_carrot_lines)
	splits5:=0
	for i:=range len(just_carrot_lines) {
		splits5 = splits5+i+1
	}
	fmt.Println("splits5:",splits5)
	return last_carrots
}
