using System;
using System.Text;

public static class RotationalCipher
{
    public static string Rotate(string text, int shiftKey)
    {
        var shifted = new StringBuilder();
        foreach (var ch in text)
        {
            shifted.Append(Rotate(ch, shiftKey));
        }

        return shifted.ToString();
    }

    private static char Rotate(char character, int shiftKey)
    {
        if (!char.IsLetter(character))
        {
            return character;
        }

        var offset = char.IsUpper(character) ? 'A' : 'a';
        return (char)(((character + shiftKey) - offset) % 26 + offset);
    }
}