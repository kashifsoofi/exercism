/// Compute the Scrabble score for a word.
use std::collections::HashMap;

pub fn score(word: &str) -> u64 {
    word.chars().map(|x| x.to_ascii_uppercase())
    .map(letter_score)
    .sum()
}

fn letter_score(c: char) -> u64 {
    match c {
        'A' | 'E' | 'I' | 'O' | 'U' | 'L' | 'N' | 'R' | 'S' | 'T' => 1,
        'D' | 'G' => 2,
        'B' | 'C' | 'M' | 'P' => 3,
        'F' | 'H' | 'V' | 'W' | 'Y' => 4,
        'K' => 5,
        'J' | 'X' => 8,
        'Q' | 'Z' => 10,
        _ => 0,
    }
}