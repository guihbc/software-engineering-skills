package guihbc.stack;

public class Stack {
    private int top;
    private final int capacity;
    private final double[] data;
    private int length;

    public Stack(int capacity) {
        this.top = -1;
        this.capacity = capacity;
        this.data = new double[capacity];
        this.length = 0;
    }

    public int getLength() {
        return this.length;
    }

    public boolean isEmpty() {
        return this.top == -1;
    }

    public boolean isFull() {
        return this.top == this.capacity - 1;
    }

    public void push(double element) throws RuntimeException {
        if (this.isFull()) {
            throw new RuntimeException("stack overflow");
        }

        this.top++;
        this.data[this.top] = element;
        this.length++;
    }

    public double pop() {
        if (this.isEmpty()) {
            throw new RuntimeException("empty stack");
        }

        double element = this.data[this.top];
        this.top--;
        this.length--;

        return element;
    }

    public double getTopElement() {
        return this.data[this.top];
    }
}
