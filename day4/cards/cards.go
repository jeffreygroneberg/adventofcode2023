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

// This function is doing the following:
// 1. Get the number of hits for each card: This is our winning table
// We have then in our wins for each card the number of hits (careful: index 0 is card 1 and so on!!)

// 2. We create a distribution rate for each card with 1
// We are just going to keep track of the number of cards we have.

// 3. We are now taking each slot of distribution rate and determine based on our winning table how many copies we need to create for each card
// We do this simply by determing the next cards we have to consider (thats just the number of hits for the current card) and how often we have to this
// (thats just the number of cards we have in the distribution rate for the current card)

// 4. We are now creating the copies and increase the counter for each id we get back by one.

// 5. We are now returning the distribution rate which is our result and from which we just have to calculate the certain values for each card (each index)
func GetCardIdsWithWinningCopyCondition(inputCards []Card) []int {

	// sort inputcards by card id
	sort.Slice(inputCards, func(i, j int) bool {
		return inputCards[i].CardId < inputCards[j].CardId
	})

	// init distribution rate for each card with 1
	var distributionRate []int = make([]int, len(inputCards))
	for i := 0; i < len(distributionRate); i++ {
		distributionRate[i] = 1
	}

	// create an array and just put the len of the wins in there this will be used as mapping table to create copies
	var wins []int
	for _, card := range inputCards {
		numberOfHits, _ := GetHittingNumbersWithPower(card)
		wins = append(wins, len(numberOfHits))
	}

	for i := 0; i < len(distributionRate); i++ {

		ids := GetCopies(i, i+wins[i], distributionRate[i])

		for _, id := range ids {
			distributionRate[id] = distributionRate[id] + 1
		}

	}

	return distributionRate

}

func GetCopies(from, to, times int) []int {

	var ids []int

	for i := 0; i < times; i++ {
		for j := from + 1; j <= to; j++ {
			ids = append(ids, j)
		}
	}

	return ids

}
