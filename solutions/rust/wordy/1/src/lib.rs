#[derive(Debug)]
enum QuestionError {
    InvalidToken,
    NotAQuestion,
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

    Raised(String),
    To(String),
    The(String),
    Power(String),

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

    fn is_raised(&self) -> bool {
        match self {
            Token::Raised(..) => true,
            _ => false,
        }
    }

    fn is_to(&self) -> bool {
        match self {
            Token::To(..) => true,
            _ => false,
        }
    }

    fn is_the(&self) -> bool {
        match self {
            Token::The(..) => true,
            _ => false,
        }
    }

    fn is_power(&self) -> bool {
        match self {
            Token::Power(..) => true,
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
            "raised" => Token::Raised(literal),
            "to" => Token::To(literal),
            "the" => Token::The(literal),
            "power" => Token::Power(literal),
            _ => {
                let num = literal
                    .chars()
                    .filter(|x| match x {
                        'a'..='z' => false,
                        _ => true,
                    })
                    .collect::<String>();

                match num.parse::<i32>() {
                    Ok(n) => Token::Number(n),
                    _ => Token::Invalid(literal),
                }
            }
        }
    }

    fn unscan(&mut self) {
        self.read_index -= 1;
    }
}

#[derive(Debug, Clone, PartialEq)]
enum Operator {
    Plus,
    Minus,
    Times,
    Divide,
    Exponential,
}

struct Parser {
    scanner: Scanner,
}

impl Parser {
    fn new(command: &str) -> Self {
        Parser { scanner: Scanner::new(command) }
    }

    fn scan(&mut self) -> Token {
        self.scanner.scan()
    }

    fn unscan(&mut self) {
        self.scanner.unscan()
    }

    fn parse(&mut self) -> Result<Expression, QuestionError> {
        let mut token = self.scan();
        if !token.is_what() {
            return Err(QuestionError::InvalidToken)
        }

        token = self.scan();
        if !token.is_is() {
            return Err(QuestionError::InvalidToken)
        }

        let mut expression = Expression { value: 0, op_terms: Vec::default() };

        token = self.scan();
        match token {
            Token::Number(value) => expression.value = value,
            _ => return Err(QuestionError::InvalidToken),
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
                Token::Raised(_) => operator = Operator::Exponential,
                _ => return Err(QuestionError::InvalidToken)
            }

            if token.is_times() || token.is_divide() {
                token = self.scan();
                if !token.is_by() {
                    return Err(QuestionError::InvalidToken)
                }
            }

            if token.is_raised() {
                token = self.scan();
                if !token.is_to() {
                    return Err(QuestionError::InvalidToken)
                }

                token = self.scan();
                if !token.is_the() {
                    return Err(QuestionError::InvalidToken)
                }
            }

            let value: i32;
            token = self.scan();
            match token {
                Token::Number(n) => value = n,
                _ => return Err(QuestionError::InvalidToken),
            }
            
            if operator == Operator::Exponential {
                token = self.scan();
                if !token.is_power() {
                    return Err(QuestionError::InvalidToken)
                }
            }
            expression.op_terms.push(OperatorTerm { operator: operator, value: value });
        }
        
        token = self.scan();
        if !token.is_question_mark() {
            return Err(QuestionError::NotAQuestion)
        }
        
        Ok(expression)
    }
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
    expression: Expression,
}

impl WordProblem {
    fn new(expression: &Expression) -> Self {
        WordProblem { expression: expression.clone() }
    }


    fn eval(&self) -> Option<i32> {
        let mut result = self.expression.value;

        for op_term in self.expression.op_terms.iter() {
            match op_term.operator {
                Operator::Plus => result += op_term.value,
                Operator::Minus => result -= op_term.value,
                Operator::Times => result *= op_term.value,
                Operator::Divide => result /= op_term.value,
                Operator::Exponential => result = result.pow(op_term.value as u32),
            }
        }

        Some(result)
    }
}

pub fn answer(command: &str) -> Option<i32> {
    let mut parser = Parser::new(command);
    match parser.parse() {
        Ok(expr) => {
            let word_problem = WordProblem::new(&expr);
            word_problem.eval()
        },
        Err(_) => None,
    }
}
