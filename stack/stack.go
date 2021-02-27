package stack

import (
	"container/list"
	"fmt"
)

type MyStack struct {
	Stack *list.List
}

func (c *MyStack) Push(value string) {
	c.Stack.PushFront(value)
}

func (c *MyStack) Pop() string {
	if c.Stack.Len() > 0 {
		ele := c.Stack.Front()
		c.Stack.Remove(ele)
		return ele.Value.(string)
	}
	return "null"
}

func (c *MyStack) Front() (string, error) {
	if c.Stack.Len() > 0 {
		if val, ok := c.Stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *MyStack) Size() int {
	return c.Stack.Len()
}

func (c *MyStack) Empty() bool {
	return c.Stack.Len() == 0
}
