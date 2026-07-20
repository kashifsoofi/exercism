using System;
using System.Collections.Generic;

public class CircularBuffer<T>
{
    private readonly Queue<T> buffer;
    private readonly int capacity;

    public CircularBuffer(int capacity)
    {
        this.capacity = capacity;
        buffer = new Queue<T>();
    }

    public T Read()
    {
        if (buffer.Count == 0)
        {
            throw new InvalidOperationException("Cannot read empty buffer.");
        }

        return buffer.Dequeue();
    }

    public void Write(T value)
    {
        if (buffer.Count == capacity)
        {
            throw new InvalidOperationException("Cannot write to full buffer.");
        }
        
        buffer.Enqueue(value);
    }

    public void Overwrite(T value)
    {
        if (buffer.Count == capacity)
        {
            buffer.Dequeue();
        }
        Write(value);
    }

    public void Clear()
    {
        buffer.Clear();
    }
}