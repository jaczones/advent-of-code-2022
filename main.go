package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	getHighestCalories(readLines("input1.txt"))
}

func getHighestCalories(countPerElf []int) {
	var highest int
	var currentCount int
	var totals []int
	for _, item := range countPerElf {
		if item != 0 {
			currentCount = currentCount + item
			if currentCount > highest {
				highest = currentCount
			}
		} else {
			totals = append(totals, currentCount)
			currentCount = 0
		}
	}
	sort.Ints(totals)
	fmt.Printf("Solution 1: %v \n", totals[len(totals)-1])
	fmt.Printf("Solution 2: %v \n", totals[len(totals)-1]+totals[len(totals)-2]+totals[len(totals)-3])
}

func readLines(path string) []int {
	file, err := os.Open(path)
	var lines []int
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, i)
	}
	return lines
}
