using System;
using System.Collections.Generic;

public static class PrimeFactors
{
    public static long[] Factors(long number)
    {
        var factors = new List<long>();
        for (var factor = 2L; number > 1; factor++)
        {
            for (; number % factor == 0; number = number / factor)
            {
                factors.Add(factor);
            }
        }
        return factors.ToArray();
    }
}