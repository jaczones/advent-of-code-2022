package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

var values = []string{"0", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func main() {
	part1(readLines("input.txt"))
	part2(readLines("input.txt"))
}

func part2(rucksacks []string) {
	total := 0
	for i := 0; i < len(rucksacks); i++ {
		group := []string{rucksacks[i], rucksacks[i+1], rucksacks[i+2]}
		elf1, elf2, elf3 := group[0], group[1], group[2]
		var priority int
		for _, i := range elf1 {
			for _, j := range elf2 {
				for _, k := range elf3 {
					if i == j && j == k {
						priority = slices.Index(values, string(i))
					}
				}
			}
		}
		total += priority
		i += 2
	}
	fmt.Println(total)
}

func part1(rucksacks []string) {
	total := 0
	for _, rucksack := range rucksacks {
		half := len(rucksack) / 2
		first := rucksack[:half]
		second := rucksack[half:]
		var priority int
		for _, i := range first {
			for _, j := range second {
				if i == j {
					priority = slices.Index(values, string(i))
				}
			}
		}
		total += priority
	}
	fmt.Println(total)
}

func readLines(path string) []string {
	file, err := os.Open(path)
	var lines []string
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
