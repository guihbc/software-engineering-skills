# Queues

A queue is a data structure that allows the addition and removal of elements with a specific rule: the first element to be added is the first one to be removed. This rule follows the **FIFO (First-In, First-Out)** order, which means that the first item to enter the queue is also the first to leave. Queues are similar to a real-life queue of people, where the person who arrives first is the one to be served first.

Here's a simple explanation with a visual representation:

![queue](https://user-images.githubusercontent.com/48635609/102701712-87454e00-4238-11eb-8a5a-973837dad29e.gif)

In the example above, you can see how the queue works. New elements are added at the back of the queue, and when an element is removed, it's always the first one (at the front of the queue) that gets removed.

Now, let's look at a Java code example that represents a simple implementation of a queue:

```java
import java.util.LinkedList;
import java.util.List;

public class Queue<T> {
    private List<T> data;
    private int length;

    public Queue() {
        this.data = new LinkedList<>();
        this.length = 0;
    }

    public void enqueue(T value) {
        this.data.add(value);
        this.length++;
    }

    public T dequeue() {
        if (this.length == 0) {
            throw new RuntimeException("Empty queue");
        }

        this.length--;
        return this.data.remove(0);
    }
}
```

In the code above we have a simple queue structure with 2 properties:

- `data`: A List to store the queue data
- `length`: The queue length

And we have 2 main methods:

- `enqueue`: This method adds the new element to the List and increments the `length` variable by 1.
- `dequeue`: This method verifies if the queue is empty if it's true an exception will be thrown saying the queue is empty, if it's false the length variable is decreased by 1, then the first element of the list is removed and returned.
