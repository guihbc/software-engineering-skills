package guihbc.queue;

public class Queue {
    private int size;
    private double[] data;
    private int head;
    private int last;
    private int length;

    public Queue(int size) {
        this.size = size;
        this.data = new double[size];
        this.head = 0;
        this.last = -1;
        this.length = 0;
    }

    public int getSize() {
        return size;
    }

    public double[] getData() {
        return data;
    }

    public void setData(double[] data) {
        this.data = data;
    }

    public int getHead() {
        return head;
    }

    public double getHeadValue() {
        return this.data[this.head];
    }

    public void setHead(int head) {
        this.head = head;
    }

    public int getLast() {
        return last;
    }

    public void setLast(int last) {
        this.last = last;
    }

    public int getLength() {
        return length;
    }
    
    public void setLength(int length) {
        this.length = length;
    }

    public void insert(double value) {
        if (this.last == this.size - 1) {
            throw new RuntimeException("queue overflow");
        }

        this.last++;
        this.data[this.last] = value;
        this.length++;
    }

    public void remove() {
        this.head++;

        if (this.head == this.size) {
            this.head = 0;
        }

        this.length--;
    }
}
