package cards

import (
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestGetCardId(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"card 123", 123},
		{"card 456", 456},
		{"card 789", 789},
		{"card 1", 1},
	}

	for _, test := range tests {
		result := GetCardId(test.input)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d for input %s", test.expected, result, test.input)
		}
	}
}

func TestExtractCardIdWinningNumbersAndDraws(t *testing.T) {
	tests := []struct {
		input    string
		expected Card
	}{

		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Card{CardId: 1, WinningNumbers: []int{41, 48, 83, 86, 17}, Draws: []int{83, 86, 6, 31, 17, 9, 48, 53}}},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Card{CardId: 2, WinningNumbers: []int{13, 32, 20, 16, 61}, Draws: []int{61, 30, 68, 82, 17, 32, 24, 19}}},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Card{CardId: 3, WinningNumbers: []int{1, 21, 53, 59, 44}, Draws: []int{69, 82, 63, 72, 16, 21, 14, 1}}},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Card{CardId: 4, WinningNumbers: []int{41, 92, 73, 84, 69}, Draws: []int{59, 84, 76, 51, 58, 5, 54, 83}}},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Card{CardId: 5, WinningNumbers: []int{87, 83, 26, 28, 32}, Draws: []int{88, 30, 70, 12, 93, 22, 82, 36}}},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Card{CardId: 6, WinningNumbers: []int{31, 18, 13, 56, 72}, Draws: []int{74, 77, 10, 23, 35, 67, 36, 11}}},
	}

	for _, test := range tests {

		result := ExtractCardIdWinningNumbersAndDraws(test.input)

		// compare result, but you need to check the slices manually
		if result.CardId != test.expected.CardId {
			t.Errorf("Expected %d, but got %d for input %s", test.expected.CardId, result.CardId, test.input)
		}

		if len(result.WinningNumbers) != len(test.expected.WinningNumbers) {
			t.Errorf("Expected %d, but got %d for input %s", len(test.expected.WinningNumbers), len(result.WinningNumbers), test.input)
		}

		if len(result.Draws) != len(test.expected.Draws) {
			t.Errorf("Expected %d, but got %d for input %s", len(test.expected.Draws), len(result.Draws), test.input)
		}

		for i, winningNumber := range result.WinningNumbers {
			if winningNumber != test.expected.WinningNumbers[i] {
				t.Errorf("Expected %d, but got %d for input %s", test.expected.WinningNumbers[i], winningNumber, test.input)
			}
		}

		for i, draw := range result.Draws {
			if draw != test.expected.Draws[i] {
				t.Errorf("Expected %d, but got %d for input %s", test.expected.Draws[i], draw, test.input)
			}
		}
	}

}

