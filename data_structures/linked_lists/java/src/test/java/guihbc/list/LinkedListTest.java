package guihbc.list;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;

public class LinkedListTest {
    @Test
    public void shouldInsertInTheList() {
        LinkedList<Integer> linkedList = new LinkedList();
        linkedList.add(1);
        linkedList.add(2);
        linkedList.add(3);

        assertTrue(linkedList.length() == 3);
    }

    @Test
    public void shouldRemoveFromTheList() {
        LinkedList<Integer> linkedList = new LinkedList();
        linkedList.add(0);
        linkedList.add(1);
        linkedList.add(2);

        linkedList.remove(2);
        assertTrue(linkedList.length() == 2);
        assertTrue(linkedList.tail().value() == 1);
    }
}
