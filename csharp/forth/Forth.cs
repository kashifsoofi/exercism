using System;
using System.Collections.Generic;
using System.Linq;
using Sprache;

public static class Forth
{
    private static readonly Parser<char> Delimiter = Parse.Char(' ');
    private static readonly Parser<string> Word = Parse.CharExcept(' ').Many().Text();
    private static readonly Parser<IEnumerable<string>> Words = Word.DelimitedBy(Delimiter);
    private static readonly Parser<int> IntNumber = Parse.Number.Select(int.Parse);

    private static readonly Parser<string> Plus = Parse.IgnoreCase("+").Text();
    private static readonly Parser<string> Minus = Parse.IgnoreCase("-").Text();
    private static readonly Parser<string> Multiply = Parse.IgnoreCase("*").Text();
    private static readonly Parser<string> Divide = Parse.IgnoreCase("/").Text();
    private static readonly Parser<string> Duplicate = Parse.IgnoreCase("dup").Text();
    private static readonly Parser<string> Drop = Parse.IgnoreCase("drop").Text();
    private static readonly Parser<string> Swap = Parse.IgnoreCase("swap").Text();
    private static readonly Parser<string> Over = Parse.IgnoreCase("over").Text();

    private static readonly Parser<string> Operator = 
        Plus.Or(Minus).Or(Multiply).Or(Divide).Or(Duplicate).Or(Drop).Or(Swap).Or(Over);

    private delegate int OperationAction(params int[] arguments);
    private delegate void EvaluateAction(Stack<int> values);
    private static readonly Dictionary<string, EvaluateAction> evaluateActions = new Dictionary<string, EvaluateAction>
    {
        ["+"] = (values) => Binary(values, x => x[0] + x[1]),
        ["-"] = (values) => Binary(values, x => x[0] - x[1]),
        ["*"] = (values) => Binary(values, x => x[0] * x[1]),
        ["/"] = (values) => Binary(values, x => x[1] == 0 ? throw new InvalidOperationException() : x[0] / x[1]),
    };

    private static void Binary(Stack<int> values, OperationAction op)
    {
        if (!values.TryPop(out var x2))
        {
            throw new InvalidOperationException();
        }
        if (!values.TryPop(out var x1))
        {
            throw new InvalidOperationException();
        }

        var result = op(x1, x2);
        values.Push(result);
    }

    public static string Evaluate(string[] instructions)
    {
        var values = new Stack<int>();

        foreach (var instruction in instructions)
        {
            var words = Words.Parse(instruction);
            foreach (var word in words)
            {
                var op = Operator.TryParse(word);
                if (op.WasSuccessful)
                {
                    evaluateActions[op.Value.ToString()](values);
                }
                else 
                {
                var result = IntNumber.TryParse(word);
                if (result.WasSuccessful)
                {
                    values.Push(result.Value);
                }
                }
            }
        }

        return string.Join(" ", values.ToArray().Reverse());
    }
}
