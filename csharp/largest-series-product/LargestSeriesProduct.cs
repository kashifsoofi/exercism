using System;

public static class LargestSeriesProduct
{
    public static long GetLargestProduct(string digits, int span)
    {
        if (digits.Length < span || span < 0)
        {
            throw new ArgumentException();
        }

        var numbers = new long[digits.Length];
        for (var i = 0; i < digits.Length; i++)
        {
            if (!long.TryParse(digits[i].ToString(), out numbers[i]))
            {
                throw new ArgumentException();
            }
        }

        if (span == 0)
        {
            return 1L;
        }

        var largestProduct = 0L;
        for (var i = 0; i < numbers.Length - span + 1; i++)
        {
            var product = numbers[i];
            for (var j = 1; j < span; j++)
            {
                product *= numbers[i + j];
            }

            largestProduct = product > largestProduct ? product : largestProduct;
        }
        return largestProduct;
    }
}