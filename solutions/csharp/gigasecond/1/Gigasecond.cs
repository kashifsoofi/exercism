using System;

public static class Gigasecond
{
    public static DateTime Add(DateTime moment)
    {
        var gigaSecond = TimeSpan.FromSeconds(1000000000);
        return moment.Add(gigaSecond);
    }
}