package main

import (
	"fmt"
)

// The function that is deferred here is the `fmt.Print()` statement
// as a result when the d1() function is about to return, you get the
// three values of the i variable of the for loop in reverse order.
// This is due to LIFO order.
func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

// The value returned here is really strange.
// The reason for that is after the for loop ended,
// the value of i is 0, because it is that value of i
// that made the for loop terminate. The tricky point
// here is that the deferred anonymous function is
// evaluated after the for loop ends since there are
// no parameters. Which means that it is evaluated
// three times for an i value of 0.
func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}() // this is dangerous as the value of i depends
		// on when the anonymous function is executed
	}
	fmt.Println()
}

// Due to the parameter of the anonymous function, each time
// the anonymous function is deferred, it gets and therefore uses
// the current value of i. Each execution of the anonymous function
// has a different value to process without any ambiguities. This is
// the best apprach to using defer. Since you are intentionally
// passing the desired variable in the anonymous function in an easy-to-read way.
func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
