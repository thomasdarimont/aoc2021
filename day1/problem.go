package day1

import (
	"aoc2021/aoc"
	"bufio"
	"log"
	"strconv"
)

func Solve() {
	aoc.ProcessInput("day1/input.txt", "1", part1)
	aoc.ProcessInput("day1/input.txt", "2", part2)
}

func part1(s *bufio.Scanner) interface{} {

	incs := 0
	prev := -1

	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		if prev != -1 && prev < val {
			incs++
		}
		prev = val
	}

	return incs
}

func part2(s *bufio.Scanner) interface{} {

	var items []int

	for s.Scan() {
		val, _ := strconv.Atoi(s.Text())
		items = append(items, val)
	}

	incs := 0
	wndSize := 3

	for i := 0; i < len(items)-wndSize; i++ {
		if items[i+wndSize] > items[i] {
			incs++
		}
	}

	return incs
}
