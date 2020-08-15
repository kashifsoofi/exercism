using System;

public static class Luhn
{
    public static bool IsValid(string number)
    {
        number = number.Replace(" ", "");
        if (number.Length < 2)
        {
            return false;
        }

        var sum = 0;
        var @double = number.Length % 2 == 0;

        foreach (var ch in number)
        {
            if (!char.IsDigit(ch))
            {
                return false;
            }

            var digit = (int) char.GetNumericValue(ch);
            if (@double)
            {
                digit *= 2;
                if (digit > 9)
                {
                    digit -= 9;
                }                
            }
            @double = !@double;

            sum += digit;
        }

        return sum % 10 == 0;
    }
}