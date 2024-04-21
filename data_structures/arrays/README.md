# Arrays

To understand what an array is, we first need to understand what a variable is.

In programming, a `variable` is a space allocated in the computer's RAM memory that we use to store a value or expression during the execution of a program.

Example:

```Go
variable := 10
name := "Guilherme"
```

In the example written in `Go`, two variables are declared, meaning that we allocate two spaces in the computer's RAM memory to store the values 10 and Guilherme, respectively, and we can use them as needed.

Once we understand what a variable is, we can then understand what an array is.

In programming, an `array`, also called a compound variable, is a set of elements of the same type and an specified size. The `array` is considered a data structure where each element can be identified by its index within the structure, and like a variable each element of the array is allocated in a memory address.

Example:

```Go
numbers := [5]int{1, 2, 3, 4, 5}
num := numbers[1]
```

In the example above, written in `Go`, we have an array of type `int` (integers) called `numbers`, which contains a set of elements. To access a specific element in `numbers`, we simply specify its index in the array, as seen in the variable `num`, which stores the value of `numbers` at position 1, meaning element `2`. In most programming languages, array indices generally start at position 0, so index 1 corresponds to element `2`.