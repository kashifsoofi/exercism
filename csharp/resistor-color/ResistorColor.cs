using System;
using System.Collections.Generic;
using System.Linq;

public enum Color : int
{
    black = 0,
    brown = 1,
    red = 2,
    orange = 3,
    yellow = 4,
    green = 5,
    blue = 6,
    violet = 7,
    grey = 8,
    white = 9,
}

public static class ResistorColor
{
    private static Dictionary<string, int> _colors = Enum.GetValues(typeof(Color))
        .Cast<Color>()
        .ToDictionary(t => t.ToString(), t => (int)t);

    public static int ColorCode(string color)
    {
        if (_colors.TryGetValue(color.ToLower(), out var code))
        {
            return code;
        }
        
        throw new ArgumentException("Color not supported");
    }

    public static string[] Colors() => _colors.Keys.ToArray();
}