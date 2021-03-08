package poker

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type suit int

const (
	Diamond suit = iota
	Heart
	Club
	Spade
)

type card struct {
	rank int
	suit suit
}

func newCard(c string) (*card, error) {
	runes := []rune(c)
	rank, err := parseCardRank(string(runes[0 : len(runes)-1]))
	if err != nil {
		return nil, err
	}

	suit, err := parseSuit(runes[len(runes)-1])
	if err != nil {
		return nil, err
	}

	return &card{rank: rank, suit: suit}, nil
}

func parseCardRank(input string) (int, error) {
	var rank int
	switch input {
	case "J":
		rank = 11
	case "Q":
		rank = 12
	case "K":
		rank = 13
	case "A":
		rank = 14
	default:
		parsedValue, err := strconv.Atoi(input)
		if err != nil || parsedValue < 2 || parsedValue > 10 {
			return 0, fmt.Errorf("%v is an invalid card rank", input)
		}
		rank = parsedValue
	}

	return rank, nil
}

func parseSuit(input rune) (s suit, err error) {
	switch input {
	case '♢':
		s = Diamond
	case '♡':
		s = Heart
	case '♤':
		s = Club
	case '♧':
		s = Spade
	default:
		err = fmt.Errorf("%c is an invalid suit", input)
	}
	return
}

type handRank int

const (
	HighCard handRank = iota
	Pair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type hand struct {
	input            string
	handRank         handRank
	orderedCardRanks []int
}

func newHand(input string) (*hand, error) {
	var cardInputs = strings.Split(input, " ")
	if len(cardInputs) < 5 {
		return nil, errors.New("too few cards")
	}

	if len(cardInputs) > 5 {
		return nil, errors.New("too many cards")
	}

	cards := make([]*card, 0)
	for _, ci := range cardInputs {
		c, err := newCard(ci)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].rank < cards[j].rank
	})

	handRank, cardRanks := getHandAndCardRanks(cards)

	return &hand{input: input, handRank: handRank, orderedCardRanks: cardRanks}, nil
}

func getHandAndCardRanks(cards []*card) (handRank, []int) {
	counts, cardRanks := getCountsAndCardRanks(cards)

	if counts[0] == 4 {
		return FourOfAKind, cardRanks
	}

	if counts[0] == 3 && counts[1] == 2 {
		return FullHouse, cardRanks
	}

	if counts[0] == 3 {
		return ThreeOfAKind, cardRanks
	}

	if counts[0] == 2 && counts[1] == 2 {
		return TwoPairs, cardRanks
	}

	if len(counts) == 4 {
		return Pair, cardRanks
	}

	var isFlush = true
	suit := cards[0].suit
	for i := 1; i < len(cards); i++ {
		if cards[i].suit != suit {
			isFlush = false
			break
		}
	}
	var isStraight = (cardRanks[0]-cardRanks[4] == 4) || (cardRanks[0] == 14 && cardRanks[1] == 5)
	if isStraight && cardRanks[0] == 14 {
		cardRanks[0] = 1
	}

	if isStraight && isFlush {
		return StraightFlush, cardRanks
	}

	if isStraight {
		return Straight, cardRanks
	}

	if isFlush {
		return Flush, cardRanks
	}

	return HighCard, cardRanks
}

type rankAndCount struct {
	rank  int
	count int
}

func compareRankAndCounts(a, b rankAndCount) bool {
	if a.count == b.count {
		return a.rank > b.rank
	}

	return a.count > b.count
}

func getCountsAndCardRanks(cards []*card) ([]int, []int) {
	countsByRank := make(map[int]rankAndCount)
	for _, c := range cards {
		countByRank, ok := countsByRank[c.rank]
		if ok {
			countByRank.count++
			countsByRank[c.rank] = countByRank
		} else {
			countsByRank[c.rank] = rankAndCount{rank: c.rank, count: 1}
		}
	}

	rankAndCounts := make([]rankAndCount, 0)
	for _, v := range countsByRank {
		rankAndCounts = append(rankAndCounts, v)
	}

	sort.Slice(rankAndCounts, func(i, j int) bool {
		return compareRankAndCounts(rankAndCounts[i], rankAndCounts[j])
	})

	counts := make([]int, 0)
	cardRanks := make([]int, 0)
	for i := 0; i < len(rankAndCounts); i++ {
		counts = append(counts, rankAndCounts[i].count)
		cardRanks = append(cardRanks, rankAndCounts[i].rank)
	}

	return counts, cardRanks
}

func compareHands(a, b *hand) int {
	if a.handRank > b.handRank {
		return 1
	} else if a.handRank < b.handRank {
		return -1
	}

	for i := range a.orderedCardRanks {
		if a.orderedCardRanks[i] > b.orderedCardRanks[i] {
			return 1
		} else if a.orderedCardRanks[i] < b.orderedCardRanks[i] {
			return -1
		}
	}

	return 0
}

// BestHand given poker hands returns best hands
func BestHand(inputs []string) ([]string, error) {
	bestHands := make([]string, 0)
	allHands := make([]*hand, 0)
	for _, input := range inputs {
		h, err := newHand(input)
		if err != nil {
			return bestHands, err
		}
		allHands = append(allHands, h)
	}

	bestHand := allHands[0]
	bestHands = append(bestHands, bestHand.input)
	for i := 1; i < len(allHands); i++ {
		r := compareHands(allHands[i], bestHand)
		if r > 0 {
			bestHands = bestHands[:0]
			bestHand = allHands[i]
			bestHands = append(bestHands, bestHand.input)
		} else if r == 0 {
			bestHands = append(bestHands, allHands[i].input)
		}
	}

	return bestHands, nil
}
