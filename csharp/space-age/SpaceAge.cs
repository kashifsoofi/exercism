using System;

public class SpaceAge
{
    private readonly double ageOnEarth;

    public SpaceAge(int seconds)
    {
        ageOnEarth = seconds / 31557600D;
    }

    public double OnEarth() => ageOnEarth;

    public double OnMercury() => ageOnEarth / 0.2408467;

    public double OnVenus() => ageOnEarth / 0.61519726;

    public double OnMars() => ageOnEarth / 1.8808158;

    public double OnJupiter() => ageOnEarth / 11.862615;

    public double OnSaturn() => ageOnEarth / 29.447498;

    public double OnUranus() => ageOnEarth / 84.016846;

    public double OnNeptune() => ageOnEarth / 164.79132;
}