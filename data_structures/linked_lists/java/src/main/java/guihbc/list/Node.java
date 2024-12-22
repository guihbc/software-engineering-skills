package guihbc.list;

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
