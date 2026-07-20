using System;
using System.Linq;

public class Matrix
{
    private readonly int[][] _matrix;

    public Matrix(string input)
    {
        var rows = input.Split("\n", StringSplitOptions.RemoveEmptyEntries);
        _matrix = new int[rows.Length][];
        for (var i = 0; i < rows.Length; i++)
        {
            var vals = rows[i].Split(" ", StringSplitOptions.RemoveEmptyEntries).Select(Int32.Parse).ToArray();
            _matrix[i] = new int[vals.Length];
            for (var j = 0; j < vals.Length; j++)
            {
                _matrix[i][j] = vals[j];
            }
        }
    }

    public int Rows => _matrix.Length;

    public int Cols => _matrix[0].Length;

    public int[] Row(int row) => _matrix[row-1];

    public int[] Column(int col)
    {
        var column = new int[_matrix.Length];
        for (var i = 0; i < _matrix.Length; i++)
        {
            column[i] = _matrix[i][col-1];
        }
        return column;
    }
}