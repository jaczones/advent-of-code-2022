package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")
	getFolderSize(lines)
}

func getFolderSize(lines []string) {
	dirs := make(map[string]Directory)
	pwd := "/"
	dirs[pwd] = Directory{Name: "/", Parent: ""}
	var newDirName string
	var fileSize int
	total := 0

	for _, line := range lines {
		toParse := strings.Split(line, " ")
		if toParse[0] == "$" && toParse[1] == "cd" {
			if toParse[2] == "/" {
				continue
			}
			if toParse[2] == ".." {
				pwd = dirs[pwd].Parent
				continue
			}
			if pwd == "/" {
				newDirName = pwd + toParse[2]
			} else {
				newDirName = pwd + "/" + toParse[2]
			}
			if _, ok := dirs[toParse[2]]; !ok {
				newDir := Directory{Name: toParse[2], Parent: pwd}
				dirs[newDirName] = newDir
			}
			pwd = newDirName
			continue
		}
		fileSize, _ = strconv.Atoi(toParse[0])
		currentDir := dirs[pwd]
		currentDir.Size += fileSize
		dirs[pwd] = currentDir
		parentDir := currentDir.Parent
		for len(parentDir) > 0 {
			folder := dirs[parentDir]
			folder.Size += fileSize
			dirs[parentDir] = folder
			parentDir = folder.Parent
		}
	}
	sizes := []int{}
	for _, dir := range dirs {
		if dir.Size <= 100000 {
			total += dir.Size
		}
		sizes = append(sizes, dir.Size)
	}
	fmt.Println("Part 1: ", total)
	sort.Ints(sizes)
	unused := 70000000 - sizes[len(sizes)-1]
	needed := 30000000 - unused
	dirsBigEnough := []int{}
	for _, size := range sizes {
		if size > needed {
			dirsBigEnough = append(dirsBigEnough, size)
		}
	}
	fmt.Println("Part 2: ", dirsBigEnough[0])
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

type Directory struct {
	Name   string
	Parent string
	Size   int
}
