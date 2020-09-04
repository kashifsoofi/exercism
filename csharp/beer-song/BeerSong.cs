using System.Collections.Generic;

public static class BeerSong
{
    public static string Recite(int startBottles, int takeDown)
    {
        var verses = new List<string>();
        for (var i = startBottles; i > startBottles - takeDown; i--)
        {
            var verse = i switch
            {
                0 => "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.",
                1 => "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.",
                2 => "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.",
                _ => $"{i} bottles of beer on the wall, {i} bottles of beer.\nTake one down and pass it around, {i - 1} bottles of beer on the wall."
            };
            verses.Add(verse);
        }
        return string.Join("\n\n", verses);
    }
}