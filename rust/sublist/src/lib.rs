use std::cmp::Ordering;

#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> Comparison {
    match first_list.len().cmp(&second_list.len()) {
        Ordering::Equal if is_sublist(first_list, second_list) => Comparison::Equal,
        Ordering::Less if is_sublist(first_list, second_list) => Comparison::Sublist,
        Ordering::Greater if is_sublist(second_list, first_list) => Comparison::Superlist,
        _ => Comparison::Unequal,
    }
}

pub fn is_sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> bool {
    first_list.is_empty()
        || second_list
            .windows(first_list.len())
            .any(|w| w == first_list)
}
