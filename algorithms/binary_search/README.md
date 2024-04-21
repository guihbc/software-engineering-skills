# Binary Search

When we want to search for an element in an array, we can use a loop such as `for` or `foreach` to iterate through it and return the element when we find it, right? However, the number of iterations the algorithm will have will be quite high when the array contains a large number of elements. This type of search is called linear search, and its complexity function grows linearly.

With that said, let's look at an example of this type of algorithm:

```Go
package main

import "fmt"

func main() {
    numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
    fmt.Println(search(numbers, 30))
}

func search(arr []int, element int) int {
    for i := 0; i < len(arr); i++ {
        if arr[i] == element {
            return i
        }
    }

    return -1
}
```

The graph below illustrates the algorithm's complexity growing linearly, denoted by O(n).

![001_linear](https://user-images.githubusercontent.com/48635609/91091280-c36ef600-e62c-11ea-9f31-55708fb8196f.PNG)

To overcome this problem, there's an algorithm called binary search. Binary search is an algorithm used to find a specific item in an array, which must be sorted.

The algorithm works by repeatedly dividing the array in half (a paradigm called "divide and conquer") to break down our larger problem into smaller sub-problems and solve them independently. This reduces the number of interactions with the array and, consequently, reduces the algorithm's complexity.

Now, let's look at another example using binary search:

```Go
package main

import "fmt"

func main() {
	numbers := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(binarySearch(numbers, 30))
}

func binarySearch(arr []int, element int) int {
	var start, final, mid int
	start = 0
	final = len(arr) - 1

	for start <= final {
		mid = (start + final) / 2
		if arr[mid] == element {
			return mid
		}
		if arr[mid] < element {
			start = mid + 1
		} else {
			final = mid - 1
		}
	}

	return -1
}
```

The graph below shows the algorithm's complexity growing logarithmically, denoted by O(log n).

![002_log](https://user-images.githubusercontent.com/48635609/91092368-4a709e00-e62e-11ea-91b2-ddb17b872586.PNG)

The gif below demonstrates the difference in the number of interactions between linear search and binary search:

![004_compare](https://user-images.githubusercontent.com/48635609/91093901-73922e00-e630-11ea-98b0-a7370b856144.gif)

Analyzing the gif, we see each step that the binary search algorithm takes. It initially divides the array into two parts and checks if the middle element is the one we're looking for. If it is, it returns the position. If not, it checks whether the middle element is greater or smaller than the target element. If it's smaller, the algorithm knows the target element is in the right half of the array, and it can ignore the left half. The array is divided in half again, and the process is repeated until the desired value is found. This technique reduces the number of interactions the algorithm performs with the array, as demonstrated in the gif.

![003_table](https://user-images.githubusercontent.com/48635609/91092411-5eb49b00-e62e-11ea-9466-1d7a393f26e0.PNG)

### References

- [Binary Search](https://algoritmosempython.com.br/cursos/algoritmos-python/pesquisa-ordenacao/pesquisa-binaria)

- [Khan Academy - Binary Search](https://pt.khanacademy.org/computing/computer-science/algorithms/binary-search/a/binary-search)

- [Cyberini - Binary Search](https://www.blogcyberini.com/2017/09/busca-binaria.html)

### Additional Materials

- [How to implement BINARY SEARCH? *You should learn this!* | Algorithms #10](https://www.youtube.com/watch?v=EgLE5HwRy_M) - video

- [Binary Search](https://www.youtube.com/watch?v=l6pxuyV3mKQ) - video

- [Binary Search](https://www.youtube.com/watch?v=NxriUYWxoEU) - video

```
