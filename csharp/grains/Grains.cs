using System;

public static class Grains
{
    public static ulong Square(int n)
    {
        if (n < 1 || n > 64)
        {
            throw new ArgumentOutOfRangeException("Number should be between 1 and 64");
        }

        return 1UL << (n - 1);
    }

    public static ulong Total()
    {
        var total = 0UL;
        for (int i = 1; i <= 64; i++)
        {
            total += Square(i);
        }
        return total;
    }
}