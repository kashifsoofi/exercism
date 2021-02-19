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
            Value = ParseValue(card.Substring(0, card.Length - 1));
            Suit = ParseSuit(card[card.Length - 1]);
        }

        public int Value { get; }
        public Suit Suit { get; }

        private int ParseValue(string value)
        {
            switch (value)
            {
                case "J": return 11;
                case "Q": return 12;
                case "K": return 13;
                case "A": return 14;
                default: return int.Parse(value);
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

    public class Hand
    {
        private readonly Card[] cards;
        private readonly int handScore;
        public Hand(string input)
        {
            Input = input;
            cards = input.Split(" ").Select(ci => new Card(ci)).OrderBy(c => c.Value).ToArray();
            handScore = CalculateHandScore();
        }

        public string Input { get; }
        public int Score => handScore;

        private int CalculateHandScore()
        {
            var handScore = 0;
            for (var i = cards.Length; i > 0; i--)
            {
                var cardScore = cards[i-1].Value - 2;
                handScore += cardScore * (int)Math.Pow(13, i);
            }
            return handScore;
        }
    }

    public static IEnumerable<string> BestHands(IEnumerable<string> hands)
    {
        var allHands = hands.Select(input => new Hand(input)).ToList();
        var bestScore = allHands.Max(h => h.Score);
        return allHands.Where(h => h.Score == bestScore).Select(h => h.Input);
    }
}