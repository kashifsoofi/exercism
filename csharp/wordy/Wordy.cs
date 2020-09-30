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
        [Lexeme("multiplied by")] TIMES = 7,
        [Lexeme("divided by")] DIVIDE = 8,
    }

    public class ExpressionParser
    {
        [Production("primary: WHAT_IS expression QUESTION_MARK")]
        public int Expression( Token<ExpressionToken> forget, int value, Token<ExpressionToken> ignore)
        {
            return value;
        }

        [Production("expression : term PLUS expression")]
        [Production("expression : term MINUS expression")]
        [Production("expression : term TIMES expression")]
        [Production("expression : term DIVIDE expression")]
        public int Binary(int left, Token<ExpressionToken> operatorToken, int right)
        {
            var result = 0;
            switch (operatorToken.TokenID)
            {
                case ExpressionToken.PLUS:
                    result = left + right;
                    break;
                case ExpressionToken.MINUS:
                    result = left - right;
                    break;
                case ExpressionToken.TIMES:
                    result = left * right;
                    break;
                case ExpressionToken.DIVIDE:
                    result = left / right;
                    break;
            }
            return result;
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
        var parser = builder.BuildParser(parserInstance, ParserType.LL_RECURSIVE_DESCENT, "primary");
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
        throw new ArgumentException();
    }
}