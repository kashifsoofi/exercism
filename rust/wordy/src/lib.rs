#[derive(Debug)]
pub enum QuestionError {
    InvalidToken,
    NotAQuestion,
}

#[derive(Debug, PartialEq)]
pub enum Token {
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
    pub fn is_what(&self) -> bool {
        match self {
            Token::What(..) => true,
            _ => false,
        }
    }

    pub fn is_is(&self) -> bool {
        match self {
            Token::Is(..) => true,
            _ => false,
        }
    }

    pub fn is_times(&self) -> bool {
        match self {
            Token::Times(..) => true,
            _ => false,
        }
    }

    pub fn is_divide(&self) -> bool {
        match self {
            Token::Divide(..) => true,
            _ => false,
        }
    }

    pub fn is_by(&self) -> bool {
        match self {
            Token::By(..) => true,
            _ => false,
        }
    }

    pub fn is_question_mark(&self) -> bool {
        match self {
            Token::QuestionMark(..) => true,
            _ => false,
        }
    }
}

#[derive(Debug, Clone)]
pub struct Scanner {
    literals: Vec<String>,
    read_index: usize,
}

impl Scanner {
    pub fn new(command: &str) -> Self {
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

    pub fn scan(&mut self) -> Token {
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

    pub fn unscan(&mut self) {
        self.read_index -= 1;
    }
}

#[derive(Debug, Clone)]
enum Operator {
    Plus,
    Minus,
    Times,
    Divide,
}

#[derive(Debug, Clone)]
struct OperatorTerm {
    operator: Operator,
    value: i32,
}

#[derive(Debug, Clone)]
struct Expression {
    value: i32,
    op_terms: Vec<OperatorTerm>,
}

pub struct WordProblem {
    scanner: Scanner,
    expression: Option<Expression>,
}

impl WordProblem {
    fn new(command: &str) -> Self {
        WordProblem { scanner: Scanner::new(command), expression: None }
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
            return Some(QuestionError::InvalidToken)
        }

        token = self.scan();
        if !token.is_is() {
            return Some(QuestionError::InvalidToken)
        }

        let mut expression = Expression { value: 0, op_terms: Vec::default() };

        token = self.scan();
        match token {
            Token::Number(value) => expression.value = value,
            _ => return Some(QuestionError::InvalidToken),
        }

        loop {
            token = self.scan();
            if token.is_question_mark() {
                self.unscan();
                break;
            }

            let operator;
            match token {
                Token::Plus(_) => operator = Operator::Plus,
                Token::Minus(_) => operator = Operator::Minus,
                Token::Times(_) => operator = Operator::Times,
                Token::Divide(_) => operator = Operator::Divide,
                _ => return Some(QuestionError::InvalidToken)
            }

            if token.is_times() || token.is_divide() {
                token = self.scan();
                if !token.is_by() {
                    return Some(QuestionError::InvalidToken)
                }
            }

            let value: i32;
            token = self.scan();
            match token {
                Token::Number(n) => value = n,
                _ => return Some(QuestionError::InvalidToken),
            }

            expression.op_terms.push(OperatorTerm { operator: operator, value: value })
        }

        token = self.scan();
        if !token.is_question_mark() {
            return Some(QuestionError::NotAQuestion)
        }

        self.expression = Some(expression);

        None
    }

    fn eval(&self) -> Option<i32> {
        match self.expression.clone() {
            None => None,
            Some(expr) => {
                let mut result = expr.value;

                for op_term in expr.op_terms.iter() {
                    match op_term.operator {
                        Operator::Plus => result += op_term.value,
                        Operator::Minus => result -= op_term.value,
                        Operator::Times => result *= op_term.value,
                        Operator::Divide => result /= op_term.value,
                    }
                }

                Some(result)
            }
        }
    }
}

pub fn answer(command: &str) -> Option<i32> {
    let mut word_problem = WordProblem::new(command);
    match word_problem.parse() {
        None => word_problem.eval(),
        _ => None
    }
}
