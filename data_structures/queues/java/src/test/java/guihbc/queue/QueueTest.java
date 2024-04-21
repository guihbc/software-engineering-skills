package guihbc.queue;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import org.junit.jupiter.api.Test;

public class QueueTest {
    
    @Test
    public void shouldInsertInTheQueue() {
        Queue queue = new Queue(1);
        queue.insert(10);

        assertEquals(1, queue.getLength());
    }

    @Test
    public void shouldRemoveFromTheQueue() {
        Queue queue = new Queue(3);
        queue.insert(10);
        queue.insert(20);
        queue.insert(30);
        queue.remove();

        assertEquals(2, queue.getLength());
        assertEquals(20, queue.getHeadValue());
    }

    @Test
    public void shouldTestQueueOverflow() {
        Queue queue = new Queue(3);
        queue.insert(10);
        queue.insert(20);
        queue.insert(30);
        
        RuntimeException exception = assertThrows(RuntimeException.class, () -> queue.insert(40));
        assertEquals("queue overflow", exception.getMessage());
    }
}
