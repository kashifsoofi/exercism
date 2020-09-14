using System;
using System.Collections;
using System.Collections.Generic;

public static class FlattenArray
{
    public static IEnumerable Flatten(IEnumerable input)
    {
        void Flatten(IEnumerable input, List<object> output)
        {
            foreach (var item in input)
            {
                var enumerable = item as IEnumerable;
                if (enumerable != null)
                {
                    Flatten(enumerable, output);
                }
                else if (item != null)
                {
                    output.Add(item);
                }
            }
        }

        var flattened = new List<object>();
        Flatten(input, flattened);
        return flattened;
    }
}