use std::cmp::Ordering;
use std::collections::HashMap;

/// Given a list of poker hands, return a list of those hands which win.
///
/// Note the type signature: this function should return _the same_ reference to
/// the winning hand(s) as were passed in, not reconstructed strings which happen to be equal.
pub fn winning_hands<'a>(hands: &[&'a str]) -> Option<Vec<&'a str>> {
    let hands = hands.iter()
        .map(|hand| Hand::from(*hand))
        .collect::<Vec<Hand>>();

    let mut best_hand = &hands[0];
    let mut best_hands: Vec<&Hand> = vec![best_hand];
    for i in 1..hands.len() {
        match hands[i].partial_cmp(best_hand) {
            Some(Ordering::Less) => {},
            Some(Ordering::Equal) => best_hands.push(&hands[i]),
            Some(Ordering::Greater) => {
                best_hand = &hands[i];
                best_hands = vec![best_hand];
            },
            None => {},
        }
    }

    let winning_hands = best_hands.iter()
        .map(|h| h.input)
        .collect();

    Some(winning_hands)
}

#[derive(Debug, Eq, PartialEq, Copy, Clone)]
pub enum Suit {
    Spade,
    Club,
    Heart,
    Diamond,
}

#[derive(Debug, Copy, Clone)]
pub struct Card {
    rank: u8,
    suit: Suit,
}

impl From<&str> for Card {
    fn from(input: &str) -> Self {
        let str_rank = &input[..(input.len()-1)];
        let rank: u8 = match str_rank {
            "A" => 14,
            "K" => 13,
            "Q" => 12,
            "J" => 11,
            _ => str_rank.parse().unwrap_or_else(|_| panic!("Invalid rank")),
        };
        let char_suit = input.chars().last().unwrap();
        let suit = match char_suit {
            'S' => Suit::Spade,
            'C' => Suit::Club,
            'H' => Suit::Heart,
            'D' => Suit::Diamond,
            _ => panic!("Invalid suit")
        };

        Self { rank: rank, suit: suit }
    }
}

impl PartialEq for Card {
    fn eq(&self, other: &Self) -> bool {
        self.rank.eq(&other.rank)
    }
}

impl PartialOrd for Card {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        self.rank.partial_cmp(&other.rank)
    }
}

#[derive(Debug, Eq, PartialEq, PartialOrd, Clone)]
pub enum HandRank {
    HighCard(Vec<u8>),
    Pair(Vec<u8>),
    TwoPairs(Vec<u8>),
    ThreeOfAKind(Vec<u8>),
    Straight(Vec<u8>),
    Flush(Vec<u8>),
    FullHouse(Vec<u8>),
    FourOfAKind(Vec<u8>),
    StraightFlush(Vec<u8>),
}

pub struct Hand<'a> {
    input: &'a str,
    hand_rank: HandRank,
}

impl<'a> From<&'a str> for Hand<'a> {
    fn from(input: &'a str) -> Self {
        let cards = input
            .split_whitespace()
            .map(|s| Card::from(s))
            .collect();

        Self {input: input, hand_rank: Hand::evaluate_hand_rank(cards)}
    }
}

impl<'a> Hand<'a> {
    fn evaluate_hand_rank(cards: Vec<Card>) -> HandRank {
        let (counts, mut card_ranks) = Hand::get_counts_and_ranks(&cards);

        if counts[0] == 4 {
            return HandRank::FourOfAKind(card_ranks)
        }

        if counts[0] == 3 && counts[1] == 2 {
            return HandRank::FullHouse(card_ranks);
        }

        if counts[0] == 3 {
            return HandRank::ThreeOfAKind(card_ranks);
        }

        if counts[0] == 2 && counts[1] == 2 {
            return HandRank::TwoPairs(card_ranks);
        }

        if counts.len() == 4 {
            return HandRank::Pair(card_ranks)
        }

        let suit = cards.first().unwrap().suit;
        let is_flush = cards.iter().all(|c| c.suit == suit);
        let is_straight = (card_ranks[0] - card_ranks[4] == 4) || (card_ranks[0] == 14 && card_ranks[1] == 5);
        if is_straight && card_ranks[0] == 14 {
            card_ranks[0] = 1;
        }

        if is_straight && is_flush {
            return HandRank::StraightFlush(card_ranks);
        }

        if is_straight {
            return HandRank::Straight(card_ranks);
        }

        if is_flush {
            return HandRank::Flush(card_ranks);
        }

        HandRank::HighCard(card_ranks)
    }

    fn get_counts_and_ranks(cards: &Vec<Card>) -> (Vec<u8>, Vec<u8>) {
        let mut counts_by_rank: HashMap::<u8, u8> = HashMap::new();
        for c in cards {
            *counts_by_rank.entry(c.rank).or_insert(0) += 1
        }

        let mut counts_and_ranks: Vec<(u8,u8)> = counts_by_rank
            .iter()
            .map(|(k, v)| (*k, *v))
            .collect();

        counts_and_ranks.sort_by(|a,b| {
            if b.1 == a.1 {
                return b.0.cmp(&a.0);
            }

            b.1.cmp(&a.1)
        });

        let ranks: Vec<u8> = counts_and_ranks.iter().map(|cr| cr.0).collect();
        let counts: Vec<u8> = counts_and_ranks.iter().map(|cr| cr.1).collect();

        (counts, ranks)
    }
}

impl<'a> PartialEq for Hand<'a> {
    fn eq(&self, other: &Self) -> bool {
        self.hand_rank == other.hand_rank
    }
}

impl<'a> PartialOrd for Hand<'a> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        self.hand_rank.partial_cmp(&other.hand_rank)
    }
}
