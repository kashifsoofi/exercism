pub fn verse(n: u32) -> String {
    let mut verse = String::new();

    if n == 0 {
        verse.push_str("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n");
    } else {
        verse.push_str(&n.to_string());
        verse.push_str(" bottle");
        if n > 1 {
            verse.push_str("s");
        }
        verse.push_str(" of beer on the wall, ");
        verse.push_str(&n.to_string());
        verse.push_str(" bottle");
        if n > 1 {
            verse.push_str("s");
        }
        verse.push_str(" of beer.\n");

        let m = n - 1;
        if m == 0 {
            verse.push_str(
                "Take it down and pass it around, no more bottles of beer on the wall.\n",
            );
        } else {
            verse.push_str("Take one down and pass it around, ");
            verse.push_str(&m.to_string());
            verse.push_str(" bottle");
            if m > 1 {
                verse.push_str("s");
            }
            verse.push_str(" of beer on the wall.\n");
        }
    }

    verse
}

pub fn sing(start: u32, end: u32) -> String {
    let mut song = String::new();

    let mut n = start;
    loop {
        song.push_str(&verse(n));
        if n != end {
            song.push_str("\n");
        }

        if n == end {
            break song;
        }

        n -= 1;
    }
}
