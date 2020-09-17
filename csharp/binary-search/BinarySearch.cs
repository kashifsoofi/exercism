using System;

public static class BinarySearch
{
    public static int Find(int[] input, int value)
    {
        int left = 0, right = input.Length - 1;
        while (left <= right)
        {
            var mid = left + (right - left) / 2;
            if (input[mid] == value)
            {
                return mid;
            }

            if (input[mid] < value)
            {
                left = mid + 1;
            }
            else
            {
                right = mid - 1;
            }
        }
        return -1;
    }
}