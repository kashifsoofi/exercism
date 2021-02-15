use std::cmp::Ordering;

#[derive(Debug, PartialEq, Eq)]
pub enum Classification {
    Abundant,
    Perfect,
    Deficient,
}

pub fn classify(num: u64) -> Option<Classification> {
    if num < 1 {
        return None;
    }

    let mut aliquot_sum: u64 = 0;
    for i in 1..=num/2 {
        if num % i == 0 {
            aliquot_sum += i;
        }
    }

    let classification = match aliquot_sum.cmp(&num) {
        Ordering::Greater => Classification::Abundant,
        Ordering::Equal => Classification::Perfect,
        Ordering::Less => Classification::Deficient
    };

    Some(classification)
}
