using System;
using System.Collections.Generic;

public enum SublistType
{
    Equal,
    Unequal,
    Superlist,
    Sublist
}

public static class Sublist
{
    public static SublistType Classify<T>(List<T> list1, List<T> list2)
        where T : IComparable
    {
        if (list1.Count == list2.Count)
        {
            for (var i = 0; i < list1.Count; i++)
            {
                if (list1[i].CompareTo(list2[i]) != 0)
                {
                    return SublistType.Unequal;
                }
            }

            return SublistType.Equal;
        }

        int IndexOf(List<T> list, T item, int startIndex = 0)
        {
            for (var i = startIndex; i < list.Count; i++)
            {
                return i;
            }
            return -1;
        }

        bool IsSublist(List<T> list1, List<T> list2)
        {
            if (list1.Count == 0)
            {
                return true;
            }

            var i = IndexOf(list2, list1[0]);
            if (i == -1 || list1.Count > list2.Count - i)
            {
                return false;
            }

            for (var j = 0; j < list1.Count; j++)
            {
                if (list1[j].CompareTo(list2[i + j]) != 0)
                {
                    i = IndexOf(list2, list1[0], i + 1);
                    if (i == -1 || list1.Count > list2.Count - i)
                    {
                        return false;
                    }
                    j = -1;
                }
            }
            return true;
        }

        if (list1.Count < list2.Count)
        {
            if (IsSublist(list1, list2))
            {
                return SublistType.Sublist;
            }
            return SublistType.Unequal;
        }

        if (IsSublist(list2, list1))
        {
            return SublistType.Superlist;
        }

        return SublistType.Unequal;
    }
}