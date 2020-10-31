use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    let mut visited_letters: HashSet<char> = HashSet::new();
    for c in candidate.chars() {
        match c {
            ' ' | '-' => continue,
            _ => (),
        }
        match visited_letters.get(&c.to_ascii_lowercase()) {
            Some(ch) => return false,
            None => { visited_letters.insert(c.to_ascii_lowercase()) }
        };
    }
    true
}
