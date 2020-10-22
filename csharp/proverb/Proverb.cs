using System;
using System.Collections.Generic;

public static class Proverb
{
    public static string[] Recite(string[] subjects)
    {
        var proverbs = new string[subjects.Length];
        if (subjects.Length == 0)
        {
            return proverbs;
        }

        for (var i = 0; i < subjects.Length - 1; i++)
        {
            proverbs[i] = $"For want of a {subjects[i]} the {subjects[i + 1]} was lost.";
        }

        proverbs[proverbs.Length - 1] = $"And all for the want of a {subjects[0]}.";

        return proverbs;
    }
}