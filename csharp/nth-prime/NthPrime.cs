using System;
using System.Collections.Generic;
using System.Linq;

public static class NthPrime
{
    public static int Prime(int nth)
    {
        if (nth == 0)
        {
            throw new ArgumentOutOfRangeException();
        }

        if (nth == 1)
        {
            return 2;
        }

        var prime = 3;
        for (int i = 3, counter = 2; counter <= nth; i += 2)
        {
            if (isPrime(i))
            {
                prime = i;
                counter++;
            }
        }
        return prime;
    }

    private static bool isPrime(int n)
    {
        if (n > 2 && n % 2 == 0)
        {
            return false;
        }

        var limit = n / 2 + 1;
        for (var i = 3; i < limit; i++)
        {
            if (n % i == 0)
            {
                return false;
            }
        }

        return true;
    }
}