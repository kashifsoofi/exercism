using System;
using System.Collections.Generic;
using System.Linq;

public static class Poker
{
    public enum Suit
    {
        Spade,
        Club,
        Heart,
        Diamond,
    }

    public class Card
    {
        public Card(string card)
        {
            Rank = ParseRank(card.Substring(0, card.Length - 1));
            Suit = ParseSuit(card[card.Length - 1]);
        }

        public int Rank { get; }
        public Suit Suit { get; }

        private int ParseRank(string rank)
        {
            switch (rank)
            {
                case "J": return 11;
                case "Q": return 12;
                case "K": return 13;
                case "A": return 14;
                default: return int.Parse(rank);
            }
        }

        private Suit ParseSuit(char value)
        {
            switch (value)
            {
                case 'S': return Suit.Spade;
                case 'C': return Suit.Club;
                case 'H': return Suit.Heart;
                case 'D': return Suit.Diamond;
                default: throw new ArgumentException("Invalid Suit");
            }
        }
    }

    public enum HandRank
    {
        HighCard,
        Pair,
        TwoPairs,
        ThreeOfAKind,
        Straight,
        Flush,
        FullHouse,
        FourOfAKind,
        StraightFlush,
    }

    public class Hand : IComparable<Hand>
    {
        private readonly Card[] cards;
        private readonly HandRank handRank;

        public Hand(string input)
        {
            Input = input;
            cards = input.Split(" ").Select(ci => new Card(ci)).OrderBy(c => c.Rank).ToArray();
            handRank = GetRank();
        }

        public string Input { get; }

        public int CompareTo(Hand other)
        {
            if (handRank > other.handRank)
            {
                return 1;
            }
            else if (handRank < other.handRank)
            {
                return -1;
            }

            if (handRank == HandRank.FullHouse || handRank == HandRank.FourOfAKind)
            {
                var cardsInGroup = cards.GroupBy(c => c.Rank).OrderByDescending(g => g.Count()).Select(g => g.First()).ToList();
                var otherCardsInGroup = other.cards.GroupBy(c => c.Rank).OrderByDescending(g => g.Count()).Select(g => g.First()).ToList();
                for (var i = 0; i < cardsInGroup.Count(); i++)
                {
                    if (cardsInGroup[i].Rank > otherCardsInGroup[i].Rank)
                    {
                        return 1;
                    }
                    else if (cardsInGroup[i].Rank < otherCardsInGroup[i].Rank)
                    {
                        return -1;
                    }
                }
            }

            for (var i = 4; i >= 0; i--)
            {
                var cardRank = handRank == HandRank.Straight && cards[i].Rank == 14 ? 0 : cards[i].Rank;
                var otherCardRank = handRank == HandRank.Straight && other.cards[i].Rank == 14 ? 0 : other.cards[i].Rank;
                if (cardRank > otherCardRank)
                {
                    return 1;
                }
                else if (cardRank < otherCardRank)
                {
                    return -1;
                }
            }

            return 0;
        }

        private HandRank GetRank()
        {
            var counts = cards.GroupBy(c => c.Rank).Select(g => g.Count()).OrderByDescending(c => c).ToList();
            if (counts[0] == 4)
            {
                return HandRank.FourOfAKind;
            }

            if (counts[0] == 3 && counts[1] == 2)
            {
                return HandRank.FullHouse;
            }

            if (counts[0] == 3)
            {
                return HandRank.ThreeOfAKind;
            }

            if (counts[0] == 2 && counts[1] == 2)
            {
                return HandRank.TwoPairs;
            }

            if (counts.Count == 4)
            {
                return HandRank.Pair;
            }

            var isFlush = cards.GroupBy(c => c.Suit).Count() == 1;
            var isStraight = (cards[4].Rank == 14 && cards[3].Rank == 5) || (cards[4].Rank - cards[0].Rank == 4);
            if (isStraight && isFlush)
            {
                return HandRank.StraightFlush;
            }
            else if (isStraight)
            {
                return HandRank.Straight;
            }
            else if (isFlush)
            {
                return HandRank.Flush;
            }

            return HandRank.HighCard;
        }
    }

    public static IEnumerable<string> BestHands(IEnumerable<string> hands)
    {
        var allHands = hands.Select(input => new Hand(input)).ToList();
        allHands.Sort((x, y) => x.CompareTo(y));
        var bestHand = allHands.First();
        var bestHands = new List<Hand>() { bestHand };
        for (var i = 1; i < allHands.Count; i++)
        {
            var result = allHands[i].CompareTo(bestHand);
            if (result > 0)
            {
                bestHands.Clear();
                bestHand = allHands[i];
                bestHands.Add(bestHand);
            }
            else if (result == 0)
            {
                bestHands.Add(allHands[i]);
            }
        }

        return bestHands.Select(h => h.Input).ToArray();
    }
}