pub fn abbreviate(phrase: &str) -> String {
    let mut acronym = String::new();
    let words = phrase.split(|c: char| !c.is_alphabetic() && c != '\'');
    for word in words {
        let mut chars = word.chars();
        if let Some(ch) = chars.next() {
            acronym.extend(ch.to_uppercase());

            // handle camelCase
            let mut prev_was_lowercase = false;
            for c in chars {
                if c.is_uppercase() && prev_was_lowercase {
                    acronym.push(c);
                }
                prev_was_lowercase = c.is_lowercase();
            }
        }
    }
    acronym
}
