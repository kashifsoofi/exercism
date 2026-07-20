pub fn reply(message: &str) -> &str {
    let trimmed_message = message.trim_end();

    let is_empty = trimmed_message.is_empty();
    let is_question = trimmed_message.ends_with("?");
    let is_yelling = trimmed_message == trimmed_message.to_uppercase() && trimmed_message.chars().any(char::is_alphabetic);

    match (is_empty, is_question, is_yelling) {
        (true, _, _) => "Fine. Be that way!",
        (_, true, true) => "Calm down, I know what I'm doing!",
        (_, true, false) => "Sure.",
        (_, false, true) => "Whoa, chill out!",
        (_, _, _) => "Whatever."
    }
}
