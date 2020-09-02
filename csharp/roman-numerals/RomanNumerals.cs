using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

public static class RomanNumeralExtension
{
    private static Dictionary<int, string> _valueSymbols = new Dictionary<int, string>
        {
            { 1, "I" },
            { 2, "II" },
            { 3, "III" },
            { 4, "IV" },
            { 5, "V" },
            { 9, "IX" },
            { 10, "X" },
            { 40, "XL" },
            { 50, "L" },
            { 90, "XC" },
            { 100, "C" },
            { 400, "CD" },
            { 500, "D" },
            { 900, "CM" },
            { 1000, "M" }
        };

    public static string ToRoman(this int value)
    {
        if (value < 1 || value > 3999)
        {
                throw new ArgumentOutOfRangeException("Value must be between 1 and 3999.");
        }

        var valueToConvert = value;

        StringBuilder romanNumerals = new StringBuilder();
        var keys = _valueSymbols.Select(kv => kv.Key).OrderByDescending(x => x).ToList();
        for (int i = 0; i < keys.Count; i++)
        {
            int key = keys[i];
            while (valueToConvert >= key)
            {
                romanNumerals.Append(_valueSymbols[key]);
                valueToConvert -= key;
            }

            if (valueToConvert == 0)
                break;
        }

        return romanNumerals.ToString();
    }
}