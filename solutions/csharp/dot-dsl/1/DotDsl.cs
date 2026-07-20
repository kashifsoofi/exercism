using System.Collections;
using System.Collections.Generic;

public class Node : AttrCollection
{
    public string Name { get; }

    public Node(string name)
    {
        Name = name;
    }

    public override bool Equals(object obj)
    {
        var n = obj as Node;
        if (n == null)
        {
            return false;
        }

        return Name.Equals(n.Name);
    }

    public override int GetHashCode() => Name.GetHashCode();
}

public class Edge : AttrCollection
{
    private readonly string from;
    private readonly string to;

    public Edge(string from, string to)
    {
        this.from = from;
        this.to = to;
    }

    public override bool Equals(object obj)
    {
        var e = obj as Edge;
        if (e == null)
        {
            return false;
        }

        return from.Equals(e.from) && to.Equals(e.to);
    }

    public override int GetHashCode() => from.GetHashCode() ^ to.GetHashCode();
}

public class Attr
{
    public string Name { get; }
    public string Value { get; }

    public Attr(string name, string value)
    {
        Name = name;
        Value = value;
    }

    public override bool Equals(object obj)
    {
        var a = obj as Attr;
        if (a == null)
        {
            return false;
        }


        return Name.Equals(a.Name) && Value.Equals(a.Value);
    }

    public override int GetHashCode() => Name.GetHashCode() ^ Value.GetHashCode();
}

public class AttrCollection : IEnumerable<Attr>
{
    public List<Attr> Attrs { get; } = new List<Attr>();

    public void Add(Attr attr) => Attrs.Add(attr);

    public void Add(string name, string value) => Add(new Attr(name, value));

    public IEnumerator<Attr> GetEnumerator() => Attrs.GetEnumerator();

    IEnumerator IEnumerable.GetEnumerator() => GetEnumerator();
}

public class Graph : AttrCollection
{
    public List<Node> Nodes { get; } = new List<Node>();

    public List<Edge> Edges { get; } = new List<Edge>();

    public void Add(Node node) => Nodes.Add(node);

    public void Add(Edge edge) => Edges.Add(edge);

}