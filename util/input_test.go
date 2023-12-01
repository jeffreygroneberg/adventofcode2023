package util_test

import (
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		filename string
		expected []string
	}{
		{"testdata/test1.txt", []string{"line 1", "line 2", "line 3"}},
		{"testdata/test2.txt", []string{"hello", "world"}},
		{"testdata/test3.txt", []string{"123", "456", "789"}},
	}

	for _, test := range tests {
		result, _ := util.ReadFile(test.filename)
		if len(result) != len(test.expected) {
			t.Errorf("Expected %d lines, but got %d lines for file %s", len(test.expected), len(result), test.filename)
		}
		for i, line := range result {
			if line != test.expected[i] {
				t.Errorf("Expected '%s', but got '%s' for line %d in file %s", test.expected[i], line, i+1, test.filename)
			}
		}
	}
}

func TestReadFileNonExistent(t *testing.T) {
	filename := "testdata/nonexistent.txt"
	result, err := util.ReadFile(filename)

	if err == nil {
		t.Errorf("Expected error, but got %v", result)
	}
}
