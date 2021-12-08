package day3

import (
	"aoc2021/aoc"
	"bufio"
	"strconv"
)

func Solve() {
	//aoc.ProcessInput("day3/sample.txt", "sample", part1)
	aoc.ProcessInput("day3/input.txt", "1", part1)

	//aoc.ProcessInput("day3/sample.txt", "sample2", part2)
	aoc.ProcessInput("day3/input.txt", "2", part2)
}

func part1(s *bufio.Scanner) interface{} {

	var gamma int64
	var epsilon int64
	var counts []int64

	for s.Scan() {
		bitstring := s.Text()
		if counts == nil {
			counts = make([]int64, len(bitstring))
		}

		for pos := 0; pos < len(counts); pos++ {
			if bitstring[pos] == '1' {
				counts[pos] += 1
			} else {
				counts[pos] -= 1
			}
		}
	}

	for pos := 0; pos < len(counts); pos++ {
		if counts[pos] > 0 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
		if pos+1 < len(counts) {
			gamma <<= 1
			epsilon <<= 1
		}
	}

	//fmt.Println(strconv.FormatInt(gamma, 2))
	//fmt.Println(strconv.FormatInt(epsilon, 2))

	return gamma * epsilon
}

func part2(s *bufio.Scanner) interface{} {

	var numbers []string

	for s.Scan() {
		numbers = append(numbers, s.Text())
	}

	oxygenNumbers := filterNumbers(0, numbers, 1)
	//fmt.Println(oxygenNumbers)
	oxygen, _ := strconv.ParseInt(oxygenNumbers[0], 2, 64)

	co2Numbers := filterNumbers(0, numbers, -1)
	//fmt.Println(co2Numbers)
	co2, _ := strconv.ParseInt(co2Numbers[0], 2, 64)

	return oxygen * co2
}

func filterNumbers(bitpos int64, numbers []string, inc int64) []string {

	if len(numbers) == 1 {
		return numbers
	}

	counts := make([]int64, len(numbers[0]))

	for _, bitstring := range numbers {
		if bitstring[bitpos] == '1' {
			counts[bitpos] += inc
		} else {
			counts[bitpos] -= inc
		}
	}

	needle := needleFor(bitpos, counts, inc)

	var filtered []string
	for _, bitstring := range numbers {
		if bitstring[bitpos] == needle {
			filtered = append(filtered, bitstring)
		}
	}

	return filterNumbers(bitpos+1, filtered, inc)
}

func needleFor(bitpos int64, counts []int64, inc int64) uint8 {

	if counts[bitpos] > 0 {
		return '1'
	}

	if counts[bitpos] < 0 {
		return '0'
	}

	if inc > 0 {
		return '1'
	}

	return '0'
}
