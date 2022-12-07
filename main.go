package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var crate = make(map[int]string)

func main() {
	lines := readLines("input.txt")
	instructions := readLines("instructions.txt")
	makeCrateMap(lines)
	parseStrings(instructions)
	fmt.Println(crate)
}

func parseStrings(instructions []string) {
	for _, line := range instructions {
		str1 := line
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

		submatchall := re.FindAllString(str1, -1)
		var nums []int
		for _, element := range submatchall {
			num, _ := strconv.Atoi(element)
			nums = append(nums, num)
		}
		moveCrate(nums[0], nums[1], nums[2])
	}
}

func moveCrate(amount int, from int, to int) {
	var stack string
	if amount > 1 {
		stack = string(crate[from][len(crate[from])-amount:])
		remaining := string(crate[from][:len(crate[from])-amount])
		crate[from] = remaining
		crate[to] = crate[to] + stack
	} else {
		for i := 1; i <= amount; i++ {
			stack = string(crate[from][len(crate[from])-1])
			crate[from] = crate[from][:len(crate[from])-1]
			crate[to] = crate[to] + stack
		}
	}
}

func makeCrateMap(lines []string) map[int]string {
	count := 1
	for _, line := range lines {
		crate[count] = line
		count++
	}
	return crate
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
