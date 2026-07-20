using System;
using System.Collections.Generic;

[Flags]
public enum Allergen
{
    Eggs = 1 << 0,
    Peanuts = 1 << 1,
    Shellfish = 1 << 2,
    Strawberries = 1 << 3,
    Tomatoes = 1 << 4,
    Chocolate = 1 << 5,
    Pollen = 1 << 6,
    Cats = 1 << 7
}

public class Allergies
{
    private readonly Allergen mask;

    public Allergies(int mask)
    {
        this.mask = (Allergen) mask;
    }

    public bool IsAllergicTo(Allergen allergen)
    {
        return mask.HasFlag(allergen);
    }

    public Allergen[] List()
    {
        var allergies = new List<Allergen>();
        foreach (Allergen allergen in Enum.GetValues(typeof(Allergen)))
        {
            if (IsAllergicTo(allergen))
            {
                allergies.Add(allergen);
            }
        }

        return allergies.ToArray();
    }
}