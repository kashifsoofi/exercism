pub fn verse(n: u32) -> String {
    let verse1_bottles = match n {
        0 => String::from("No more bottles"),
        1 => String::from("1 bottle"),
        _ => format!("{} bottles", n)
    };

    let verse2 = match n {
        0 => String::from("Go to the store and buy some more,"),
        1 => String::from("Take it down and pass it around,"),
        _ => String::from("Take one down and pass it around,")
    };

    let verse2_bottles = match n {
        0 => String::from("99 bottles"),
        1 => String::from("no more bottles"),
        2 => String::from("1 bottle"),
        _ => format!("{} bottles", n - 1)
    };

    format!(
        "{} of beer on the wall, {} of beer.\n{} {} of beer on the wall.\n",
        verse1_bottles,
        verse1_bottles.to_lowercase(),
        verse2,
        verse2_bottles
    )
}

pub fn sing(start: u32, end: u32) -> String {    
    let mut v: Vec<String> = Vec::new();

    for i in (end..=start).rev() {
        v.push(verse(i));
    }
    
    v.join("\n")
}
