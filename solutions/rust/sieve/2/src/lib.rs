pub fn primes_up_to(upper_bound: u64) -> Vec<u64> {
    let mut primes: Vec<u64> = vec![];
    let mut marked: Vec<bool> = vec![false; 1 + upper_bound as usize];
    for i in 2..marked.len() {
        println!("{}: {}", i, marked[i]);
        if !marked[i] {
            primes.push(i as u64);
            for j in (i * 2..marked.len()).step_by(i) {
                println!("{}", j);
                marked[j as usize] = true;
            }
        }
    }
    primes
}
