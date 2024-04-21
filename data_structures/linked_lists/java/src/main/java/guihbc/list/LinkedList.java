package guihbc.list;

public class LinkedList {
    private Node head;
    private int length;

    public Node getHead() {
        return head;
    }

    public void setHead(Node first) {
        this.head = first;
    }

    public void insert(Object element) {
        Node node = new Node();
        node.setElement(element);
        node.setNext(this.head);
        this.setHead(node);
        length++;
    }

    public boolean isEmpty() {
        return this.head == null;
    }

    public Node search(Object element) {
        for(Node n = this.head; n != null; n = n.getNext()) {
            if(n.getElement() == element) {
                return n;
            }
        }

        return null;
    }

    public int length() {
        return this.length;
    }

    public int remove(Object element) {
        Node previous = null;
        Node first = this.head;

        while(first != null && first.getElement() != element) {
            previous = first;
            first = first.getNext();
        }

        if (first == null) {
            return -1; // Not found
        }

        if (previous == null) {
            this.head = first;
        } else {
            previous.setNext(first.getNext());
        }

        length--;
        return 0;
    }
}
