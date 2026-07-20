pub fn find<T: PartialOrd, A: AsRef<[T]>>(array: A, key: T) -> Option<usize> {
    let slice = array.as_ref();
    if slice.is_empty() {
        return None;
    }

    let (mut left, mut right) = (0, slice.len() - 1);
    while left <= right {
        let index = left + (right - left) / 2;
        if slice[index] == key {
            return Some(index);
        }

        if index != 0 && slice[index] > key {
            right = index - 1;
        }
        else {
            left = index + 1;
        }
    }
    None
}
