package noam

import "C"
import (
	"fmt"
	"strings"
)

func toAnySlice(v any) []any {
	if v == nil {
		return nil
	}
	return v.([]any)
}

type Node struct {
	Head     string
	Leaf     string
	Children []*Node
}

func (n *Node) String() string {
	var sb strings.Builder
	sb.WriteString("[.")
	sb.WriteString(n.Head)
	if n.Leaf != "" {
		sb.WriteString(" ")
		sb.WriteString(n.Leaf)
	} else {
		for _, child := range n.Children {
			sb.WriteString(" ")
			sb.WriteString(child.String())
		}
	}
	sb.WriteString(" ]")
	return sb.String()
}

func makeNode(head any, value any) (*Node, error) {
	node := &Node{
		Head: head.(string),
	}
	switch v := value.(type) {
	default:
		return nil, fmt.Errorf("unknown type %T", v)
	case string:
		node.Leaf = v
	case []*Node:
		node.Children = v
	}
	return node, nil
}

func makeChildren(first any, rest any) ([]*Node, error) {
	pairs := toAnySlice(rest)
	var rs []*Node

	rs = append(rs, first.(*Node))
	for _, pair := range pairs {
		ps := toAnySlice(pair)
		rs = append(rs, ps[3].(*Node))
	}
	return rs, nil
}
