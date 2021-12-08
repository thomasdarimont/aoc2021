package day2

import (
	"aoc2021/aoc"
	"bufio"
	"log"
	"strconv"
	"strings"
)

func Solve() {
	aoc.ProcessInput("day2/input.txt", "1", part1)
	aoc.ProcessInput("day2/input.txt", "2", part2)
}

func part1(s *bufio.Scanner) interface{} {

	pos := 0
	depth := 0

	for s.Scan() {

		tokens := strings.Split(s.Text(), " ")

		dir := tokens[0]
		steps, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case "forward":
			pos += steps
		case "up":
			depth -= steps
		case "down":
			depth += steps
		}
	}

	return pos * depth
}

func part2(s *bufio.Scanner) interface{} {

	aim := 0
	pos := 0
	depth := 0

	for s.Scan() {

		tokens := strings.Split(s.Text(), " ")

		dir := tokens[0]
		steps, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal(err)
		}

		switch dir {
		case "forward":
			pos += steps
			depth += aim * steps
		case "up":
			aim -= steps
		case "down":
			aim += steps
		}
	}

	return pos * depth
}
