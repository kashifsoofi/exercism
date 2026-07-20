// Package tree implements a simple library that
// builds a tree view of forum posts
package tree

import (
	"errors"
	"sort"
)

// Record struct
type Record struct {
	ID     int
	Parent int
}

// Node struct represnt Tree
type Node struct {
	ID       int
	Children []*Node
}

// Build builds and return Tree structure from Records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := make([]*Node, len(records))
	for i, record := range records {
		if record.ID != i || record.Parent > record.ID || record.ID > 0 && record.Parent == record.ID {
			return nil, errors.New("invalid record")
		}

		node := &Node{
			ID: record.ID,
		}
		nodes[i] = node
		if i > 0 {
			parentNode := nodes[record.Parent]
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	return nodes[0], nil
}
