using System;
using System.Collections.Generic;

public class CircularBuffer<T> : Queue<T>
{
    private readonly int capacity;

    public CircularBuffer(int capacity)
        : base(capacity)
    {
        this.capacity = capacity;
    }

    public T Read() => Dequeue();

    public void Write(T value)
    {
        if (Count == capacity)
        {
            throw new InvalidOperationException("Cannot write to full buffer.");
        }
        
        Enqueue(value);
    }

    public void Overwrite(T value)
    {
        if (Count == capacity)
        {
            Dequeue();
        }

        Enqueue(value);
    }
}