pub fn nth(n: u32) -> u32 {
    let (mut i, mut counter) = (0, 0);

    loop {
        if is_prime(i) {
            counter += 1;
        }

        if counter == n + 1 {
            break i;
        }

        i += 1;
    }
}

fn is_prime(n: u32) -> bool {
    if n < 2 {
        return false;
    }

    for i in 2..n {
        if n % i == 0 {
            return false;
        }
    }

    true
}
