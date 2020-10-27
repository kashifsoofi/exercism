/// Return the Hamming distance between the strings,
/// or None if the lengths are mismatched.
pub fn hamming_distance(s1: &str, s2: &str) -> Option<usize> {
    match s1.len() == s2.len() {
        false => None,
        true => Some(s1.chars().zip(s2.chars()).map(|(c1, c2)| if c1 == c2 { 0 } else { 1 }).sum())
    }
}
