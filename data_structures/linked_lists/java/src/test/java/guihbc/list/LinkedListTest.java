package guihbc.list;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;

public class LinkedListTest {
    @Test
    public void shouldInsertInTheList() {
        LinkedList linkedList = new LinkedList();
        linkedList.insert(1);
        linkedList.insert(2);
        linkedList.insert(3);

        assertTrue(linkedList.length() == 3);

        Node third = linkedList.getHead();
        Node second = third.getNext();
        Node first = second.getNext();

        assertEquals(3, (int) third.getElement());
        assertEquals(2, (int) second.getElement());
        assertEquals(1, (int) first.getElement());
    }

    @Test
    public void shouldRemoveFromTheList() {
        LinkedList linkedList = new LinkedList();
        linkedList.insert(1);
        linkedList.insert(2);
        linkedList.insert(3);

        linkedList.remove(2);
        assertTrue(linkedList.length() == 2);

        Node third = linkedList.getHead();
        Node second = third.getNext();

        
        assertEquals(3, (int) third.getElement());
        assertEquals(1, (int) second.getElement());
    }

    @Test
    public void shouldSearchInTheList() {
        LinkedList linkedList = new LinkedList();
        linkedList.insert(1);
        linkedList.insert(2);
        linkedList.insert(3);

        Node node = linkedList.search(2);
        assertNotNull(node);
        assertEquals(2, node.getElement());
    }
}
