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
        private int highestCard;
        public Hand(string input)
        {
            Input = input;
            cards = input.Split(" ").Select(ci => new Card(ci)).ToArray();
            highestCard = cards.Max(c => c.Value);
        }

        public string Input { get; }
        public int HighestCard => highestCard;
    }

    public static IEnumerable<string> BestHands(IEnumerable<string> hands)
    {
        var allHands = hands.Select(input => new Hand(input)).ToList();
        var highestCard = allHands.Max(h => h.HighestCard);
        return allHands.Where(h => h.HighestCard == highestCard).Select(h => h.Input);
    }
}