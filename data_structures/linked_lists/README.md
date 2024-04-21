# Linked Lists

A linked list is a data structure that represents a sequence of data elements of the same type. Each data element is stored in a node, which contains the element itself and the address of the next node. In this explanation, elements are of type `interface{}`, meaning they can be any data type.

It's important to note that a linked list does not have a specific size limit; these data structures, known as dynamic lists, can grow or shrink in size as needed.

In the diagram below, you can see the operation of a linked list, where each node is "linked" to the next node, hence the name "linked list." Each node contains an object and a reference to the next node.

![linked_list](https://user-images.githubusercontent.com/48635609/102827421-4c225680-43c1-11eb-8f9d-67812d386563.gif)

Here's an example of a linked list implemented in Go:

```Go
package main

import (
    "fmt"
)

type Node struct {
    element interface{}
    next    *Node
}

func (n *Node) isEmpty() bool {
    return n == nil
}

type LinkedList struct {
    first *Node
}

func (l *LinkedList) create() {
    l.first = nil
}

func (l *LinkedList) insert(e interface{}) {
    var node *Node = &Node{}
    node.element = e
    node.next = l.first
    l.first = node
}

func (l *LinkedList) isEmpty() bool {
    return l.first == nil
}

func (l *LinkedList) search(e interface{}) *Node {
    for n := l.first; n != nil; n = n.next {
        if n.element == e {
            return n
        }
    }

    return nil
}

func (l *LinkedList) remove(e interface{}) int {
    var previous *Node = nil
    var p *Node = l.first

    for p != nil && p.element != e {
        previous = p
        p = p.next
    }

    if p == nil {
        return -1
    }

    if previous == nil {
        l.first = p.next
    } else {
        previous.next = p.next
    }

    return 0
}

func (l *LinkedList) free() {
    for l.first != nil {
        var temp *Node = l.first.next
        l.first = nil
        l.first = temp
    }
}

func (l *LinkedList) show() {
    fmt.Print("[ ")
    print(l.first)
    fmt.Println("]")
}

func print(n *Node) {
    if !n.isEmpty() {
        print(n.next)
        fmt.Print(n.element)
        fmt.Print(" ")
    }
}

func main() {
    var list *LinkedList = &LinkedList{}
    list.create()
    list.insert(1)
    list.insert(2)
    list.insert(3)
    list.insert(4)
    list.insert(5)
    list.insert(6)
    list.insert(7)
    list.show()
    list.remove(2)
    list.show()
}
```

In the code above, the main components are as follows:

1. A `Node` struct representing the nodes of the linked list, with the attributes:

   - `element` - The data to be stored in the node.
   - `next` - The address of the next node in the list.

   The `isEmpty` method checks if a node is empty (nil).

2. A `LinkedList` struct representing the linked list, with the attribute:

   - `first` - The address of the first node in the list.

   The `create` method initializes a new list.

   The `insert` method inserts a new element into the list. It creates a new node, sets its element to the given value, and makes it point to the current first node. Then, it updates the first node to be the newly created node.

   The `isEmpty` method checks if the list is empty.

   The `search` method searches for a specific element in the list. It iterates through the nodes in the list and returns the node containing the element if found, or nil if not.

   The `remove` method removes an element from the list. It creates two nodes: one that keeps track of the previous node (before the one to be removed) and another to traverse the list. It iterates through the list to find the element to be removed. If it's found, the `previous` node is updated to point to the node after the one being removed. If the `previous` node is nil, it means the first node is being removed, so the first node is updated to the next node. If the `previous` node is not nil, it simply bypasses the node to be removed.

   The `free` method frees up memory by removing all nodes in the list.

   The `show` method displays the elements of the list.

   The `print` function is a helper function to display the elements of the list in the `show` method.

In the `main` function, an example linked list is created, elements are inserted, and then one element is removed, followed by displaying the list before and after the removal.

The comments in the code provide additional explanations for each part.