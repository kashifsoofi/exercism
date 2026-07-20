using System;
using System.Collections.Generic;

public static class VariableLengthQuantity
{
    public static uint[] Encode(uint[] numbers)
    {
        var bytes = new List<uint>();
        foreach (var n in numbers)
        {
            var encoded = new List<uint> { 0x7Fu & n };
            for (var i = n >> 7; i > 0; i >>= 7)
            {
                var b = (0x7Fu & i) | 0x80u;
                encoded.Insert(0, b);
            }
            bytes.AddRange(encoded);
        }
        return bytes.ToArray();
    }

    public static uint[] Decode(uint[] bytes)
    {
        if (bytes[bytes.Length -1] > 0x7Fu)
        {
            throw new InvalidOperationException("Invalid sequence");
        }

        var number = 0u;
        var numbers = new List<uint>();
        foreach (var b in bytes)
        {
            number = (number << 7) + (0x7F & b);
            if ((0x80 & b) == 0)
            {
                numbers.Add(number);
                number = 0;
            }
        }

        return numbers.ToArray();
    }
}
