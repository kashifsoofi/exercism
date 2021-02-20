package poker

import (
	"errors"
	"fmt"
	"math"
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

func parseCardRank(card string) (int, error) {
	var rank int
	switch card {
	case "J":
		rank = 11
	case "Q":
		rank = 12
	case "K":
		rank = 13
	case "A":
		rank = 14
	default:
		parsedValue, err := strconv.Atoi(card)
		if err != nil || parsedValue < 2 || parsedValue > 10 {
			return 0, errors.New(fmt.Sprintf("%v is an invalid card rank", card))
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
		err = errors.New(fmt.Sprintf("%c is an invalid suit", input))
	}
	return
}

type hand struct {
	input string
	cards []*card
	score float64
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

	score := calculateHandScore(cards)

	return &hand{input: input, cards: cards, score: score}, nil
}

func calculateHandScore(cards []*card) float64 {
	score := 0.0
	for i, c := range cards {
		score += (float64(c.rank) - 2) * math.Pow(13, float64(i))
	}

	return score / 402234
}

func BestHand(inputs []string) ([]string, error) {
	bestHands := make([]string, 0)
	hands := make([]*hand, 0)
	for _, input := range inputs {
		h, err := newHand(input)
		if err != nil {
			return bestHands, err
		}
		hands = append(hands, h)
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].score > hands[j].score
	})

	maxScore := hands[0].score
	for _, h := range hands {
		if h.score == maxScore {
			bestHands = append(bestHands, h.input)
		}
	}

	return bestHands, nil
}
