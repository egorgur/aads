package compress

import (
	"image/color"
)

type QuadTree struct {
	FirstNode *Node
}

func NewQuadTree(fistNode *Node) *QuadTree {
	return &QuadTree{
		FirstNode: fistNode,
	}
}
// A color block
type Node struct {
	Parent   *Node
	Children []*Node
	x0       int
	y0       int
	width    int
	height   int
	color    color.Color
}

func NewNode(parent *Node, x0, y0, width, height int, color color.Color) *Node {
	return &Node{
		Parent:   parent,
		Children: make([]*Node, 4),
		x0:       x0,
		y0:       y0,
		height:   height,
		width:    width,
		color:    color,
	}
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}
func (n *Node) HasChild() bool {
	for _, v := range n.Children {
		if v != nil {
			return true
		}
	}
	return false
}
