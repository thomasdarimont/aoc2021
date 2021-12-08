package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	Name  string
	Solve Solver
}

type Solver func()

func (p *Problem) Solution() {
	fmt.Printf("### Solving Problem '%s'\n", p.Name)
	p.Solve()
	fmt.Println()
}

func ProcessInput(filename string, part string, fun func(s *bufio.Scanner) interface{}) {

	fmt.Printf("# Part %s\n", part)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)

	result := fun(s)
	fmt.Printf("# Answer '%v'\n", result)
}
