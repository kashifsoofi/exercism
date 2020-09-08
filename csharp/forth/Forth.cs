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

    public static string Evaluate(string[] instructions)
    {
        var values = new Stack<int>();

        foreach (var instruction in instructions)
        {
            var words = Words.Parse(instruction);
            foreach (var word in words)
            {
                var result = IntNumber.TryParse(word);
                if (result.WasSuccessful)
                {
                    values.Push(result.Value);
                }
            }
        }

        return string.Join(" ", values.ToArray().Reverse());
    }
}
