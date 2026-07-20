using System;
using System.Linq;

public static class Bob
{
    public static string Response(string statement)
    {
        var isSilence = string.IsNullOrWhiteSpace(statement);
        if (isSilence)
        {
            return "Fine. Be that way!";
        }

        var hasLetters = statement.Any(Char.IsLetter);
        var isYell = hasLetters && statement.ToUpperInvariant() == statement;
        var isQuestion = statement.TrimEnd().EndsWith("?");

        if (isYell && isQuestion)
        {
            return "Calm down, I know what I'm doing!";
        }

        if (isYell)
        {
            return "Whoa, chill out!";
        }

        if (isQuestion)
        {
            return "Sure.";
        }

        return "Whatever.";
    }
}