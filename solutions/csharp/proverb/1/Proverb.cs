using System;
using System.Collections.Generic;

public static class Proverb
{
    public static string[] Recite(string[] subjects)
    {
        var proverbs = new List<string>();

        for (var i = 0; i < subjects.Length; i++)
        {
            if ( i < subjects.Length - 1)
            {
                proverbs.Add($"For want of a {subjects[i]} the {subjects[i + 1]} was lost.");
            }
            else if (i == subjects.Length - 1)
            {
                proverbs.Add($"And all for the want of a {subjects[0]}.");
            }
        }

        return proverbs.ToArray();
    }
}