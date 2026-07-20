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
        public Card(int rank, Suit suit)
        {
            Rank = rank;
            Suit = suit;
        }

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
        private readonly Card[] rankedCards;
        private readonly HandRank handRank;

        public Hand(string input)
        {
            Input = input;
            var cards = input.Split(" ").Select(ci => new Card(ci)).OrderBy(c => c.Rank).ToArray();
            (handRank, rankedCards) = GetRank(cards);
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

            for (var i = 0; i < rankedCards.Length; i++)
            {
                if (rankedCards[i].Rank > other.rankedCards[i].Rank)
                {
                    return 1;
                }
                else if (rankedCards[i].Rank < other.rankedCards[i].Rank)
                {
                    return -1;
                }
            }

            return 0;
        }

        private (HandRank, Card[]) GetRank(Card[] cards)
        {
            var cardsByCount = cards.GroupBy(c => c.Rank).Select(g => new { count = g.Count(), card = g.First() } ).OrderByDescending(c => c.count).ThenByDescending(c => c.card.Rank).ToList();
            var rankedCards = cardsByCount.Select(c => c.card).ToArray();
            if (cardsByCount[0].count == 4)
            {
                return (HandRank.FourOfAKind, rankedCards);
            }

            if (cardsByCount[0].count == 3 && cardsByCount[1].count == 2)
            {
                return (HandRank.FullHouse, rankedCards);
            }

            if (cardsByCount[0].count == 3)
            {
                return (HandRank.ThreeOfAKind, rankedCards);
            }

            if (cardsByCount[0].count == 2 && cardsByCount[1].count == 2)
            {
                return (HandRank.TwoPairs, rankedCards);
            }

            if (cardsByCount.Count == 4)
            {
                return (HandRank.Pair, rankedCards);
            }

            var isFlush = cards.GroupBy(c => c.Suit).Count() == 1;
            var isStraight = (rankedCards[0].Rank == 14 && rankedCards[1].Rank == 5) || (rankedCards[0].Rank - rankedCards[4].Rank == 4);
            if (isStraight && rankedCards[0].Rank == 14)
            {
                rankedCards[0] = new Card(1, rankedCards[0].Suit);
            }

            if (isStraight && isFlush)
            {
                return (HandRank.StraightFlush, rankedCards);
            }
            else if (isStraight)
            {
                return (HandRank.Straight, rankedCards);
            }
            else if (isFlush)
            {
                return (HandRank.Flush, rankedCards);
            }

            return (HandRank.HighCard, rankedCards);
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