using System;

public class Clock
{
    private readonly int hours;
    private readonly int minutes;

    public Clock(int hours, int minutes)
        : this(minutes + (hours * 60))
    {
    }

    private Clock(int totalMinutes)
    {
        (this.hours, this.minutes) = GetHoursAndMinutes(totalMinutes);
    }

    public Clock Add(int minutesToAdd)
    {
        return new Clock(minutes + (hours * 60) + minutesToAdd);
    }

    public Clock Subtract(int minutesToSubtract)
    {
        return new Clock(minutes + (hours * 60) - minutesToSubtract);
    }

    public override string ToString()
    {
        return $"{hours:00}:{minutes:00}";
    }

    public override bool Equals(object obj)
    {
        var other = obj as Clock;
        if (other == null)
        {
            return false;
        }

        return hours == other.hours && minutes == other.minutes;
    }

    private (int hours, int minutes) GetHoursAndMinutes(int totalMinutes)
    {
        var minutes = (totalMinutes % 60 + (totalMinutes < 0 ? 60 : 0)) % 60;

        var totalHours = (totalMinutes - minutes) / 60;
        var hours = (totalHours % 24 + (totalHours < 0 ? 24 : 0)) % 24;
        return (hours, minutes);
    }
}
