using System;
using sly.lexer;
using sly.parser;
using sly.parser.generator;

public static class Wordy
{
    public enum ExpressionToken
    {
        [Lexeme("What is")] WHAT_IS = 1,
        [Lexeme("\\?")] QUESTION_MARK = 2,
        [Lexeme("[ \\t]+", true)] WS = 3,
        [Lexeme("(-)?[0-9]+")] INT = 4,
        [Lexeme("plus")] PLUS = 5,
        [Lexeme("minus")] MINUS = 6,
    }

    public class ExpressionParser
    {
        [Production("primary: WHAT_IS term QUESTION_MARK")]
        public int Expression( Token<ExpressionToken> forget, int value, Token<ExpressionToken> ignore)
        {
            return value;
        }

        [Production("expression : term PLUS expression")]
        public int Binary(int left, Token<ExpressionToken> operatorToken, int right)
        {
            return left + right;
        }

        [Production("expression : term")]
        public int Expression_Term(int termValue)
        {
            return termValue;
        }

        [Production("term : INT")]
        public int Expression(Token<ExpressionToken> intToken)
        {
            return intToken.IntValue;
        }
    }

    public static Parser<ExpressionToken, int> GetParser()
    {
        var parserInstance = new ExpressionParser();
        var builder = new ParserBuilder<ExpressionToken, int>();
        var parser = builder.BuildParser(parserInstance, ParserType.LL_RECURSIVE_DESCENT, "expression");
        return parser.Result;
    }

    public static int Answer(string question)
    {
        var parser = GetParser();
        var r = parser.Parse(question);
        if (!r.IsError && r.Result != null && r.Result is int)
        {
            return (int)r.Result;
        }
        return -1;
    }
}