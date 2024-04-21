package guihbc.list;

public class Node {
    private Object element;
    private Node next;
    
    public Object getElement() {
        return element;
    }
    public void setElement(Object element) {
        this.element = element;
    }
    public Node getNext() {
        return next;
    }
    public void setNext(Node next) {
        this.next = next;
    }

    public boolean isEmpty() {
        return this == null;
    }
}
