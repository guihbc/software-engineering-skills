# Stacks

A stack is a data structure that allows the addition and removal of elements with a specific rule: the last element to be added is the first one to be removed. This rule follows the **LIFO (Last-In, First-Out)** order, which means that the last item to enter the stack is the first to leave. Stacks are similar to a real-life stack of plates, where the plate placed on top of the stack is the first to be removed.

Here's a simple explanation with a visual representation:

![stack](https://user-images.githubusercontent.com/48635609/102821122-4757a580-43b5-11eb-8e79-63c915d24378.gif)

In the example above, you can see how the stack works. New elements are added on top of the stack, and when an element is removed, it's always the last one (at the top of the stack) that gets removed.

Now, let's take a look at a Java code example representing a simple stack implementation:

```java
import java.util.List;
import java.util.ArrayList;
import java.util.function.Consumer;

public class Stack<T> {
    private int top;
    private int length;
    private List<T> data;

    public Stack() {
        this.top = -1;
        this.length = 0;
        this.data = new ArrayList<T>();
    }

    public boolean isEmpty() {
        return this.length == 0;
    }

    public void push(T element) {
        this.top++;
        this.length++;
        this.data.add(element);
    }

    public T getTop() {
        if (this.isEmpty()) {
            throw new RuntimeException("empty stack");
        }

        return this.data.get(top);
    }

    public T pop() {
        if (this.isEmpty()) {
            throw new RuntimeException("empty stack");
        }

        T element = this.data.remove(top);

        this.top--;
        this.length--;
        return element;
    }

    public void forEach(Consumer<T> consumer) {
        for (int i = this.top; i >= 0; i--) {
            consumer.accept(this.data.get(i));
        }
    }
}
```

In the code above we have a simple stack structure with 3 properties:

- `top`: The index of the top element.
- `length`: The length.
- `data`: A list of a generic type to store the data.

And we have 3 main methods:

- `push`: The top position will be the next one, the length will be incremented by 1 and the new element will be added to the top of the stack.
- `pop`: This method checks if the stack is empty. If it is, an exception is thrown indicating that the stack is empty. If it is not empty, it retrieves the top element, decrements the top index and the length by 1, and then returns the removed top element.
