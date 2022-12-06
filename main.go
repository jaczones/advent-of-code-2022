package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1(readCSV())
}

func part1(rows [][]string) {
	count := 0
	count2 := 0
	for _, rec := range rows {
		elf1 := strings.Split(string(rec[0]), "-")
		elf2 := strings.Split(string(rec[1]), "-")
		elf1start, _ := strconv.Atoi(elf1[0])
		elf1end, _ := strconv.Atoi(elf1[1])
		elf2start, _ := strconv.Atoi(elf2[0])
		elf2end, _ := strconv.Atoi(elf2[1])
		if elf1start <= elf2start && elf1end >= elf2end {
			count++
		} else if elf2start <= elf1start && elf2end >= elf1end {
			count++
		} else if elf1start == elf2start && elf1end == elf2end {
			count++
		}
		if elf1end >= elf2start && elf2end >= elf1start {
			count2++
		}
	}
	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", count2)
}

func readCSV() [][]string {
	f, err := os.Open("data.csv")
	var rows [][]string
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		rows = append(rows, rec)
	}
	return rows
}
