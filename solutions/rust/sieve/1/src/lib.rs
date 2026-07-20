pub fn primes_up_to(upper_bound: u64) -> Vec<u64> {
    if upper_bound < 2 {
        return vec![];
    }

    let mut marked: Vec<bool> = vec![false; (upper_bound + 1) as usize];
    let mut primes: Vec<u64> = vec![];
    for i in 2..marked.len() {
        if !marked[i] {
            primes.push(i as u64);
            let mut j: usize = 2;
            let mut index: usize = i * j;
            while index < marked.len() {
                marked[index] = true;
                j += 1;
                index = i * j;
            }
        }
    }
    primes
}
