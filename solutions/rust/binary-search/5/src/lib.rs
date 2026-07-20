use ::std::cmp::{Ord, Ordering};

pub fn find<A, T>(array: A, key: T) -> Option<usize>
where
    A: AsRef<[T]>,
    T: Ord,
{
    let slice = array.as_ref();
    let (mut left, mut right) = (0, slice.len());
    while left < right {
        let index = left + (right - left) / 2;
        match slice[index].cmp(&key) {
            Ordering::Equal => {
                return Some(index);
            }
            Ordering::Less => {
                left = index + 1;
            }
            Ordering::Greater => {
                right = index;
            }
        }
    }
    None
}
