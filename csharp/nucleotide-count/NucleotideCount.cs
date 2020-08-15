using System;
using System.Collections.Generic;

public static class NucleotideCount
{
    public static IDictionary<char, int> Count(string sequence)
    {
        var nucleotideCounts = new Dictionary<char, int>
        {
            ['A'] = 0,
            ['C'] = 0,
            ['G'] = 0,
            ['T'] = 0
        };

        foreach (var ch in sequence)
        {
            if (!nucleotideCounts.ContainsKey(ch))
            {
                throw new ArgumentException("Invalid nucleotide");
            }

            nucleotideCounts[ch]++;
        }

        return nucleotideCounts;
    }
}