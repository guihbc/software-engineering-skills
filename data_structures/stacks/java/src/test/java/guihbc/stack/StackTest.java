package guihbc.stack;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class StackTest {

    @Test
    void shouldPushStack() {
        Stack stack = new Stack(3);
        stack.push(150);
        stack.push(199);

        Assertions.assertEquals(2, stack.getLength());
        Assertions.assertFalse(stack.isFull());
        Assertions.assertFalse(stack.isEmpty());
    }

    @Test
    void shouldPopStack() {
        Stack stack = new Stack(3);
        stack.push(150);
        stack.push(199);
        stack.push(250);
        stack.pop();

        Assertions.assertEquals(2, stack.getLength());
        Assertions.assertFalse(stack.isFull());
        Assertions.assertFalse(stack.isEmpty());
    }

    @Test
    void shouldTestStackOverflow() {
        Stack stack = new Stack(3);
        stack.push(150);
        stack.push(199);
        stack.push(250);

        Assertions.assertThrows(RuntimeException.class, () -> stack.push(299));
        Assertions.assertEquals(3, stack.getLength());
        Assertions.assertFalse(stack.isEmpty());
    }

    @Test
    void shouldPopEmptyStack() {
        Stack stack = new Stack(3);

        Assertions.assertThrows(RuntimeException.class, stack::pop);
        Assertions.assertEquals(0, stack.getLength());
        Assertions.assertTrue(stack.isEmpty());
    }
}
