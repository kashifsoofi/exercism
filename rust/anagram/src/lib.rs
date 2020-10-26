use std::collections::HashMap;
use std::collections::HashSet;
use unicode_segmentation::UnicodeSegmentation;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let word_letter_counts = get_letter_counts(word);

    possible_anagrams
        .iter()
        .filter(|candidate| candidate.to_lowercase() != word.to_lowercase())
        .fold(HashSet::new(), |mut anagrams, candidate| {
            let candidate_letter_counts = get_letter_counts(candidate);
            if candidate_letter_counts == word_letter_counts {
                anagrams.insert(candidate);
            }
            anagrams
        })
}

fn get_letter_counts(string: &str) -> HashMap<String, u32> {
    string
        .to_lowercase()
        .graphemes(true)
        .map(|grapheme| grapheme.to_string())
        .fold(HashMap::new(), |mut map, grapheme| {
            *map.entry(grapheme).or_insert(0) += 1;
            map
        })
}
