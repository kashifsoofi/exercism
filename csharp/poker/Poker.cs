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
            Suit = ParseSuit(card[^1]);
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

        public Hand(string input)
        {
            Input = input;
            cards = input.Split(" ").Select(ci => new Card(ci)).OrderBy(c => c.Value).ToArray();
            Score = CalculateHandScore();
        }

        public string Input { get; }
        public double Score { get; }

        private double CalculateHandScore()
        {
            var pair1Value = -1;
            var pair2Value = -1;
            for (var i = 0; i < cards.Length - 1; i++)
            {
                var currentValue = cards[i].Value;
                var nextValue = cards[i+1].Value;

                if (pair1Value == -1 && currentValue == nextValue)
                {
                    pair1Value = currentValue;
                    continue;
                }

                if (pair2Value == -1 && currentValue != pair1Value && currentValue == nextValue)
                {
                    pair2Value = currentValue;
                    continue;
                }
            }

            var handScore = 0.0;
            for (var i = cards.Length; i > 0; i--)
            {
                if (cards[i-1].Value == pair1Value || cards[i-1].Value == pair2Value)
                {
                    continue;
                }

                var cardScore = cards[i-1].Value - 2;
                handScore += cardScore * Math.Pow(13, i);
            }

            handScore /= 433175;

            if (pair1Value != -1)
            {
                handScore += 100 + pair1Value / 14.0 * 50;
            }

            if (pair2Value != -1)
            {
                handScore += 100 + pair2Value / 14.0 * 50;
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