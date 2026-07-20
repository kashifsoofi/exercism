using System;

public enum Direction
{
    North,
    East,
    South,
    West
}

public class RobotSimulator
{
    private Direction direction;
    private int x;
    private int y;

    public RobotSimulator(Direction direction, int x, int y)
    {
        this.direction = direction;
        this.x = x;
        this.y = y;
    }

    public Direction Direction => direction;

    public int X => x;

    public int Y => y;

    public void TurnRight()
    {
        switch (Direction)
        {
            case Direction.North:
                direction = Direction.East;
                break;
            case Direction.East:
                direction = Direction.South;
                break;
            case Direction.South:
                direction = Direction.West;
                break;
            case Direction.West:
                direction = Direction.North;
                break;
        }
    }

    public void TurnLeft()
    {
        switch (Direction)
        {
            case Direction.North:
                direction = Direction.West;
                break;
            case Direction.West:
                direction = Direction.South;
                break;
            case Direction.South:
                direction = Direction.East;
                break;
            case Direction.East:
                direction = Direction.North;
                break;
        }
    }

    public void Advance()
    {
        switch (Direction)
        {
            case Direction.North:
                y += 1;
                break;
            case Direction.East:
                x += 1;
                break;
            case Direction.South:
                y -= 1;
                break;
            case Direction.West:
                x -= 1;
                break;
        }
    }

    public void Move(string instructions)
    {
        foreach (var c in instructions)
        {
            switch (c)
            {
                case 'A':
                    Advance();
                    break;
                case 'L':
                    TurnLeft();
                    break;
                case 'R':
                    TurnRight();
                    break;
            }
        }
    }
}