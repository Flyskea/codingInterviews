package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	data []string
	str  strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		c.str.WriteString("null,")
	} else {
		c.str.WriteString(strconv.Itoa(root.Val))
		c.str.WriteString(",")
		c.serialize(root.Left)
		c.serialize(root.Right)
	}
	return c.str.String()
}

func (c *Codec) deserializeCore() *TreeNode {
	if c.data[0] == "null" {
		c.data = c.data[1:]
		return nil
	}
	v, _ := strconv.Atoi(c.data[0])
	root := &TreeNode{
		Val: v,
	}
	c.data = c.data[1:]
	root.Left = c.deserializeCore()
	root.Right = c.deserializeCore()
	return root
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	c.data = l
	return c.deserializeCore()
}

func main() {
	c := Constructor()
	s := c.serialize(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	})
	fmt.Println(c.deserialize(s))
}
