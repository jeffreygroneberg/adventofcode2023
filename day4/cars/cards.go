package cards

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	CardId         int
	WinningNumbers []int
	Draws          []int
}

func ExtractCardIdWinningNumbersAndDraws(line string) Card {

	split := strings.Split(line, ":")

	i := GetCardId(split[0])
	winningNumbers, draws := GetWinningNumbersAndDraws(split[1])

	card := Card{
		CardId:         i,
		WinningNumbers: winningNumbers,
		Draws:          draws,
	}

	return card

}

func GetCardId(s string) int {

	// Get Id from string "Card   1"
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")

	// last element of split is the id
	id := split[len(split)-1]

	i, _ := strconv.Atoi(id)

	return i

}

func GetWinningNumbersAndDraws(s string) ([]int, []int) {

	// trim spaces first
	s = strings.TrimSpace(s)
	split := strings.Split(s, "|")

	// remove double whitespaces
	winningNumbers := strings.Split(strings.Join(strings.Fields(split[0]), " "), " ")
	draws := strings.Split(strings.Join(strings.Fields(split[1]), " "), " ")

	var winningNumbersInt []int
	var drawsInt []int

	for _, winningNumber := range winningNumbers {
		i, _ := strconv.Atoi(winningNumber)
		winningNumbersInt = append(winningNumbersInt, i)
	}

	for _, draw := range draws {
		i, _ := strconv.Atoi(draw)
		drawsInt = append(drawsInt, i)
	}

	return winningNumbersInt, drawsInt

}

func GetHittingNumbersWithPower(testCard Card) ([]int, int) {

	// check for all numbers in draws if they are in winning numbers

	var wins []int

	for _, draw := range testCard.Draws {

		for _, winningNumber := range testCard.WinningNumbers {

			if draw == winningNumber {
				wins = append(wins, draw)
			}
		}
	}

	return wins, int(math.Pow(float64(2), float64(len(wins)-1)))

}

// assuming the array with cards is sorted by card id
func GetCardIdsWithWinningCopyCondition(inputCards []Card) []int {

	var extractedIds []int

	// sort inputcards by card id
	sort.Slice(inputCards, func(i, j int) bool {
		return inputCards[i].CardId < inputCards[j].CardId
	})

	// copy all cards with only their id
	for _, card := range inputCards {
		extractedIds = append(extractedIds, card.CardId)
	}

	// create an array and just put the len of the wins in there this will be used as mapping table to create copies
	var wins []int

	for _, card := range inputCards {

		numberOfHits, _ := GetHittingNumbersWithPower(card)
		wins = append(wins, len(numberOfHits))

	}

	sameNumberCount := 0
	lastNumber := extractedIds[0]

	for i := 0; i < len(extractedIds); i++ {

		currentNumber := extractedIds[i]

		// new number add wins of last number to the slice
		if lastNumber != currentNumber {
			for j := 0; j < sameNumberCount; j++ {
				extractedIds = AddIdsToSlice(lastNumber, wins[lastNumber-1], extractedIds)
			}
			sameNumberCount = 1
		} else {
			sameNumberCount++
		}

		lastNumber = extractedIds[i]

	}

	return extractedIds

}

func AddIdsToSlice(startIndex int, nextNumbers int, ids []int) []int {

	// protection to overextend the slice. Get the last number if the slice first as maximum
	maxNumber := ids[len(ids)-1]

	for i := startIndex + 1; i <= startIndex+nextNumbers && i <= maxNumber; i++ {

		ids = InsertSorted(ids, i)

	}

	return ids

}

func InsertSorted(s []int, e int) []int {
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}
