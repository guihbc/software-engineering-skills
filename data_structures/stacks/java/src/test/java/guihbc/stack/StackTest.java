package guihbc.stack;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class StackTest {

    @Test
    void shouldPushStack() {
        Stack<Integer> stack = new Stack();
        stack.push(150);
        stack.push(199);

        Assertions.assertEquals(2, stack.getLength());
        Assertions.assertFalse(stack.isEmpty());
    }

    @Test
    void shouldPopStack() {
        Stack<Integer> stack = new Stack();
        stack.push(150);
        stack.push(199);
        stack.push(250);
        stack.pop();

        Assertions.assertEquals(2, stack.getLength());
        Assertions.assertFalse(stack.isEmpty());
    }

    @Test
    void shouldNotPopEmptyStack() {
        Stack<Integer> stack = new Stack();

        Assertions.assertThrows(RuntimeException.class, stack::pop);
        Assertions.assertEquals(0, stack.getLength());
        Assertions.assertTrue(stack.isEmpty());
    }
}
