use std::{char::from_digit, usize};

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    minefield
        .iter()
        .enumerate()
        .map(|(y, row)| {
            row
                .char_indices()
                .map(|(x, ch)| {
                    match ch {
                        '*' => '*',
                        _ => match count_mines(minefield, (x, y)) {
                            0 => ' ',
                            n => from_digit(n, 10).unwrap(),
                        },
                    }
                })
                .collect()
        })
        .collect()
}

fn count_mines(minefield: &[&str], square: (usize, usize)) -> u32 {
    minefield
        .iter()
        .take(square.1 + 2)
        .skip((square.1 as i32 - 1).max(0) as usize)
        .flat_map(|row| {
            row.chars()
                .take(square.0 + 2)
                .skip((square.0 as i32 - 1).max(0) as usize)
        })
        .filter(|&ch| ch == '*')
        .count() as u32
}
