package engine

import (
	"reflect"
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestGetPartNumbersFromEngineString(t *testing.T) {

	lines := []string{
		"...123...456",
		"789...",
		"...0...",
	}

	// expected are the part numbers we expect to get from the engine string
	expected := []PartNumber{
		{Value: 123, StartIndex: 3, EndIndex: 5, Row: 0},
		{Value: 456, StartIndex: 9, EndIndex: 11, Row: 0},
		{Value: 789, StartIndex: 0, EndIndex: 2, Row: 1},
		{Value: 0, StartIndex: 3, EndIndex: 3, Row: 2},
	}

	result := GetPartNumbersFromEngineString(lines)

	if len(result) != len(expected) {
		t.Errorf("Unexpected number of part numbers. Expected %d, got %d", len(expected), len(result))
		return
	}

	// check if the part numbers are correct
	for i, partNumber := range result {
		if partNumber != expected[i] {
			t.Errorf("Unexpected part number at index %d. Expected %+v, got %+v", i, expected[i], partNumber)
		}
	}

}
func TestGetAllAdjacentNumbersWithStringLines(t *testing.T) {

	lines := []string{
		"...123...456",
		"789...",
		"...0...",
	}

	// expected are the part numbers we expect to get from the engine string
	partNumbers := GetPartNumbersFromEngineString(lines)

	// Single Number
	result := GetAllAdjacentNumbers(partNumbers, 0, 0)
	expected := []PartNumber{
		{Value: 789, StartIndex: 0, EndIndex: 2, Row: 1},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Expected %+v, got %+v", expected, result)
	}

	// Test case 3: Three numbers
	result = GetAllAdjacentNumbers(partNumbers, 3, 1)
	expected = []PartNumber{
		{Value: 123, StartIndex: 3, EndIndex: 5, Row: 0},
		{Value: 789, StartIndex: 0, EndIndex: 2, Row: 1},
		{Value: 0, StartIndex: 3, EndIndex: 3, Row: 2},
	}

	// Test case 4: No numbers
	result = GetAllAdjacentNumbers(partNumbers, 6, 2)
	// check for empty array
	if len(result) != 0 {
		t.Errorf("Unexpected result. Expected empty array, got %+v", result)
	}

}

func TestGetSymbolsFromEngineString(t *testing.T) {

	// symbols are not a number and NOT "."
	// symbols might be # or * or &

	lines := []string{
		".$.123.#.456",
		"789.*.",
		"...0.#.",
	}

	// expected are the part numbers we expect to get from the engine string
	expected := []Symbol{
		{Value: "$", x: 1, y: 0},
		{Value: "#", x: 7, y: 0},
		{Value: "*", x: 4, y: 1},
		{Value: "#", x: 5, y: 2},
	}

	result := GetSymbolsFromEngineString(lines)

	if len(result) != len(expected) {
		t.Errorf("Unexpected number of symbols. Expected %d, got %d", len(expected), len(result))
		return
	}

	// check if the part numbers are correct
	for i, symbol := range result {
		if symbol != expected[i] {
			t.Errorf("Unexpected symbol at index %d. Expected %+v, got %+v", i, expected[i], symbol)
		}
	}
}

func TestNumbersSumAroundSymbolsWithExample1(t *testing.T) {

	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 4361
	// expected are the part numbers we expect to get from the engine string
	partNumbers := GetPartNumbersFromEngineString(lines)
	symbols := GetSymbolsFromEngineString(lines)

	sum := 0
	for _, symbol := range symbols {

		// get all the numbers around it
		adjacentNumbers := GetAllAdjacentNumbers(partNumbers, symbol.x, symbol.y)

		for _, adjacentNumber := range adjacentNumbers {
			sum += adjacentNumber.Value
		}

	}

	if sum != expected {
		t.Errorf("Unexpected sum. Expected %d, got %d", expected, sum)
	}

}
func TestNumbersSumAroundSymbolsWithData1(t *testing.T) {

	lines, _ := util.ReadFile("testdata/part1.txt")

	expected := 539433
	// expected are the part numbers we expect to get from the engine string
	partNumbers := GetPartNumbersFromEngineString(lines)
	symbols := GetSymbolsFromEngineString(lines)

	sum := 0
	for _, symbol := range symbols {

		// get all the numbers around it
		adjacentNumbers := GetAllAdjacentNumbers(partNumbers, symbol.x, symbol.y)

		for _, adjacentNumber := range adjacentNumbers {
			sum += adjacentNumber.Value
		}

	}

	if sum != expected {
		t.Errorf("Unexpected sum. Expected %d, got %d", expected, sum)
	}
}

func TestGetSymbolsFromEngineStringWith(t *testing.T) {
	lines := []string{
		".$.123.#.456",
		"789.*.",
		"...0.#.",
	}

	symbol := "#"

	expected := []Symbol{
		{Value: "#", x: 7, y: 0},
		{Value: "#", x: 5, y: 2},
	}

	result := GetSymbolsFromEngineStringWith(lines, symbol)

	if len(result) != len(expected) {
		t.Errorf("Unexpected number of symbols. Expected %d, got %d", len(expected), len(result))
		return
	}

	// check if the symbols are correct
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected symbols. Expected %+v, got %+v", expected, result)
	}
}

func TestNumbersSumAndMultiAroundSymbolsWithExample2(t *testing.T) {

	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 467835

	// get all symbols with exactly 2 adjacent numbers
	symbols := GetSymbolsFromEngineStringWithExactAdjacentNumbers(lines, "*", 2)

	// take every symbol and multiply the value of the adjacent numbers

	sum := 0
	for _, symbol := range symbols {

		// get all the numbers around it
		adjacentNumbers := GetAllAdjacentNumbers(GetPartNumbersFromEngineString(lines), symbol.x, symbol.y)

		multi := 1
		// multiply the values of the adjacent numbers
		for _, adjacentNumber := range adjacentNumbers {
			multi *= adjacentNumber.Value
		}

		sum += multi
	}

	if sum != expected {
		t.Errorf("Unexpected sum. Expected %d, got %d", expected, sum)
	}

}

func TestNumbersSumAndMultiAroundSymbolsWithData2(t *testing.T) {

	lines, _ := util.ReadFile("testdata/part2.txt")

	expected := 75847567

	// get all symbols with exactly 2 adjacent numbers
	symbols := GetSymbolsFromEngineStringWithExactAdjacentNumbers(lines, "*", 2)

	// take every symbol and multiply the value of the adjacent numbers

	sum := 0
	for _, symbol := range symbols {

		// get all the numbers around it
		adjacentNumbers := GetAllAdjacentNumbers(GetPartNumbersFromEngineString(lines), symbol.x, symbol.y)

		multi := 1
		// multiply the values of the adjacent numbers
		for _, adjacentNumber := range adjacentNumbers {
			multi *= adjacentNumber.Value
		}

		sum += multi
	}

	if sum != expected {
		t.Errorf("Unexpected sum. Expected %d, got %d", expected, sum)
	}

}
