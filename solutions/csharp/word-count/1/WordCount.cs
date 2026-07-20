using System;
using System.Collections.Generic;

public static class WordCount
{
    public static IDictionary<string, int> CountWords(string phrase)
    {
        var delimeters = new [] { ' ', ',', ':', '.', '!', '@', '$', '%', '^', '&', '\n' };
        var words = phrase.ToLower().Split(delimeters, StringSplitOptions.RemoveEmptyEntries);

        var wordCounts = new Dictionary<string, int>();
        foreach (var word in words)
        {
            var key = word.Trim('\'');
            if (!wordCounts.ContainsKey(key))
            {
                wordCounts[key] = 0;
            }
            wordCounts[key]++;
        }

        return wordCounts;
    }
}