func TestGetWinsExample1(t *testing.T) {
	tests := []struct {
		card     Card
		expected []int
	}{
		{
			card: Card{
				WinningNumbers: []int{1, 2, 3, 4, 5},
				Draws:          []int{1, 2, 3, 4, 5},
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			card: Card{
				WinningNumbers: []int{1, 2, 3, 4, 5},
				Draws:          []int{6, 7, 8, 9, 10},
			},
			expected: []int{},
		},
		{
			card: Card{
				WinningNumbers: []int{1, 2, 3, 4, 5},
				Draws:          []int{1, 3, 5, 7, 9},
			},
			expected: []int{1, 3, 5},
		},
	}

	for _, test := range tests {
		result, _ := GetHittingNumbersWithPower(test.card)

		if len(result) != len(test.expected) {
			t.Errorf("Expected %v, but got %v for card %+v", test.expected, result, test.card)
		}

		for i, win := range result {
			if win != test.expected[i] {
				t.Errorf("Expected %v, but got %v for card %+v", test.expected, result, test.card)
				break
			}
		}
	}
}

func TestWinsWithExample1(t *testing.T) {

	tests := []struct {
		card     Card
		expected int
	}{
		{
			card: Card{
				CardId:         1,
				WinningNumbers: []int{41, 48, 83, 86, 17},
				Draws:          []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			expected: 8,
		},
		{
			card: Card{
				CardId:         2,
				WinningNumbers: []int{13, 32, 20, 16, 61},
				Draws:          []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			expected: 2,
		},
		{
			card: Card{
				CardId:         3,
				WinningNumbers: []int{1, 21, 53, 59, 44},
				Draws:          []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
			expected: 2,
		},
		{
			card: Card{
				CardId:         4,
				WinningNumbers: []int{41, 92, 73, 84, 69},
				Draws:          []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
			expected: 1,
		},
		{
			card: Card{
				CardId:         5,
				WinningNumbers: []int{87, 83, 26, 28, 32},
				Draws:          []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			expected: 0,
		},
		{
			card: Card{
				CardId:         6,
				WinningNumbers: []int{31, 18, 13, 56, 72},
				Draws:          []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
			expected: 0,
		},
	}

	sumOfPower := 0
	for _, test := range tests {
		_, power := GetHittingNumbersWithPower(test.card)

		if power != test.expected {
			t.Errorf("Expected %d, but got %d for card %+v", test.expected, power, test.card)
		}

		sumOfPower += power
	}

	if sumOfPower != 13 {
		t.Errorf("Expected %d, but got %d", 13, sumOfPower)
	}
}

func TestWinsWithData1(t *testing.T) {

	tests, _ := util.ReadFile("testdata/part1.txt")

	cards := make([]Card, len(tests))

	for i, test := range tests {
		cards[i] = ExtractCardIdWinningNumbersAndDraws(test)
	}

	sumOfPower := 0

	for _, card := range cards {
		_, power := GetHittingNumbersWithPower(card)
		sumOfPower += power
	}

	if sumOfPower != 21821 {
		t.Errorf("Expected %d, but got %d", 13, sumOfPower)
	}
}

func TestAddIdsToSlice(t *testing.T) {

	// create a slice of ints
	slice := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 2, 3, 3, 4, 4, 5}

	// add the ids to the slice
	slice = AddIdsToSlice(1, 3, slice)

	// check if the slice is correct
	if len(slice) != len(expected) {
		t.Errorf("Expected %v, but got %v", expected, slice)
	}

}

func TestGetCardIdsWithWinningCopyCondition(t *testing.T) {

	inputCards := []Card{
		{CardId: 1, WinningNumbers: []int{1, 2, 3}, Draws: []int{1, 5, 7}},
		{CardId: 2, WinningNumbers: []int{4, 5, 6}, Draws: []int{4, 5, 6}},
		{CardId: 3, WinningNumbers: []int{7, 8, 9}, Draws: []int{7, 8, 9}},
	}

	expected := []int{1, 2, 2, 3, 3, 3}

	result := GetCardIdsWithWinningCopyCondition(inputCards)

	if len(result) != len(expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	for i, id := range result {
		if id != expected[i] {
			t.Errorf("Expected %v, but got %v", expected, result)
			break
		}
	}
}

func TestGetCardIdsWithWinningCopyConditionExample2(t *testing.T) {

	inputCards := []Card{

		{CardId: 1, WinningNumbers: []int{41, 48, 83, 86, 17}, Draws: []int{83, 86, 6, 31, 17, 9, 48, 53}},
		{CardId: 2, WinningNumbers: []int{13, 32, 20, 16, 61}, Draws: []int{61, 30, 68, 82, 17, 32, 24, 19}},
		{CardId: 3, WinningNumbers: []int{1, 21, 53, 59, 44}, Draws: []int{69, 82, 63, 72, 16, 21, 14, 1}},
		{CardId: 4, WinningNumbers: []int{41, 92, 73, 84, 69}, Draws: []int{59, 84, 76, 51, 58, 5, 54, 83}},
		{CardId: 5, WinningNumbers: []int{87, 83, 26, 28, 32}, Draws: []int{88, 30, 70, 12, 93, 22, 82, 36}},
		{CardId: 6, WinningNumbers: []int{31, 18, 13, 56, 72}, Draws: []int{74, 77, 10, 23, 35, 67, 36, 11}},
	}

	expected := 30

	result := GetCardIdsWithWinningCopyCondition(inputCards)

	if len(result) != expected {
		t.Errorf("Expected %d, but got %d", expected, len(result))
	}

}

func TestGetCardIdsWithWinningCopyConditionData2(t *testing.T) {

	tests, _ := util.ReadFile("testdata/part2.txt")

	cards := make([]Card, len(tests))

	for i, test := range tests {
		cards[i] = ExtractCardIdWinningNumbersAndDraws(test)
	}
	expected := 30

	result := GetCardIdsWithWinningCopyCondition(cards)

	if len(result) != expected {
		t.Errorf("Expected %d, but got %d", expected, len(result))
	}

}
