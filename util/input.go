package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// reads a file and returns a slice of strings
func ReadFile(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("failed to open file: %s", err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	// read file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("failed to read file: %s", err)
		return nil, err
	}

	// return slice of strings
	return lines, nil
}

func IsNumber(char byte) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}

func IsNumberOrDot(char byte) bool {
	return IsNumber(char) || char == '.'
}
