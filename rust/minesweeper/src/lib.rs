use std::{char::from_digit, usize};

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let mut result: Vec<String> = Vec::with_capacity(minefield.len());
    for (x, row) in minefield.iter().enumerate() {
        let s = row
            .chars()
            .enumerate()
            .map(|(y, ch)| match ch {
                ' ' => {
                    match get_adjacent_mine_count(minefield, x, y) {
                        0 => ' ',
                        c => from_digit(c, 10).unwrap()
                    }
                },
                _ => ch,
            })
            .collect();

        result.push(s);
    }

    result
}

fn get_adjacent_mine_count(minefield: &[&str], x: usize, y: usize) -> u32 {
    let mut mine_count = 0;
    for i in x.saturating_sub(1)..=(x+1).min(minefield.len() - 1) {
        for j in y.saturating_sub(1)..=(y+1).min(minefield[x].len() - 1) {
            match minefield.get(i).map(|&s| s.get(j..j+1)) {
                Some(Some("*")) => {
                    mine_count += 1;
                },
                _ => continue,
            }
        }
    }
    mine_count
}