using System;
using System.Collections.Generic;
using System.Linq;

public static class Isogram
{
    public static bool IsIsogram(string word)
    {
        var timmedWord = new string(word.ToLower().Where(c => char.IsLetter(c)).ToArray());
        var charCount = new Dictionary<char, int>();
        foreach(var ch in timmedWord)
        {
            if (charCount.ContainsKey(ch))
            {
                return false;
            }

            charCount.Add(ch, 1);
        }

        return true;
    }
}
