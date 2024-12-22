# Linked Lists

A linked list is a data structure that represents a sequence of data elements of the same type. Each data element is stored in a node, which contains the element itself and the address of the next node, that's why it's called a Linked list because each node is linked to the next one.

It's important to note that a linked list does not have a specific size limit; these data structures, known as dynamic lists, can grow or shrink in size as needed.

In the diagram below, you can see the operation of a linked list, where each node is "linked" to the next node, hence the name "linked list." Each node contains an object and a reference to the next node.

![linked_list](https://user-images.githubusercontent.com/48635609/102827421-4c225680-43c1-11eb-8f9d-67812d386563.gif)

Now, let's take a look at a Java code example representing a simple linked list implementation:

```java
// Node.java

public class Node<T> {
    private Node<T> next;
    private T value;

    public Node(T value) {
        this.value = value;
        this.next = null;
    }

    public Node next() {
        return this.next;
    }

    public void setNext(Node next) {
        this.next = next;
    }

    public T value() {
        return this.value;
    }

    public void setValue(T value) {
        this.value = value;
    }
}
```

```java
// LinkedList.java

public class LinkedList<T> {
    private Node<T> head;
    private Node<T> tail;
    private int length;

    public LinkedList() {
        this.head = null;
        this.tail = null;
        this.length = 0;
    }

    public boolean isEmpty() {
        return this.length == 0;
    }

    public int length() {
        return this.length;
    }

    public Node<T> head() {
        return this.head;
    }

    public Node<T> tail() {
        return this.tail;
    }

    public void add(T element) {
        Node<T> node = new Node<>(element);

        if (this.isEmpty()) {
            this.head = node;
            this.tail = node;
            this.length++;

            return;
        }

        this.tail.setNext(node);
        this.tail = node;
        this.length++;
    }

    public void remove(T value) {
        if (this.isEmpty()) {
            return;
        }

        if (value.equals(this.head.value())) {
            this.head = this.head.next();

            if (this.head == null) {
                this.tail = null;
            }

            this.length--;
            return;
        }

        Node<T> current = this.head;
        Node<T> previous = null;

        // Find the element
        while (current != null && !value.equals(current.value())) {
            previous = current;
            current = current.next();
        }

        // Element not found
        if (current == null) {
            return;
        }

        // "removing" the element
        previous.setNext(current.next());

        // "removing" the last element
        if (current == this.tail) {
            this.tail = previous;
        }

        this.length--;
    }
}
```

In the example above we have a simple `LinkedList` implementation, first we have a Node class with 2 properties:

- `next`: A Node representing the next node of the list
- `value`: The value stored in the list

In the LinkedList class we have 3 properties:

- `head`: The Node representing the first element in the list
- `tail`: The Node representing the last element in the list
- `length`: The list length

And the two main methods:

- `add(T element)`: This method checks if the list is empty. If it is, it sets the new node as both the head and the tail of the list and increments the length by 1. If the list is not empty, it appends the new node to the end of the list. In this case, the next pointer of the current tail is updated to reference the new node, and the tail is updated to the new node, effectively making it the new end of the list.

- `remove(T element)`: This method removes an element from the list by its value. It does not actually delete the node but instead adjusts the links between nodes to "ignore" the node with the target value.

First, it checks if the element to be removed is the head of the list. If so, it updates the head to the next node. If the new head is null, this means the list is now empty, and the tail is also set to null.

If the element is not the head, the method searches for the element in the list by comparing values in a loop. If the element is not found (i.e., current becomes null), the method simply returns. If the element is found, it updates the next pointer of the previous node to skip over the node being removed, thereby effectively removing it from the list. Finally, if the removed node was the tail, the tail is updated to the previous node. The length of the list is then decremented by 1.
