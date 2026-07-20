using System;
using System.Linq;
using System.Linq.Expressions;
using Sprache;

public static class Wordy
{
    private static class Grammer
    {
        private static readonly Parser<string> WhatIs = Parse.String("What is").Text().Token();
        private static readonly Parser<char> QuestionMark = Parse.Char('?').Token();
        private static readonly Parser<Expression> Constant =
            from op in Parse.Optional(Parse.Char('-').Token())
            from n in Parse.Number.Token()
            select Expression.Constant(int.Parse(n) * (op.IsDefined ? -1 : 1));
        private static readonly Parser<ExpressionType> Operator =
            Parse.String("plus").Return(ExpressionType.Add)
                .Or(Parse.String("minus").Return(ExpressionType.Subtract))
                .Or(Parse.String("multiplied by").Return(ExpressionType.Multiply))
                .Or(Parse.String("divided by").Return(ExpressionType.Divide));
        private static readonly Parser<Expression> Operation =
            Parse.ChainOperator(Operator, Constant, Expression.MakeBinary);

        public static readonly Parser<Expression> Question =
            from whatIs in WhatIs
            from expression in Operation.Or(Constant)
            from questionMark in QuestionMark.End()
            select expression;
    }

    public static int Answer(string question)
    {
        try
        {
            var operation = Grammer.Question.Parse(question);
            var func = Expression.Lambda<Func<int>>(operation).Compile();
            return func();
        }
        catch (Exception e)
        {
            throw new ArgumentException("Invalid input", e);
        }
    }
}