using System;

public class CircularBuffer<T>
{
    private readonly T[] buffer;
    private int count = 0;
    private int readIndex = 0;

    public CircularBuffer(int capacity)
    {
        buffer = new T[capacity];
        count = 0;
        readIndex = 0;
    }

    public T Read()
    {
        if (count == 0)
        {
            throw new InvalidOperationException("Cannot read empty buffer.");
        }

        readIndex = readIndex % buffer.Length;
        var value = buffer[readIndex++];
        count--;

        return value;
    }

    public void Write(T value)
    {
        if (count == buffer.Length)
        {
            throw new InvalidOperationException("Cannot write to full buffer.");
        }
        
        var writeIndex = (readIndex + count) % buffer.Length;
        buffer[writeIndex] = value;
        count++;
    }

    public void Overwrite(T value)
    {
        if (count < buffer.Length)
        {
            Write(value);
        }
        else
        {
            buffer[readIndex] = value;
            readIndex++;
        }
    }

    public void Clear()
    {
        count = 0;
    }
}