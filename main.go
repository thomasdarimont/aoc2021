package main

import (
	"aoc2021/aoc"
	"aoc2021/day1"
	"aoc2021/day2"
	"aoc2021/day3"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Advent of Code 2021")

	var problem string
	if len(os.Args) > 1 {
		problem = os.Args[1]
	}

	problems := []*aoc.Problem{
		{Name: "day1", Solve: day1.Solve},
		{Name: "day2", Solve: day2.Solve},
		{Name: "day3", Solve: day3.Solve},
	}

	for _, p := range problems {
		if problem != "" && problem != p.Name {
			continue
		}
		p.Solution()
	}
}
