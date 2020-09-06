pub fn nth(n: u32) -> u32 {
    if n == 0 {
        return 2;
    }

    let (mut i, mut counter) = (3, 1);

    loop {
        if is_prime(i) {
            counter += 1;
        }

        if counter == n {
            break i;
        }

        i += 2;
    }
}

fn is_prime(n: u32) -> bool {
    if n > 2 && n % 2 == 0 {
        return false;
    }

    let (mut i, limit) = (3, n / 2 + 1);
    loop {
        if n % i == 0 {
            return false;
        }

        if i >= limit {
            break;
        }

        i += 2;
    }

    true
}
