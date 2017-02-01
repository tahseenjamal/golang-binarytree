package main

import (
	"fmt"
	"strconv"
)

type node interface {
	add(a int) bool
	getval() string
}

type node_element struct {
	left node

	element int

	right node
}

func (c *node_element) add(a int) bool {

	if c.element == 0 {

		c.element = a
	}

	if a > c.element {

		if c.right == nil {

			c.right = &node_element{element: a, left: nil, right: nil}

		} else {

			c.right.add(a)

		}
	}

	if a < c.element {

		if c.left == nil {

			c.left = &node_element{element: a, left: nil, right: nil}

		} else {

			c.left.add(a)

		}

	}

	return true

}

func (c *node_element) getval() string {

	var left, right string

	if c.left == nil {
		left = "nil"
	} else {
		left = c.left.getval()
	}

	if c.right == nil {
		right = "nil"
	} else {
		right = c.right.getval()

	}

	return "{" + left + "." + strconv.Itoa(c.element) + "." + right + "}"

}

func main() {

	a := &node_element{}

	a.add(1)

	fmt.Println(a.getval())

	a.add(2)

	fmt.Println(a.getval())

	fmt.Println(a.right.getval())

	a.right.add(10)

	fmt.Println(a.right.getval())

	a.add(3)

	fmt.Println(a.getval())

	a.add(4)

	fmt.Println(a.getval())

}
