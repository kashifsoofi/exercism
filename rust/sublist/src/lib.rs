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
        _ => Comparison::Unequal
    }
}

pub fn is_sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> bool {
    if first_list.len() == 0 {
        return true;
    }

    let mut i = 0;
    let mut j = 0;
    while i < first_list.len() && j < second_list.len() {
        if first_list[i] != second_list[j] {
            i = 0;
            while j < second_list.len() && second_list[j] != first_list[i] {
                j += 1;
            }
        }
        else {
            i += 1;
            j += 1;
        }
    }

    if i < first_list.len() {
        return false;
    }

    true
}
