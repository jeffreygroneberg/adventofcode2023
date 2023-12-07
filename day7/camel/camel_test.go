package camel

import (
	"log"
	"sort"
	"strconv"
	"testing"

	"github.com/jeffreygroneberg/adventofcode2023/util"
)

func TestDetermineTypeAndPower(t *testing.T) {

	tests := []struct {
		input                 string
		expectedType          string
		expectedPower         int
		expectedRelevantCards string
	}{
		{input: "22222", expectedType: "Five of a kind", expectedPower: 1, expectedRelevantCards: "22222"},
		{input: "11A11", expectedType: "Four of a kind", expectedPower: 2, expectedRelevantCards: "1111"},
		{input: "22211", expectedType: "Full house", expectedPower: 3, expectedRelevantCards: "11222"},
		{input: "AAA31", expectedType: "Three of a kind", expectedPower: 4, expectedRelevantCards: "AAA"},
		{input: "1122A", expectedType: "Two pair", expectedPower: 5, expectedRelevantCards: "1122"},
		{input: "11BCD", expectedType: "One pair", expectedPower: 6, expectedRelevantCards: "11"},
		{input: "5A321", expectedType: "Highest card", expectedPower: 7, expectedRelevantCards: "5"},
		{input: "A5432", expectedType: "Highest card", expectedPower: 7, expectedRelevantCards: "A"},
		{input: "5432A", expectedType: "Highest card", expectedPower: 7, expectedRelevantCards: "A"},
		{input: "ACA2A", expectedType: "Three of a kind", expectedPower: 4},
	}

	for _, test := range tests {

		rating := determineTypeAndPower(test.input, orderOfCardsNoJoker)

		if rating.CardType != test.expectedType || rating.Power != test.expectedPower {
			t.Errorf("Test Case: %s\nExpected: %s, %d\nGot: %s, %d", test.input, test.expectedType, test.expectedPower, rating.CardType, rating.Power)
		}

	}

}

func TestCompare(t *testing.T) {

	// So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger.
	// 33332 > 2AAAA

	tests := []struct {
		name    string
		ratings []Rating
		loser   int // 1 if Rating1 > Rating2, -1 if Rating1 < Rating2, 0 if Rating1 == Rating2
	}{

		{ratings: []Rating{
			{CardType: "Five of a kind", Power: 1, RelevantCards: "22222", Source: "22222"},
			{CardType: "Five of a kind", Power: 1, RelevantCards: "11111", Source: "11111"},
		}, loser: 1, name: "Five of a kind"},

		{ratings: []Rating{
			{CardType: "Four of a kind", Power: 2, RelevantCards: "1111", Source: "1111A"},
			{CardType: "Four of a kind", Power: 2, RelevantCards: "2222", Source: "2222B"},
		}, loser: 0, name: "Four of a kind"},

		{ratings: []Rating{
			{CardType: "Full house", Power: 3, RelevantCards: "11222", Source: "22211"},
			{CardType: "Full house", Power: 3, RelevantCards: "11122", Source: "11122"},
		}, loser: 1, name: "Full house"},

		{ratings: []Rating{
			{CardType: "Four of a kind", Power: 2, RelevantCards: "3333", Source: "3333A"},
			{CardType: "Four of a kind", Power: 2, RelevantCards: "AAAA", Source: "AAAA2"},
		}, loser: 0, name: "Four of a kind - High card"},
	}

	for _, test := range tests {
		loser := compare(test.ratings[0], test.ratings[1], false, orderOfCardsNoJoker)

		if loser != test.ratings[test.loser] {
			t.Errorf("Test Case: %s\nExpected: %v\nGot: %v", test.name, test.ratings[test.loser], loser)
		}
	}

}

func TestExample1(t *testing.T) {

	inputs := map[string]struct {
		cards string
		bid   int
	}{
		"32T3K": {cards: "32T3K", bid: 765},
		"T55J5": {cards: "T55J5", bid: 684},
		"KK677": {cards: "KK677", bid: 28},
		"KTJJT": {cards: "KTJJT", bid: 220},
		"QQQJA": {cards: "QQQJA", bid: 483},
	}

	var ratings []Rating

	// determine type and power of each hand
	for _, value := range inputs {

		ratings = append(ratings, determineTypeAndPower(value.cards, orderOfCardsNoJoker))
	}

	// sort ratings by power using our compare function going from low rank to high rank
	sort.Slice(ratings, func(i, j int) bool {
		return compare(ratings[i], ratings[j], false, orderOfCardsNoJoker) == ratings[i]
	})

	score := 0
	for i, rating := range ratings {
		score += (i + 1) * inputs[rating.Source].bid
	}

	expectedScore := 6440

	if score != expectedScore {
		t.Errorf("Expected score: %d, Got: %d", expectedScore, score)
	}

}

