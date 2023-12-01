package calibration_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/day1/calibration"
	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestCalculateCalibrationForLineUsingExample(t *testing.T) {

	tests := []struct {
		line     string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		result, err := calibration.CalculateCalibrationForLine(test.line, false)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != test.expected {
			t.Errorf("Expected %d, but got %d for line %s", test.expected, result, test.line)
		}
	}
}

func TestCalculateCalibrationSumUsingExample1(t *testing.T) {

	var lines = []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	var result int
	for _, line := range lines {
		value, err := calibration.CalculateCalibrationForLine(line, false)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		fmt.Printf("Adding %d to %d\n", value, result)
		result += value
	}

	// check if the result is 142
	if result != 142 {
		t.Errorf("Expected %d, but got %d", 142, result)
	}
}

func TestCalculateCalibrationForLineUsingExample2(t *testing.T) {

	var lines = []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}

	var result int
	for _, line := range lines {
		value, err := calibration.CalculateCalibrationForLine(line, true)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		fmt.Printf("Adding %d to %d\n", value, result)
		result += value
	}

	// check if the result is 281
	if result != 281 {
		t.Errorf("Expected %d, but got %d", 281, result)
	}

}

// create a test that is using the input file part1.txt and checks the result against the expected result
func TestCalculateCalibrationSumForFile_Part1(t *testing.T) {

	// read the file
	lines, _ := util.ReadFile("testdata/part1.txt")

	// loop through each line and cummulate the returned integer values
	var result int
	for _, line := range lines {
		value, err := calibration.CalculateCalibrationForLine(line, false)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		result += value
	}

	// check if the result is 54573
	if result != 54573 {
		t.Errorf("Expected %d, but got %d", 54573, result)
	}

}

// create a test that is using the input file part1.txt and checks the result against the expected result
func TestCalculateCalibrationSumForFile_Part2(t *testing.T) {

	// read the file
	lines, _ := util.ReadFile("testdata/part2.txt")

	// loop through each line and cummulate the returned integer values
	var result int
	for _, line := range lines {
		value, err := calibration.CalculateCalibrationForLine(line, true)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		log.Println(line)
		log.Printf("Adding %d to %d\n", value, result)
		result += value
	}

	// check if the result is 53868
	if result != 54591 {
		t.Errorf("Expected %d, but got %d", 54591, result)
	}

}

func TestIsNumber(t *testing.T) {

	tests := []struct {
		char     byte
		expected bool
	}{
		{'1', true},
		{'a', false},
		{'b', false}}

	for _, test := range tests {
		result := calibration.IsNumber(test.char)
		if result != test.expected {
			t.Errorf("Expected %t, but got %t for char %s", test.expected, result, string(test.char))
		}
	}
}

func TestReplacer(t *testing.T) {

	tests := []struct {
		line     string
		expected string
	}{
		{"one123", "1123"},
		{"twothree", "23"},
		{"oneeighthree", "183"},
		{"three4five", "345"},
		{"12xyfour", "12xy4"},
		{"er1five", "er15"},
		{"sixseven", "67"},
		{"seven8nine", "789"},
		{"eightnine", "89"},
		{"nine", "9"},
		{"ten", "ten"},
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"1111111122four", "11111111224"},
	}

	for _, test := range tests {
		result := calibration.Replacer.Replace(test.line)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s for line %s", test.expected, result, test.line)
		}
	}

}

func TestCalculateCalibrationWithReplacerUsingExample(t *testing.T) {

	tests := []struct {
		line     string
		expected int
	}{
		{"one123", 13},
		{"twothree", 23},
		{"three4five", 35},
		{"12xyfour", 14},
		{"er1five", 15},
		{"sixseven", 67},
		{"seven8nine", 79},
		{"eightnine", 89},
		{"nine", 99},
		{"ten1", 11},
		{"one", 11},
		{"two", 22},
		{"three", 33},
		{"four", 44},
		{"1111111122four", 14},
		{"nineererererererer", 99},
	}

	for _, test := range tests {
		result, err := calibration.CalculateCalibrationForLine(test.line, true)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != test.expected {
			t.Errorf("Expected %d, but got %d for line %s", test.expected, result, test.line)
		}
	}
}
