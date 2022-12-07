package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	getScorePart1(readFromCSV())
	getScorePart2(readFromCSV())
}

func getScorePart1(rounds [][]string) {
	total := 0
	for _, row := range rounds {
		elfThrow := row[0]
		myThrow := row[1]
		if myThrow == "X" {
			total += 1
			if elfThrow == "A" {
				total += 3
			} else if elfThrow == "C" {
				total += 6
			}
		} else if myThrow == "Y" {
			total += 2
			if elfThrow == "A" {
				total += 6
			} else if elfThrow == "B" {
				total += 3
			}
		} else {
			total += 3
			if elfThrow == "B" {
				total += 6
			} else if elfThrow == "C" {
				total += 3
			}
		}
	}
	fmt.Printf("Part 1 score: %v \n", total)
}

func getScorePart2(rounds [][]string) {
	total := 0
	for _, row := range rounds {
		elfThrow := row[0]
		result := row[1]
		if result == "X" {
			if elfThrow == "A" {
				total += 3
			} else if elfThrow == "B" {
				total += 1
			} else if elfThrow == "C" {
				total += 2
			}
		} else if result == "Y" {
			total += 3
			if elfThrow == "A" {
				total += 1
			} else if elfThrow == "B" {
				total += 2
			} else if elfThrow == "C" {
				total += 3
			}
		} else {
			total += 6
			if elfThrow == "A" {
				total += 2
			} else if elfThrow == "B" {
				total += 3
			} else if elfThrow == "C" {
				total += 1
			}
		}
	}
	fmt.Printf("Part 2 score: %v \n", total)
}

func readFromCSV() [][]string {
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