func TestData1(t *testing.T) {

	lines, _ := util.ReadFile("testdata/part1.txt")

	// Each line is 529J8 290
	// Thats the card drawn and the bid

	var inputs = make(map[string]struct {
		cards string
		bid   int
	})

	for _, line := range lines {

		convbid, _ := strconv.Atoi(line[6:])
		cards := line[0:5]

		inputs[cards] = struct {
			cards string
			bid   int
		}{cards: cards, bid: convbid}
	}

	var ratings []Rating

	// determine type and power of each hand
	for _, value := range inputs {

		ratings = append(ratings, determineTypeAndPower(value.cards, orderOfCardsNoJoker))
	}

	// sort ratings by power using our compare function going from low rank to high rank
	sort.Slice(ratings, func(i, j int) bool {
		return compare(ratings[i], ratings[j], false, orderOfCardsNoJoker) == ratings[i]
	})

	score := 0
	for i, rating := range ratings {
		score += (i + 1) * inputs[rating.Source].bid
	}

	expectedScore := 247961593

	if score != expectedScore {
		t.Errorf("Expected score: %d, Got: %d", expectedScore, score)
	}

}
func TestReplaceJokers(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{

		{input: "J567J", expectedOutput: "75677"},
		{input: "JJKJQ", expectedOutput: "KKKKQ"},
		{input: "JJJJJ", expectedOutput: "AAAAA"},
		{input: "JJJAA", expectedOutput: "AAAAA"},
		{input: "JJJTT", expectedOutput: "TTTTT"},
		{input: "JJJ99", expectedOutput: "99999"},
		{input: "JJJ88", expectedOutput: "88888"},
		{input: "JJJ77", expectedOutput: "77777"},
		{input: "JJJ66", expectedOutput: "66666"},
		{input: "JJJ55", expectedOutput: "55555"},
		{input: "JJAAA", expectedOutput: "AAAAA"},
		{input: "T55J5", expectedOutput: "T5555"},
		{input: "QQQJA", expectedOutput: "QQQQA"},

		// test full houses
		{input: "JJJTT", expectedOutput: "TTTTT"},
		{input: "JJJ99", expectedOutput: "99999"},
		{input: "JJJ88", expectedOutput: "88888"},
		{input: "JJJ77", expectedOutput: "77777"},
		{input: "JJJ66", expectedOutput: "66666"},
		{input: "JJJ55", expectedOutput: "55555"},
		{input: "JJJ44", expectedOutput: "44444"},
		{input: "JJJ33", expectedOutput: "33333"},
		{input: "JJJ22", expectedOutput: "22222"},

		{input: "JJJAA", expectedOutput: "AAAAA"},
		{input: "JJJKK", expectedOutput: "KKKKK"},
		{input: "JJJQQ", expectedOutput: "QQQQQ"},
		{input: "JJJTT", expectedOutput: "TTTTT"},
		{input: "JJJ99", expectedOutput: "99999"},
		{input: "JJJ88", expectedOutput: "88888"},
		{input: "JJJ77", expectedOutput: "77777"},
		{input: "JJJ66", expectedOutput: "66666"},
		{input: "JJJ55", expectedOutput: "55555"},
		{input: "JJJ44", expectedOutput: "44444"},
		{input: "JJJ33", expectedOutput: "33333"},
		{input: "JJJ22", expectedOutput: "22222"},

		{input: "JJJAA", expectedOutput: "AAAAA"},
		{input: "JJKKK", expectedOutput: "KKKKK"},
		{input: "JJQQQ", expectedOutput: "QQQQQ"},
		{input: "JJTTT", expectedOutput: "TTTTT"},
		{input: "JJ999", expectedOutput: "99999"},
		{input: "JJ888", expectedOutput: "88888"},
		{input: "JJ777", expectedOutput: "77777"},
		{input: "JJ666", expectedOutput: "66666"},
		{input: "JJ555", expectedOutput: "55555"},
		{input: "JJ444", expectedOutput: "44444"},
		{input: "JJ333", expectedOutput: "33333"},
	}

	for _, test := range tests {
		output := replaceJokers(test.input)
		if output != test.expectedOutput {
			t.Errorf("Input: %s\nExpected Output: %s\nGot: %s", test.input, test.expectedOutput, output)
		}
	}
}

func TestExample2(t *testing.T) {

	inputs := map[string]struct {
		cards string
		bid   int
	}{
		"32T3K": {cards: "32T3K", bid: 765},
		"T55J5": {cards: "T55J5", bid: 684},
		"KK677": {cards: "KK677", bid: 28},
		"KTJJT": {cards: "KTJJT", bid: 220},
		"QQQJA": {cards: "QQQJA", bid: 483},
	}

	var ratings []Rating

	// determine type and power of each hand
	for _, value := range inputs {

		r := determineTypeAndPower(replaceJokers(value.cards), orderOfCardsJoker)
		r.BeforeJokerReplace = value.cards

		ratings = append(ratings, r)
	}

	// sort ratings by power using our compare function going from low rank to high rank
	sort.Slice(ratings, func(i, j int) bool {
		return compare(ratings[i], ratings[j], false, orderOfCardsJoker) == ratings[i]
	})

	score := 0
	for i, rating := range ratings {

		bid := inputs[rating.BeforeJokerReplace].bid
		score += (i + 1) * bid
	}

	expectedScore := 5905

	if score != expectedScore {
		t.Errorf("Expected score: %d, Got: %d", expectedScore, score)
	}

}

func TestData2(t *testing.T) {

	lines, _ := util.ReadFile("testdata/part1.txt")

	var inputs = make(map[string]struct {
		cards string
		bid   int
	})

	for _, line := range lines {

		convbid, _ := strconv.Atoi(line[6:])
		cards := line[0:5]

		inputs[cards] = struct {
			cards string
			bid   int
		}{cards: cards, bid: convbid}
	}

	var ratings []Rating

	// determine type and power of each hand
	for _, value := range inputs {

		r := determineTypeAndPower(replaceJokers(value.cards), orderOfCardsJoker)
		r.BeforeJokerReplace = value.cards

		ratings = append(ratings, r)
	}

	// sort ratings by power using our compare function going from low rank to high rank
	sort.Slice(ratings, func(i, j int) bool {
		return compare(ratings[i], ratings[j], false, orderOfCardsJoker) == ratings[i]
	})

	score := 0
	for i, rating := range ratings {
		log.Printf("Rating: %+v", rating)
		bid := inputs[rating.BeforeJokerReplace].bid
		score += (i + 1) * bid
	}

	expectedScore := 248750699

	if score != expectedScore {
		t.Errorf("Expected score: %d, Got: %d", expectedScore, score)
	}

}
