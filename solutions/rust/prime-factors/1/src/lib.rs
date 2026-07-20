pub fn factors(n: u64) -> Vec<u64> {
    let mut factors : Vec<u64> = Vec::new();
    let mut number = n;
    let mut factor = 2;
    while number > 1 {
        while number % factor == 0 {
            factors.push(factor);
            number = number / factor;
        }
        factor += 1;
    }
    return factors;
}
