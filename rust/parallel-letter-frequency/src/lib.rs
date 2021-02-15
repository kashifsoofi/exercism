use std::thread;
use std::collections::HashMap;

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let mut handles = vec![];
    let chunk_size = input.len() / worker_count + 1;
    input
        .chunks(chunk_size)
        .for_each(|chunk| {
            let data: Vec<String> = chunk.iter().map(|s| s.to_lowercase()).collect();
            let handle = thread::spawn(move || {
                let mut counts : HashMap<char, usize> = HashMap::new();

                for line in data {
                    line
                        .chars()
                        .filter(|c| c.is_alphabetic())
                        .for_each(|ch| {
                            let count = counts.entry(ch).or_insert(0);
                            *count += 1;
                        });
                }

                counts
            });

            handles.push(handle);
        });


    let mut char_counts : HashMap<char, usize> = HashMap::new();
    for handle in handles {
        for (ch, c) in handle.join().unwrap() {
            let count = char_counts.entry(ch).or_insert(0);
            *count += c;
        }
    }

    char_counts
}
