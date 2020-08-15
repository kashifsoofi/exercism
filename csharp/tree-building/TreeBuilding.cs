using System;
using System.Collections.Generic;
using System.Linq;

public class TreeBuildingRecord
{
    public int ParentId { get; set; }
    public int RecordId { get; set; }
}

public class Tree
{
    public int Id { get; set; }
    public int ParentId { get; set; }

    public List<Tree> Children { get; set; }

    public bool IsLeaf => Children.Count == 0;
}

public static class TreeBuilder
{
    public static Tree BuildTree(IEnumerable<TreeBuildingRecord> records)
    {
        if (records.Count() == 0)
        {
            throw new ArgumentException();
        }

        List<Tree> trees = new List<Tree>();
        var orderedRecords = records.OrderBy(x => x.RecordId).ToList();
        for (int i = 0; i < orderedRecords.Count; i++)
        {
            var record = orderedRecords[i];

            if (record.RecordId != i ||
                (record.ParentId != 0 && record.ParentId >= record.RecordId))
            {
                throw new ArgumentException();
            }

            var item = new Tree
            {
                Id = record.RecordId,
                ParentId = record.ParentId,
                Children = new List<Tree>(),
            };
            trees.Add(item);
            if (i > 0)
            {
                trees[record.ParentId].Children.Add(item);
            }
        }

        return trees[0];
    }
}