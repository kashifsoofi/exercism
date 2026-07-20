using System;
using System.Linq;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Threading.Tasks;

public static class ParallelLetterFrequency
{
    public static Dictionary<char, int> Calculate(IEnumerable<string> texts)
    {
        var charCounts = new ConcurrentDictionary<char, int>();
        Parallel.ForEach(texts, text =>
        {
            foreach (var c in text.ToLower())
            {
                if (char.IsLetter(c))
                {
                    charCounts.AddOrUpdate(c, 1, (_, count) => count + 1);
                }
            }
        });
        return charCounts.ToDictionary(x => x.Key, x => x.Value);
    }
}