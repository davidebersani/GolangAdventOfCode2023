package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type GroupedCard struct {
	Card  string
	Count int
	Value int
}

const (
	FiveOfKind  int = 6
	FourOfKind      = 5
	FullHose        = 4
	ThreeOfKind     = 3
	TwoPairs        = 2
	OnePair         = 1
	HighCard        = 0
)

type PokerHand struct {
	Hand string
	Bid  int
}

func (hand PokerHand) GetGroupedCards() []GroupedCard {
	// Calculate attribute
	cards := strings.Split(hand.Hand, "")
	values := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	// Map cards -> []GroupedCard{}
	cardValues := funk.Map(cards, func(c string) GroupedCard {
		return GroupedCard{Card: c, Count: 1, Value: values[c]}
	})
	// Sort by value
	cardValuesSlice := cardValues.([]GroupedCard)
	sort.Slice(cardValuesSlice, func(i, j int) bool {
		return cardValuesSlice[i].Value > cardValuesSlice[j].Value
	})
	// Reduce []GroupedCard{} -> []GroupedCard{} aumentando il count
	groupedCards := funk.Reduce(cardValuesSlice, func(acc []GroupedCard, card GroupedCard) []GroupedCard {
		if len(acc) == 0 {
			acc = append(acc, card)
			return acc
		}
		if card.Card == acc[len(acc)-1].Card {
			acc[len(acc)-1].Count++
		} else {
			acc = append(acc, card)
		}
		return acc
	}, []GroupedCard{}).([]GroupedCard)
	// Sort by count
	sort.Slice(groupedCards, func(i, j int) bool {
		if groupedCards[i].Count != groupedCards[j].Count {
			return groupedCards[i].Count > groupedCards[j].Count
		}

		return groupedCards[i].Value > groupedCards[j].Value
	})

	return groupedCards
}

func (hand PokerHand) GetType() int {
	groupedCards := hand.GetGroupedCards()

	// Check if there is a J, then try to obtain a higher type
	if funk.ContainsString(strings.Split(hand.Hand, ""), "J") {
		newType := HighCard
		oc := "A"                                // In case the hand contains 5 J (we never know, the rules like to change (cit.))
		for i := 0; i < len(groupedCards); i++ { // Find the first card different from J
			if groupedCards[i].Card != "J" {
				oc = groupedCards[i].Card
				break
			}
		}
		h := PokerHand{Hand: strings.ReplaceAll(hand.Hand, "J", oc), Bid: hand.Bid}
		t := h.GetType()
		if t > newType {
			newType = t
			fmt.Printf("==> Sostituisco J con %v e ottengo la mano %v con rank %v\n", oc, h.Hand, t)
		}
		return newType
	}

	// Otherwise calculate usual type
	lenGroupedCards := len(groupedCards)
	if lenGroupedCards == 1 {
		return FiveOfKind
	}
	if lenGroupedCards == 2 {
		if groupedCards[0].Count == 4 {
			return FourOfKind
		} else {
			return FullHose
		}
	}
	if lenGroupedCards == 3 {
		if groupedCards[0].Count == 3 {
			return ThreeOfKind
		} else {
			return TwoPairs
		}
	}
	if lenGroupedCards == 4 {
		return OnePair
	}
	return HighCard
}

// Define a collection type that implements sort.Interface
type PokerHands []PokerHand

func (hands PokerHands) Len() int {
	return len(hands)
}
func (hands PokerHands) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}
func (hands PokerHands) Less(i, j int) bool {
	if hands[i].GetType() != hands[j].GetType() {
		return hands[i].GetType() < hands[j].GetType()
	}

	values := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	cardsI := strings.Split(hands[i].Hand, "")
	cardsJ := strings.Split(hands[j].Hand, "")
	for index := 0; index < len(cardsI); index++ {
		if values[cardsI[index]] != values[cardsJ[index]] {
			return values[cardsI[index]] < values[cardsJ[index]]
		}
	}

	return false
}

func ConvertToint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Error converting number ", n)
	}
	return n
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		return
	}

	hands := PokerHands{}
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		h := PokerHand{Hand: parts[0], Bid: ConvertToint(parts[1])}
		fmt.Println(h.GetGroupedCards())
		fmt.Println(h.GetType())
		hands = append(hands, h)
	}

	sort.Sort(hands)

	winSum := 0
	for i, h := range hands {
		winSum += h.Bid * (i + 1)
	}

	fmt.Println("Win sum is", winSum)
}
