using System;
using System.Collections.Generic;
using System.Linq;

public class Anagram
{
    private readonly string baseWord;

    public Anagram(string baseWord)
    {
        this.baseWord = baseWord;
    }

    public string[] FindAnagrams(string[] potentialMatches)
    {
        var anagrams = new List<string>();

        var sortedBaseWord = string.Concat(baseWord.ToLower().OrderBy(c => c));
        foreach (var candidate in potentialMatches)
        {
            if (string.Compare(baseWord, candidate, StringComparison.OrdinalIgnoreCase) != 0 &&
                string.Compare(string.Concat(candidate.ToLower().OrderBy(c => c)), sortedBaseWord, StringComparison.OrdinalIgnoreCase) == 0)
            {
                anagrams.Add(candidate);
            }
        }

        return anagrams.ToArray();
    }
}