public class Clock
{
    private readonly int timeInMinutes;

    public Clock(int hours, int minutes)
    {
        timeInMinutes = AdjustTimeInMinutes(hours * 60 + minutes);
    }

    public int Hours => timeInMinutes / 60;

    public int Minutes => timeInMinutes % 60;

    public Clock Add(int minutesToAdd) => new Clock(Hours, Minutes + minutesToAdd);

    public Clock Subtract(int minutesToSubtract) => Add(-minutesToSubtract);

    public override string ToString()
    {
        return $"{Hours:D2}:{Minutes:D2}";
    }

    public override bool Equals(object obj)
    {
        var other = obj as Clock;
        if (other == null)
        {
            return false;
        }

        return timeInMinutes == other.timeInMinutes;
    }

    private int AdjustTimeInMinutes(int minutesIn)
    {
        var maxMinutes = 24 * 60;
        return (maxMinutes + minutesIn % maxMinutes) % maxMinutes;
    }
}