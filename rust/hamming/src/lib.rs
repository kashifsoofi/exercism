/// Return the Hamming distance between the strings,
/// or None if the lengths are mismatched.
pub fn hamming_distance(s1: &str, s2: &str) -> Option<usize> {
    println!("{}:{}", s1.len(), s2.len());
    match s1.len() == s2.len() {
        false => {
            println!("false");
            None
        },
        true => {
            println!("true");
            let mut distance = 0;
            for i in 0..s1.len() {
                println!("{}", i);
                let c1 = s1.chars().nth(i).unwrap();
                let c2 = s2.chars().nth(i).unwrap();
                println!("{}:{}", c1, c2);
                if c1 != c2 {
                    distance += 1;
                }
            }
            Some(distance)
        }
    }
}
