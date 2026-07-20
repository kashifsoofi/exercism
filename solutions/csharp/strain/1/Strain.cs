using System;
using System.Collections.Generic;

public static class Strain
{
    public static IEnumerable<T> Keep<T>(this IEnumerable<T> collection, Func<T, bool> predicate)
    {
        var keepList = new List<T>();
        foreach (var item in collection)
        {
            if (predicate(item))
            {
                keepList.Add(item);
            }
        }
        return keepList;
    }

    public static IEnumerable<T> Discard<T>(this IEnumerable<T> collection, Func<T, bool> predicate)
    {
        var discardList = new List<T>();
        foreach (var item in collection)
        {
            if (!predicate(item))
            {
                discardList.Add(item);
            }
        }
        return discardList;
    }
}