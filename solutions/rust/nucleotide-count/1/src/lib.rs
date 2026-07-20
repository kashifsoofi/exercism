use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    Ok(*nucleotide_counts(dna)?.get(&nucleotide).ok_or(nucleotide)?)
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts = "ACGT".chars().map(|c| (c, 0_usize)).collect::<HashMap<_, _>>();
    for c in dna.chars() {
        match c {
            'A' | 'C' | 'G' | 'T' => {
                let count = counts.entry(c).or_insert(0);
                *count += 1;
            },
            _ => return Err(c)
        }
    }
    Ok(counts)
}
