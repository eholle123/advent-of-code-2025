package main

import ("os"; "io"; "bufio"; "fmt"; "strings"; "strconv")

func main() {
	id_ranges:=ParseIdRanges(os.Stdin)
	// fmt.Println(id_ranges)

	invalid_ids_1:=GetInvalidIds(id_ranges,"Part 1")
	// fmt.Println("Part 1 invalid_ids:",invalid_ids_1)

	sum_invalid_1:=SumInvalidIds(invalid_ids_1)
	fmt.Println("Day 2 Part 1 answer:",sum_invalid_1)

	invalid_ids_2:=GetInvalidIds(id_ranges,"Part 2")
	// fmt.Println("Part 2 invalid_ids:",invalid_ids_2)

	sum_invalid_2:=SumInvalidIds(invalid_ids_2)
	fmt.Println("Day 2 Part 2 answer:",sum_invalid_2)
}
type IDRange struct {
	start	string
	stop 	string
}
func ParseIdRanges(in io.Reader) []IDRange {
	scanner:=bufio.NewScanner(in)
	var id_ranges []IDRange
	for scanner.Scan() {
		line:=scanner.Text()
		before,after,found:=strings.Cut(line,",")
		for found==true {
			start,stop,_:=strings.Cut(before,"-")
			id_ranges = append(id_ranges,IDRange{start,stop})
			before,after,found = strings.Cut(after,",")
		}
		if before!="" {
			start,stop,_:=strings.Cut(before,"-")
			id_ranges = append(id_ranges,IDRange{start,stop})
		}
	}
	er:=scanner.Err(); if er!=nil { panic(er) }
	return id_ranges
}	
func GetInvalidIds(id_ranges []IDRange, p string) []string {
	var invalid_ids []string
	for _,id_range:=range id_ranges {
		start,er:=strconv.Atoi(id_range.start); if er!=nil { panic(er) }
		stop ,er:=strconv.Atoi(id_range.stop ); if er!=nil { panic(er) }
		r:=stop-start+1
		for i:=range r {
			id:=strconv.Itoa(start+i)
			if p=="Part 1" && CheckSymmetryPart1(id) {
				invalid_ids = append(invalid_ids,id)
			}
			if p=="Part 2" && CheckSymmetryPart2(id) {
				invalid_ids = append(invalid_ids,id)
			}
		}
	}
	return invalid_ids
}
func CheckSymmetryPart1(s string) bool {
	if len(s)%2==0 {
		p1:=s[:len(s)/2]
		p2:=s[len(s)/2:]
		if p1==p2 {
			return true
		}
	}
	return false
}
func CheckSymmetryPart2(s string) bool {
	if len(s)>1 {
		for i:=range len(s)+1 {
			if i!=0 {
				if len(s)%i==0 {
					var prefix   string
					var sub_syms []string
					if i==1 {
						prefix = string(s[0])
					} else {
						prefix = string(s[:len(s)/i])
					}
					after,found := strings.CutPrefix(s,prefix)
					sub_syms = append(sub_syms,prefix)
					for range i-1 {
						after,found = strings.CutPrefix(after,prefix)
						if found {
							sub_syms = append(sub_syms,prefix)
						}
					}
					new_s:=""
					for j:=range len(sub_syms) {
						new_s += sub_syms[j]
					}
					if string(new_s)==s {
						return true
					}
				} 
			}
		}
  }
	return false
}
func SumInvalidIds(ids []string) int {
	var sum int
	for _,id:=range ids {
		if len(id)!=1 {
			num,er:=strconv.Atoi(id); if er!=nil { panic(er) }
			sum = sum + num
		}
	}
	return sum
}
