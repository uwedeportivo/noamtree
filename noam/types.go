package noam

import "strings"

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

func makeChildren(v any) []*Node {
	cs := v.([]any)
	var result []*Node
	for _, c := range cs {
		child := c.(*Node)
		result = append(result, child)
	}
	return result
}
