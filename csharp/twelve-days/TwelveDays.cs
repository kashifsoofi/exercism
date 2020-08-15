using System;
using System.Collections.Generic;
using System.Text;

public static class TwelveDays
{
    private static Dictionary<int, string> dayName = new Dictionary<int, string>
    {
        [1] = "first",
        [2] = "second",
        [3] = "third",
        [4] = "fourth",
        [5] = "fifth",
        [6] = "sixth",
        [7] = "seventh",
        [8] = "eighth",
        [9] = "ninth",
        [10] = "tenth",
        [11] = "eleventh",
        [12] = "twelfth",
    };

    private static List<string> gifts = new List<string>
    {
        "a Partridge in a Pear Tree",
        "two Turtle Doves",
        "three French Hens",
        "four Calling Birds",
        "five Gold Rings",
        "six Geese-a-Laying",
        "seven Swans-a-Swimming",
        "eight Maids-a-Milking",
        "nine Ladies Dancing",
        "ten Lords-a-Leaping",
        "eleven Pipers Piping",
        "twelve Drummers Drumming",
    };

    public static string Recite(int verseNumber)
    {
        var giftsForDay = new StringBuilder();
        for (var i = verseNumber - 1; i > 0; i--)
        {
            giftsForDay.Append($"{gifts[i]}, ");
        }
        if (giftsForDay.Length > 0)
        {
            giftsForDay.Append("and ");
        }
        giftsForDay.Append(gifts[0]);

        return $"On the {dayName[verseNumber]} day of Christmas my true love gave to me: {giftsForDay}.";
    }

    public static string Recite(int startVerse, int endVerse)
    {
        var verses = new StringBuilder();
        for (var i = startVerse; i <= endVerse; i++)
        {
            verses.Append(Recite(i));
            if (i < endVerse)
            {
                verses.Append("\n");
            }
        }
        return verses.ToString();
    }
}