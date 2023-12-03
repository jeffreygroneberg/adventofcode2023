package engine

import (
	"strconv"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

type PartNumber struct {
	Value      int
	StartIndex int
	EndIndex   int
	Row        int
}

type Symbol struct {
	Value string
	x, y  int
}

func GetPartNumbersFromEngineString(lines []string) []PartNumber {

	var partNumbers []PartNumber

	for row, line := range lines {

		currentNumber := ""

		for index := 0; index < len(line); index++ {

			if util.IsNumber(line[index]) {
				currentNumber += string(line[index])

				// if next char is not a number, we have the full number and can write it to the array
				if index+1 == len(line) || !util.IsNumber(line[index+1]) {

					value, _ := strconv.Atoi(currentNumber)

					partNumber := PartNumber{
						Value:      value,
						StartIndex: index - len(currentNumber) + 1,
						EndIndex:   index,
						Row:        row,
					}

					partNumbers = append(partNumbers, partNumber)
					currentNumber = ""

				}
			}

		}

	}

	return partNumbers

}

func GetAllAdjacentNumbers(partNumbers []PartNumber, x int, y int) []PartNumber {

	var adjacentNumbers []PartNumber

	for _, partNumber := range partNumbers {

		// same row
		if partNumber.Row == y {

			// to the left
			if x-1 == partNumber.EndIndex {
				adjacentNumbers = append(adjacentNumbers, partNumber)
			}

			// to the right
			if x+1 == partNumber.StartIndex {
				adjacentNumbers = append(adjacentNumbers, partNumber)
			}
		}

		// a row above // bellow
		if partNumber.Row == y-1 || partNumber.Row == y+1 {

			// left or right diagonal
			if x-1 == partNumber.EndIndex || x+1 == partNumber.StartIndex {
				adjacentNumbers = append(adjacentNumbers, partNumber)
			}

			// direct above
			if x >= partNumber.StartIndex && x <= partNumber.EndIndex {
				adjacentNumbers = append(adjacentNumbers, partNumber)
			}

		}

	}

	return adjacentNumbers

}

func GetSymbolsFromEngineString(lines []string) []Symbol {

	var symbols []Symbol

	for row, line := range lines {

		for index := 0; index < len(line); index++ {

			// util should not be a number and not "."
			if !util.IsNumber(line[index]) && line[index] != '.' {

				symbol := Symbol{
					Value: string(line[index]),
					x:     index,
					y:     row,
				}

				symbols = append(symbols, symbol)

			}

		}

	}

	return symbols

}

func GetSymbolsFromEngineStringWith(lines []string, symbol string) []Symbol {

	var symbols []Symbol

	for row, line := range lines {

		for index := 0; index < len(line); index++ {

			// util should not be a number and not "."
			if string(line[index]) == symbol {

				symbol := Symbol{
					Value: string(line[index]),
					x:     index,
					y:     row,
				}

				symbols = append(symbols, symbol)

			}

		}

	}

	return symbols

}

func GetSymbolsFromEngineStringWithExactAdjacentNumbers(lines []string, symbol string, maxAdjacentNumbers int) []Symbol {

	var symbols []Symbol

	foundSymbols := GetSymbolsFromEngineStringWith(lines, symbol)

	for _, symbol := range foundSymbols {

		adjacentNumbers := GetAllAdjacentNumbers(GetPartNumbersFromEngineString(lines), symbol.x, symbol.y)

		if len(adjacentNumbers) == maxAdjacentNumbers {
			symbols = append(symbols, symbol)
		}

	}

	return symbols

}
