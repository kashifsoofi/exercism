using System;
using System.Collections.Generic;
using System.Linq;

public static class ArmstrongNumbers
{
    public static bool IsArmstrongNumber(int number)
    {
        var digits = new List<int>();
        for (var n = number; n > 0; n /= 10)
        {
            digits.Add(n % 10);
        }

        var p = digits.Count;
        var sum = digits.Select(d => Math.Pow(d, p)).Sum();
        return number == sum;
    }
}