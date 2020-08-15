using System;
using System.Collections.Generic;

public static class MatchingBrackets
{
    private static Dictionary<char, char> BracketPairs = new Dictionary<char, char>
    {
        [']'] = '[',
        [')'] = '(',
        ['}'] = '{',
    };

    public static bool IsPaired(string input)
    {
        var brackets = new Stack<char>();

        foreach (var ch in input)
        {
            if (BracketPairs.ContainsValue(ch))
            {
                brackets.Push(ch);
            }
            else if (BracketPairs.ContainsKey(ch))
            {
                if (brackets.Count == 0)
                {
                    return false;
                }

                var openningBracket = brackets.Pop();
                if (openningBracket != BracketPairs[ch])
                {
                    return false;
                }
            }
        }

        return brackets.Count == 0;
    }
}
