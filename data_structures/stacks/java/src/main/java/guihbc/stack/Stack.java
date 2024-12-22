package guihbc.stack;

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

    public int getLength() {
        return this.length;
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
