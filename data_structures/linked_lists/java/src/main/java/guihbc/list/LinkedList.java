package guihbc.list;

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
