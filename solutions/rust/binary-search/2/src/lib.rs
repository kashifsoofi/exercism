pub fn find<A, T>(array: A, key: T) -> Option<usize>
where A: AsRef<[T]>, T: PartialOrd {
    let slice = array.as_ref();
    if slice.is_empty() {
        return None;
    }

    let (mut left, mut right) = (0, slice.len() - 1);
    while left <= right {
        let index = left + (right - left) / 2;
        match slice(index).cmp(key) {
            Ordering::Equal => { return Some(index); },
            Ordering::Greater => { left = index + 1; }
            Ordering::Less => {
                if index != 0 {
                    right = index - 1;
                }
            }
        }
    }
    None
}
