package guihbc.queue;

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

    public int getLength() {
        return this.length;
    }
}
