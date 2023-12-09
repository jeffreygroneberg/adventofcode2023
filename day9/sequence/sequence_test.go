package sequence

import (
	"strconv"
	"strings"
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestCalculateNextInSequence(t *testing.T) {

	// Arrange
	sequence := []int{2, 5, 8, 11, 14}
	expected := 17

	// Act
	actual := CalculateNextInSequence(sequence)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestCalculateNextIntSequenceExample(t *testing.T) {

	// 0 3 6 9 12 15 expected is 18
	// 1 3 6 10 15 21 expected is 28
	// 10 13 16 21 30 45 expected is 68

	tests := []struct {
		sequence []int
		expected int
	}{
		{[]int{0, 3, 6, 9, 12, 15}, 18},
		{[]int{1, 3, 6, 10, 15, 21}, 28},
		{[]int{10, 13, 16, 21, 30, 45}, 68},
	}

	sum := 0
	for _, test := range tests {
		actual := CalculateNextInSequence(test.sequence)
		if actual != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, actual)
		}

		sum += actual
	}

	if sum != 114 {
		t.Errorf("Expected %d, but got %d", 114, sum)
	}
}
func TestCalculateNextIntSequenceData1(t *testing.T) {

	lines, _ := util.ReadFile("testdata/part1.txt")

	sum := 0
	for _, line := range lines {

		// convert string to int []
		// line looks like this 1 8 27 58 122 282 681 1611 3653 7980 17022 35902 75440 158200 330196 682764 1392266 2790739 5490489 10604640 20143903
		// we need to convert it to [1,8,27,58,122,282,681,1611,3653,7980,17022,35902,75440,158200,330196,682764,1392266,2790739,5490489,10604640,20143903]

		s := strings.Split(line, " ")

		sequence := make([]int, len(s))
		for i, v := range s {

			v = strings.TrimSpace(v)
			sequence[i], _ = strconv.Atoi(v)
		}

		actual := CalculateNextInSequence(sequence)
		sum += actual

	}

	if sum != 1681758908 {
		t.Errorf("Expected %d, but got %d", 114, sum)
	}

}

func TestCalculateNextIntSequenceData2(t *testing.T) {

	lines, _ := util.ReadFile("testdata/part1.txt")

	sum := 0
	for _, line := range lines {

		s := strings.Split(line, " ")

		sequence := make([]int, len(s))

		// reverse the input
		for i := len(s) - 1; i >= 0; i-- {

			v := s[i]

			v = strings.TrimSpace(v)
			sequence[len(s)-i-1], _ = strconv.Atoi(v)
		}

		actual := CalculateNextInSequence(sequence)
		sum += actual

	}

	if sum != 803 {
		t.Errorf("Expected %d, but got %d", 114, sum)
	}

}
