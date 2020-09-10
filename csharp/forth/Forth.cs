using System;
using System.Collections.Generic;
using System.Linq;
using Sprache;

public static class Forth
{
    internal static readonly Parser<char> Delimiter = Parse.Char(' ');
    internal static readonly Parser<string> Word = Parse.CharExcept(' ').Many().Text();
    internal static readonly Parser<IEnumerable<string>> Words = Word.DelimitedBy(Delimiter);
    internal static readonly Parser<int> IntNumber = Parse.Number.Select(int.Parse);

    internal static readonly Parser<char> ArithmeticOperator = Parse.Chars(new char[]{'+', '-', '*', '/'});

    internal delegate int OperationAction(params int[] arguments);

    internal static readonly Dictionary<string, OperationAction> actions = new Dictionary<string, OperationAction>
    {
        ["+"] = (args) =>  { return args[0] + args[1]; },
        ["-"] = (args) =>  { return args[0] - args[1]; },
        ["*"] = (args) =>  { return args[0] * args[1]; },
        ["/"] = (args) =>
                    {
                        if (args[1] == 0)
                        {
                            throw new InvalidOperationException();
                        }
                        return args[0] / args[1];
                    },
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
                var op = ArithmeticOperator.TryParse(word);
                if (op.WasSuccessful)
                {
                    Binary(values, actions[op.Value.ToString()]);
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
