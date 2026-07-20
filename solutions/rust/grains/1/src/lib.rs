pub fn square(s: u32) -> u64 {
    if s < 1 || s > 64 {
        panic!("Square must be between 1 and 64");
    }

    return (1 as u64) << (s as u64) - 1
}

pub fn total() -> u64 {
    let mut total:u64 = 0;
    for i in 1..65 {
        total += square(i)
    }

    total
}
