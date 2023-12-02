package cubegame

import (
	"reflect"
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestGetCubesForSingleDraw(t *testing.T) {
	tests := []struct {
		draws    string
		expected map[string]int
	}{
		{
			draws: "3 red, 5 blue, 2 green",
			expected: map[string]int{
				"red":   3,
				"blue":  5,
				"green": 2,
			},
		},
		{
			draws: "10 yellow, 4 purple, 6 orange",
			expected: map[string]int{
				"yellow": 10,
				"purple": 4,
				"orange": 6,
			},
		},
		{
			draws: "2 pink, 3 brown, 7 gray",
			expected: map[string]int{
				"pink":  2,
				"brown": 3,
				"gray":  7,
			},
		},
	}

	for _, test := range tests {
		result := GetCubesForSingleDraw(test.draws)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, but got %v for draws %s", test.expected, result, test.draws)
		}
	}
}

func TestExtractGameIdAndDraws(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{
			line:     "Game 1: 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue",
			expected: 1,
		},
		{
			line:     "Game 2: 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue",
			expected: 2,
		},
		{
			line:     "Game 3: 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue",
			expected: 3,
		},
	}

	for _, test := range tests {
		result, _ := ExtractGameIdAndDraws(test.line)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d for line %s", test.expected, result, test.line)
		}
	}
}

func TestGetCubesForAllDraws(t *testing.T) {
	tests := []struct {
		draws    string
		expected map[string]int
	}{
		{
			draws: "3 red, 5 blue, 2 green; 4 red, 6 blue, 3 green",
			expected: map[string]int{
				"red":   4,
				"blue":  6,
				"green": 3,
			},
		},
		{
			draws: "10 yellow, 4 purple, 6 orange; 5 yellow, 2 purple, 8 orange",
			expected: map[string]int{
				"yellow": 10,
				"purple": 4,
				"orange": 8,
			},
		},
		{
			draws: "2 pink, 3 brown, 7 gray; 1 pink, 4 brown, 5 gray",
			expected: map[string]int{
				"pink":  2,
				"brown": 4,
				"gray":  7,
			},
		},
	}

	for _, test := range tests {
		result := GetCubesForAllDraws(test.draws)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, but got %v for draws %s", test.expected, result, test.draws)
		}
	}
}

func TestGetCubesForAllGames(t *testing.T) {

	tests := []struct {
		games    []string
		expected map[int]map[string]int
	}{
		{
			games: []string{

				"Game 1: 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue",
				"Game 2: 3 blue, 1 green, 2 red; 2 red, 2 green, 5 blue; 3 green, 10 blue; 8 red, 1 blue; 3 red, 1 green, 5 blue; 1 blue, 5 red, 3 green",
				"Game 3: 4 green, 1 blue; 6 blue, 5 green, 1 red; 11 green, 10 blue",
			},

			expected: map[int]map[string]int{
				1: {
					"red":   1,
					"blue":  15,
					"green": 14,
				},
				2: {
					"red":   8,
					"blue":  10,
					"green": 3,
				},
				3: {
					"red":   1,
					"blue":  10,
					"green": 11,
				},
			},
		},
	}

	for _, test := range tests {
		result := GetCubesForAllGames(test.games)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, but got %v for games %s", test.expected, result, test.games)
		}
	}

}

// test if game is valid
func TestIsGameValid(t *testing.T) {

	tests := []struct {
		gameCubes   map[string]int
		constraints map[string]int
		expected    bool
	}{
		{
			gameCubes: map[string]int{
				"red":   1,
				"blue":  15,
				"green": 14,
			},
			constraints: map[string]int{
				"red":   1,
				"blue":  15,
				"green": 14,
			},
			expected: true,
		},
		{
			gameCubes: map[string]int{
				"red":   1,
				"blue":  15,
				"green": 14,
			},
			constraints: map[string]int{
				"red":   1,
				"blue":  15,
				"green": 15,
			},
			expected: true,
		},
		{
			gameCubes: map[string]int{
				"red":   1,
				"blue":  15,
				"green": 14,
			},
			constraints: map[string]int{
				"red":   1,
				"blue":  15,
				"green": 13,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		result := isGameValid(test.gameCubes, test.constraints)
		if result != test.expected {
			t.Errorf("Expected %v, but got %v for gameCubes %v and constraints %v", test.expected, result, test.gameCubes, test.constraints)
		}
	}

}

func TestGetInvalidGames(t *testing.T) {

	playedGames := []string{
		"Game 1: 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue",
		"Game 2: 3 blue, 1 green, 2 red; 2 red, 2 green, 5 blue; 3 green, 10 blue; 8 red, 1 blue; 3 red, 1 green, 5 blue; 1 blue, 5 red, 3 green",
		"Game 3: 4 green, 1 blue; 6 blue, 5 green, 1 red; 11 green, 10 blue",
	}

	constraints := map[string]int{
		"red":   5,
		"blue":  10,
		"green": 15,
	}

	expected := map[int]map[string]int{
		1: {
			"red":   1,
			"blue":  15,
			"green": 14,
		},
		2: {
			"red":   8,
			"blue":  10,
			"green": 3,
		},
	}

	result := GetGamesWithConstraints(playedGames, constraints, false)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGetValidGamesWithExample1(t *testing.T) {

	playedGames := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	constraints := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	expected := 8

	// get all valid games
	validGames := GetGamesWithConstraints(playedGames, constraints, true)

	// get the sum of the ids
	sum := 0
	for gameId, _ := range validGames {
		sum += gameId
	}

	if sum != expected {
		t.Errorf("Expected %d, but got %d", expected, sum)
	}

}

func TestGetValidGamesWithTestData1(t *testing.T) {

	playedGames, _ := util.ReadFile("testdata/part1.txt")

	constraints := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	expected := 2164
	validGames := GetGamesWithConstraints(playedGames, constraints, true)

	sum := 0
	for gameId, _ := range validGames {
		sum += gameId
	}

	if sum != expected {
		t.Errorf("Expected %d, but got %d", expected, sum)
	}

}
