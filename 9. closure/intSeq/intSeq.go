/* https://gobyexample.com/closures
Go supports anonymous functions, which can form closures.

Benefits:
1. Encapsulating private state.
2. Function factories: Closure can act as factories for generating specialized functions based on
   specific config or parameters.
3. Maintaining state across multiple calls. Retain state across successive calls.
4. Callbacks and event handlers. commonly used.
5. Asynchronous operations.

*/

package main

import "fmt"

// This function intSeq returns another function,
// which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()

	fmt.Println(nextInt()) // output: 1
	fmt.Println(nextInt()) // output: 2
	fmt.Println(nextInt()) // output: 3

	newInts := intSeq()
	fmt.Println(newInts()) // output: 1
}
