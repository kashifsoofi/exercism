use std::usize;

#[derive(Debug)]
enum QuestionError {
    InvalidQuestion
}

#[derive(Debug, PartialEq)]
enum Token {
    Invalid(String),
    Eof,

    What(String),
    Is(String),

    Number(i32),

    Plus(String),
    Minus(String),
    Times(String),
    Divide(String),
    By(String),

    QuestionMark(String),
}

impl Token {
    fn is_what(&self) -> bool {
        match self {
            Token::What(..) => true,
            _ => false,
        }
    }

    fn is_is(&self) -> bool {
        match self {
            Token::Is(..) => true,
            _ => false,
        }
    }

    fn is_number(&self) -> bool {
        match self {
            Token::Number(..) => true,
            _ => false,
        }
    }

    fn is_operator(&self) -> bool {
        self.is_plus() || self.is_minus() || self.is_times() || self.is_divide()
    }

    fn is_plus(&self) -> bool {
        match self {
            Token::Plus(..) => true,
            _ => false,
        }
    }

    fn is_minus(&self) -> bool {
        match self {
            Token::Minus(..) => true,
            _ => false,
        }
    }

    fn is_times(&self) -> bool {
        match self {
            Token::Times(..) => true,
            _ => false,
        }
    }

    fn is_divide(&self) -> bool {
        match self {
            Token::Divide(..) => true,
            _ => false,
        }
    }

    fn is_by(&self) -> bool {
        match self {
            Token::By(..) => true,
            _ => false,
        }
    }

    fn is_question_mark(&self) -> bool {
        match self {
            Token::QuestionMark(..) => true,
            _ => false,
        }
    }
}

#[derive(Debug, Clone)]
struct Scanner {
    literals: Vec<String>,
    read_index: usize,
}

impl Scanner {
    fn new(command: &str) -> Self {
        let mut tokens: Vec<_> = command.split_whitespace().map(|s| s.to_string()).collect();

        let last = tokens.last();
        if last.is_some() && last.unwrap().ends_with("?") {
            let mut last = tokens.pop().unwrap();
            let question_mark = last.pop().unwrap();
            tokens.push(last);
            tokens.push(question_mark.to_string());
        }

        Scanner { literals: tokens, read_index: 0 }
    }

    fn scan(&mut self) -> Token {
        if self.read_index == self.literals.len() {
            return Token::Eof;
        }

        let literal: String = self.literals[self.read_index].clone();
        self.read_index += 1;

        match literal.to_lowercase().as_str() {
            "what" => Token::What(literal),
            "is" => Token::Is(literal),
            "plus" => Token::Plus(literal),
            "minus" => Token::Minus(literal),
            "multiplied" => Token::Times(literal),
            "divided" => Token::Divide(literal),
            "by" => Token::By(literal),
            "?" => Token::QuestionMark(literal),
            _ => {
                match literal.parse::<i32>() {
                    Ok(n) => Token::Number(n),
                    _ => Token::Invalid(literal)
                }
            }
        }
    }

    fn unscan(&mut self) {
        self.read_index -= 1;
    }
}

pub struct WordProblem {
    scanner: Scanner,
    tokens: Vec<Token>,
}

impl WordProblem {
    fn new(command: &str) -> Self {
        WordProblem { scanner: Scanner::new(command), tokens: Vec::default() }
    }

    fn scan(&mut self) -> Token {
        self.scanner.scan()
    }

    fn unscan(&mut self) {
        self.scanner.unscan()
    }

    fn parse(&mut self) -> Option<QuestionError> {
        let mut token = self.scan();
        if !token.is_what() {
            return Some(QuestionError::InvalidQuestion)
        }

        token = self.scan();
        if !token.is_is() {
            return Some(QuestionError::InvalidQuestion)
        }

        token = self.scan();
        if !token.is_number() {
            return Some(QuestionError::InvalidQuestion)
        }
        self.tokens.push(token);

        loop {
            token = self.scan();
            if token.is_question_mark() {
                self.unscan();
                break;
            }

            if !token.is_operator() {
                return Some(QuestionError::InvalidQuestion)
            }

            if token.is_times() || token.is_divide() {
                token = self.scan();
                if !token.is_by() {
                    return Some(QuestionError::InvalidQuestion)
                }
            }
            self.tokens.push(token);

            token = self.scan();
            if !token.is_number() {
                return Some(QuestionError::InvalidQuestion)
            }
            self.tokens.push(token);
        }

        token = self.scan();
        if !token.is_question_mark() {
            return Some(QuestionError::InvalidQuestion)
        }

        None
    }

    fn eval(&self) -> Option<i32> {
        None
    }
}

pub fn answer(command: &str) -> Option<i32> {
    let mut word_problem = WordProblem::new(command);
    match word_problem.parse() {
        None => word_problem.eval(),
        _ => None
    }
}
