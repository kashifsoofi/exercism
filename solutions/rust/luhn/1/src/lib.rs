/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let trimmed = code.replace(" ", "");
    if trimmed.len() < 2 || !trimmed.chars().all(|ch| ch.is_numeric()) {
        return false;
    }

    trimmed
        .chars()
        .rev()
        .filter_map(|ch| ch.to_digit(10))
        .enumerate()
        .map(|(i, n)| {
            match i % 2 {
                1 => {
                    if n * 2 > 9 {
                        n * 2 - 9
                    } else {
                        n * 2
                    }
                },
                _ => n,
            }
        })
        .sum::<u32>() % 10 == 0

}
