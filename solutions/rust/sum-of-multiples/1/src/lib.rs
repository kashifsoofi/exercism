use std::collections::HashSet;

pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let mut multiples : HashSet<u32> = HashSet::new();
    for factor in factors.iter() {
        let mut index = 1;
        while factor > &0 && factor * index < limit {
            multiples.insert(factor * index);
            index += 1;
        }
    }

    let mut sum = 0;
    for n in &multiples {
        sum += n;
    }
    sum
}
