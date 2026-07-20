use std::collections::HashMap;

/// Count occurrences of words.
pub fn word_count(words: &str) -> HashMap<String, u32> {
    words
        .to_ascii_lowercase()
        .split(|c| c == ' ' || c == ',' || c == '\n')
        .fold(HashMap::new(), |mut map, s| {
            let key = s.trim_matches(|c: char| c.is_ascii_punctuation()).to_string();
            if !key.is_empty() {
                *map.entry(key).or_insert(0) += 1;
            }
            map
        })
}
