using System;
using System.Collections;
using System.Collections.Generic;

public static class FlattenArray
{
    public static IEnumerable Flatten(IEnumerable input)
    {
        var flattened = new List<object>();
        foreach (var item in input)
        {
            if (item is IEnumerable enumerable)
            {
                foreach (var nestedItem in Flatten(enumerable))
                {
                    yield return nestedItem;
                }
            }
            else if (item != null)
            {
                yield return item;
            }
        }
    }
}