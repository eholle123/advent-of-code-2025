package main

import ("os"; "io"; "bufio"; "fmt"; "strings"; "strconv")

func main() {
	id_ranges:=ParseIdRanges(os.Stdin)
	// fmt.Println(id_ranges)

	invalid_ids:=GetInvalidIds(id_ranges)
	// fmt.Println("invalid_ids:",invalid_ids)

	sum_invalid:=SumInvalidIds(invalid_ids)
	fmt.Println("Day 2 Part 1 answer:",sum_invalid)
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
	}
	er:=scanner.Err(); if er!=nil { panic(er) }
	return id_ranges
}	
func GetInvalidIds(id_ranges []IDRange) []string {
	var invalid_ids []string
	var c1					string
	for _,id_range:=range id_ranges {
		start,er:=strconv.Atoi(id_range.start); if er!=nil { panic(er) }
		stop ,er:=strconv.Atoi(id_range.stop ); if er!=nil { panic(er) }
		r:=stop-start+1
		for i:=range r {
			id:=strconv.Itoa(start+i)
			for range len(id)+1 {
				c1 += string(id[0])
			}
			if len(id)%2==0 {
				p1:=id[:len(id)/2]
				p2:=id[len(id)/2:]
				// fmt.Println("p1:",p1,"p2:",p2)
				if p1==p2 {
					invalid_ids = append(invalid_ids,id)
				}
			}
		}
	}
	return invalid_ids
}
func CheckSymmetry(s string, n int) bool {
	id,er:=strconv.Atoi(s); if er!=nil { panic(er) }
	for i:=range len(s)+1 {
		if len(s)%i==0 && i!=0 && i!=1 && len(s)/i!=1 && n==i {
			var prefix string
			prefix:=string(s[:len(s)/2])
			current:=string(s[:len(s)/2])
			found:=true
			for found==true {
				after,found:=strings.CutPrefix(s,prefix)

			}
		} 
	}
}
func SumInvalidIds(ids []string) int {
	var sum int
	for _,id:=range ids {
		num,er:=strconv.Atoi(id); if er!=nil { panic(er) }
		sum = sum + num
	}
	return sum
}
