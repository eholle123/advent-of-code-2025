package main

import ("os"; "io"; "fmt"; "strconv"; "bufio"; "strings")

func main() {
	hw:=ParseHomework(os.Stdin)
	fmt.Println(hw)

	sum:=DoHomework(hw)
	fmt.Println("Day 6 Part 1 answer:",sum)
}
type Homework struct {
	nums			[][]int
	operands 	[]string
}
func ParseHomework(in io.Reader) Homework {
	scanner:=bufio.NewScanner(in)
	var nums 			[][]int
	var operands 	[]string
	for scanner.Scan(){
		line:=string(scanner.Text())
		datas:=strings.Fields(line)
		if datas[0]=="+" || datas[0]=="*"{
			for _,data:=range datas{
				operands = append(operands, data)
			}
		} else if datas[0]!="" {
			var row_nums []int
			for _,data:=range datas {
				n,er:=strconv.Atoi(data); if er!=nil { panic(er) }
				row_nums = append(row_nums,n)
			}
			nums = append(nums,row_nums)
		} else {
			continue
		}
	}
	er:=scanner.Err(); if er!=nil { panic(er) }
	return Homework{nums,operands}
}
func DoHomework(hw Homework) int {
	var sum				int
	var solutions []int
	for j:=range len(hw.nums[0]) {
		operator:=hw.operands[j]
		var solution	int
		if operator=="*" { solution = 1 }
		for i:=range len(hw.nums) {
			if operator=="+" {
				solution = solution + hw.nums[i][j]
			}
			if operator=="*" {
				solution = solution * hw.nums[i][j]
			}
		}
		solutions = append(solutions, solution)
		sum += solution
	}
	fmt.Println(solutions)
	return sum
}
