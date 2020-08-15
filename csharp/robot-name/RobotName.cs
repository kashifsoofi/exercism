using System;
using System.Collections.Generic;

public class Robot
{
    const string chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

    private static readonly HashSet<string> generatedNames = new HashSet<string>();

    public Robot()
    {
        Name = GenerateName();
    }

    public string Name { get; private set; }

    public void Reset()
    {
        Name = GenerateName();
    }

    private string GenerateName()
    {
        var name = GenerateRandomName();
        while (generatedNames.Contains(name))
        {
            name = GenerateRandomName();
        }
        
        generatedNames.Add(name);
        return name;

        string GenerateRandomName()
        {
            var random = new Random();
            var letter1 = chars[random.Next(0, chars.Length)];
            var letter2 = chars[random.Next(0, chars.Length)];
            var number = random.Next(1, 1000);
            return $"{letter1}{letter2}{number.ToString("D3")}";
        }
    }
}