package camel

import (
	"sort"
	"strings"
)

type Rating struct {
	CardType           string
	Power              int
	RelevantCards      string
	Source             string
	BeforeJokerReplace string
}

var orderOfCardsNoJoker = map[string]int{

	"A": 1,
	"K": 2,
	"Q": 3,
	"J": 4,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13,
	"1": 14,
}

var orderOfCardsJoker = map[string]int{

	"A": 1,
	"K": 2,
	"Q": 3,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13,
	"J": 15,
}

func determineTypeAndPower(input string, orderOfCards map[string]int) Rating {

	// sort input string by rune
	s := []rune(input)

	sort.Slice(s, func(i, j int) bool {
		return orderOfCards[string(s[i])] < orderOfCards[string(s[j])]
	})

	// check for five of a kind
	if s[0] == s[4] {

		return Rating{CardType: "Five of a kind", Power: 1, RelevantCards: string(s), Source: input}

	}

	// check for four of a kind
	if s[0] == s[3] {
		return Rating{CardType: "Four of a kind", Power: 2, RelevantCards: string(s[0:3]), Source: input}
	}

	if s[1] == s[4] {
		return Rating{CardType: "Four of a kind", Power: 2, RelevantCards: string(s[1:4]), Source: input}
	}

	// check for full house
	if s[0] == s[2] && s[3] == s[4] {
		return Rating{CardType: "Full house", Power: 3, RelevantCards: string(s[0:3]), Source: input}
	}

	if s[0] == s[1] && s[2] == s[4] {

		return Rating{CardType: "Full house", Power: 3, RelevantCards: string(s[2:4]), Source: input}

	}

	// check for three of a kind
	if s[0] == s[2] {
		return Rating{CardType: "Three of a kind", Power: 4, RelevantCards: string(s[0:3]), Source: input}
	}

	if s[1] == s[3] {
		return Rating{CardType: "Three of a kind", Power: 4, RelevantCards: string(s[1:4]), Source: input}
	}

	if s[2] == s[4] {
		return Rating{CardType: "Three of a kind", Power: 4, RelevantCards: string(s[2:5]), Source: input}
	}

	// check for two pair
	if s[0] == s[1] && s[2] == s[3] {
		return Rating{CardType: "Two pair", Power: 5, RelevantCards: string(s[0:4]), Source: input}
	}

	if s[0] == s[1] && s[3] == s[4] {
		return Rating{CardType: "Two pair", Power: 5, RelevantCards: string(s[0:2]) + string(s[3:5]), Source: input}
	}

	if s[1] == s[2] && s[3] == s[4] {
		return Rating{CardType: "Two pair", Power: 5, RelevantCards: string(s[1:5]), Source: input}
	}

	// check for one pair
	if s[0] == s[1] {
		return Rating{CardType: "One pair", Power: 6, RelevantCards: string(s[0:2]), Source: input}
	}

	if s[1] == s[2] {
		return Rating{CardType: "One pair", Power: 6, RelevantCards: string(s[1:3]), Source: input}
	}

	if s[2] == s[3] {
		return Rating{CardType: "One pair", Power: 6, RelevantCards: string(s[2:4]), Source: input}
	}

	if s[3] == s[4] {
		return Rating{CardType: "One pair", Power: 6, RelevantCards: string(s[3:5]), Source: input}
	}

	// check for highest card
	return Rating{CardType: "Highest card", Power: 7, RelevantCards: string(s[4]), Source: input}

}

func compare(r1, r2 Rating, byRelevantCards bool, orderOfCards map[string]int) Rating {

	// return the rating with the higher power
	if r1.Power < r2.Power {
		return r2
	}

	// rating is the same
	if r1.Power == r2.Power {
		if byRelevantCards {
			if orderOfCards[string(r1.RelevantCards[0])] < orderOfCards[string(r2.RelevantCards[0])] {
				return r2
			}
		} else {

			// if r1.replaceJokers is set then we take these values for comparison
			var s string

			if r1.BeforeJokerReplace != "" {
				s = determineWinner(r1.BeforeJokerReplace, r2.BeforeJokerReplace, orderOfCards)
				if s == r1.BeforeJokerReplace {
					return r1
				} else {
					return r2
				}
			} else {
				s = determineWinner(r1.Source, r2.Source, orderOfCards)
				if s == r1.Source {
					return r1
				} else {
					return r2
				}
			}
		}
		return r1
	}
	return r1
}

func determineWinner(s1, s2 string, orderOfCards map[string]int) string {

	for i := 0; i < len(s1); i++ {

		if orderOfCards[string(s1[i])] < orderOfCards[string(s2[i])] {
			return s2
		}

		if orderOfCards[string(s1[i])] > orderOfCards[string(s2[i])] {
			return s1
		}

	}

	return s1

}

func replaceJokers(s string) string {

	if strings.Contains(s, "J") == false {
		return s
	}

	if strings.Count(s, "J") == 5 {
		return "AAAAA"
	}
	// get string with no jokers
	noJokers := strings.Replace(s, "J", "", -1)

	highestCard := "J"
	for _, c := range noJokers {

		orderOfChar := orderOfCardsJoker[string(c)]

		if orderOfChar < orderOfCardsJoker[highestCard] {
			highestCard = string(c)
		}
	}

	// count remaining letters in a map
	count := make(map[string]int)
	for _, c := range noJokers {
		count[string(c)]++
	}

	mostCounts := 0
	cardValueMaxCounts := ""
	for k, v := range count {
		if v > mostCounts {
			mostCounts = v
			cardValueMaxCounts = k
		}
	}

	// case that we have only different cards left, then we can replace the joker with the highest card
	if len(count) == len(noJokers) {
		return strings.Replace(s, "J", highestCard, -1)
	}

	// case that we have two pairs at 4 cards left, then we can replace the joker with just the highest card
	if len(noJokers) == 4 && len(count) == 2 && count[highestCard] == 2 {
		return strings.Replace(s, "J", highestCard, -1)
	}

	// all the other cases means we need to replace the joker with the card that has the most counts
	s = strings.Replace(s, "J", cardValueMaxCounts, -1)

	return s

}
