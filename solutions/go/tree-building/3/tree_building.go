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
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("invalid record")
		}

		nodes[i] = &Node{
			ID: r.ID,
		}
		if i > 0 {
			parentNode := nodes[r.Parent]
			parentNode.Children = append(parentNode.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
