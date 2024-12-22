package guihbc.queue;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import org.junit.jupiter.api.Test;

public class QueueTest {

    @Test
    public void shouldInsertInTheQueue() {
        Queue<Integer> queue = new Queue();
        queue.enqueue(10);

        assertEquals(1, queue.getLength());
    }

    @Test
    public void shouldRemoveFromTheQueue() {
        Queue<Integer> queue = new Queue();
        queue.enqueue(10);
        queue.enqueue(20);
        queue.enqueue(30);
        queue.dequeue();

        assertEquals(2, queue.getLength());
    }
}
