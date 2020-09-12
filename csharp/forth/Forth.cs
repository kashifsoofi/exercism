using System;
using System.Collections.Generic;
using System.Linq;
using Sprache;

public static class Forth
{
    private static class Grammer
    {
        private static readonly Parser<char> Colon = Parse.Char(':').Token();
        private static readonly Parser<char> SemiColon = Parse.Char(';').Token();
        private static readonly Parser<char> Space = Parse.Char(' ');

        private static readonly Parser<string> Word = Parse.CharExcept(' ').Many().Text();

        public static readonly Parser<int> IntNumber = Parse.Number.Select(int.Parse);

        public static readonly Parser<IEnumerable<string>> Tokens = Word.DelimitedBy(Space);

        public static readonly Parser<(string, IEnumerable<string>)> CustomWordDefinition =
            from colon in Colon
            from tokens in Parse.CharExcept(" ;").Many().Text().DelimitedBy(Space)
            from semiColon in SemiColon
            select ( tokens.First().ToLower(), tokens.Skip(1).Where(x => !string.IsNullOrWhiteSpace(x)) );
    }

    private delegate int[] OperationAction(params int[] arguments);
    private delegate void EvaluateAction(Stack<int> values);
    private static readonly Dictionary<string, EvaluateAction> evaluateActions = new Dictionary<string, EvaluateAction>
    {
        ["+"] = (values) => Binary(values, x => new[] { x[0] + x[1] }),
        ["-"] = (values) => Binary(values, x => new[] { x[0] - x[1] }),
        ["*"] = (values) => Binary(values, x => new[] { x[0] * x[1] }),
        ["/"] = (values) => Binary(values, x => x[1] == 0 ? throw new InvalidOperationException() : new[] { x[0] / x[1] }),
        ["dup"] = (values) => Unary(values, x => new[] { x[0], x[0] }),
        ["drop"] = (values) => Unary(values, x => new int[] {}),
        ["swap"] = (values) => Binary(values, x => new int[] { x[1], x[0] }),
        ["over"] = (values) => Binary(values, x => new int[] { x[0], x[1], x[0] }),
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
        foreach (var v in result)
        {
            values.Push(v);
        }
    }

    private static void Unary(Stack<int> values, OperationAction op)
    {
        if (!values.TryPop(out var x1))
        {
            throw new InvalidOperationException();
        }

        var result = op(x1);
        foreach (var v in result)
        {
            values.Push(v);
        }
    }

    public static string Evaluate(string[] instructions)
    {
        var values = new Stack<int>();
        var userDefinedWords = new Dictionary<string, List<string>>();

        void EvaluationToken(string token)
        {
            if (evaluateActions.ContainsKey(token.ToLower()))
            {
                evaluateActions[token.ToLower()](values);
                return;
            }

            var result = Grammer.IntNumber.TryParse(token);
            if (result.WasSuccessful)
            {
                values.Push(result.Value);
            }
            else
            {
                throw new InvalidOperationException();
            }
        }

        void EvaluationTokens(List<string> tokens)
        {
            foreach (var token in tokens)
            {
                if (userDefinedWords.ContainsKey(token.ToLower()))
                {
                    var definitions = userDefinedWords[token.ToLower()];
                    EvaluationTokens(definitions);
                    continue;
                }

                EvaluationToken(token);
            }
        }

        foreach (var instruction in instructions)
        {
            var userDefinedWordResult = Grammer.CustomWordDefinition.TryParse(instruction);
            if (userDefinedWordResult.WasSuccessful)
            {
                (string udfWord, IEnumerable<string> udfTokens) = userDefinedWordResult.Value;

                var result = Grammer.IntNumber.TryParse(udfWord);
                if (result.WasSuccessful)
                {
                    throw new InvalidOperationException();
                }

                if (!userDefinedWords.ContainsKey(udfWord) || userDefinedWords[udfWord] == null)
                {
                    userDefinedWords[udfWord] = new List<string>();
                }
                foreach (var udfToken in udfTokens)
                {
                    if (userDefinedWords.ContainsKey(udfToken))
                    {
                        userDefinedWords[udfWord].AddRange(userDefinedWords[udfToken]);
                        userDefinedWords[udfToken] = new List<string>();
                    }
                    else
                    {
                        userDefinedWords[udfWord].Add(udfToken);
                    }
                }
                continue;
            }

            var tokens = Grammer.Tokens.Parse(instruction).ToList();
            EvaluationTokens(tokens);
        }

        return string.Join(" ", values.ToArray().Reverse());
    }
}
