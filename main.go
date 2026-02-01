package main

import (
	"fmt"
	"os"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("prime check requires number >= 2")
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, fmt.Errorf("negative exponents not supported")
	}
	if exponent == 0 && base == 0 {
		return 0, fmt.Errorf("MATH ERROR!! 0^0 is undefined")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

func makeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func makeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial

	add = func(x int) {
		total += x
	}

	subtract = func(x int) {
		total -= x
	}

	get = func() int {
		return total
	}

	return
}

func Apply(nums []int, operation func(int) int) []int {
	returned := make([]int, len(nums))
	for i, v := range nums {
		returned[i] = operation(v)
	}
	return returned
}

func Filter(nums []int, predicate func(int) bool) []int {
	var filtered []int
	for _, v := range nums {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	result := initial
	for _, v := range nums {
		result = operation(result, v)
	}
	return result
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	fmt.Printf("Current Process ID: %d\n", os.Getpid())
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())

	data := []int{1, 2, 3, 4}
	fmt.Printf("Memory address of slice: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])
	fmt.Println("Note: Other processes can't see these memory addresses due to process isolation.")

	/*  --> A process ID is a unique number assigned by the operating system kernel to identify a running process.
	    --> Process isolation is important as it prevents one process from accessing or modifying the memory of another process
			       without proper permissions, enhancing security and stability.

		--> The difference between the slice header address and the element addresses is that the slice header is stored directly in memory,
		while the elements are stored in a contiguous block of memory pointing to the header. */
}

func DoubleValue(x int) int {
	return x * 2
	/*It wil not modify the original variable as primitives in GO are immutable and are passed by value
	so a copy is made and sent to the function */
}

func DoublePointer(x *int) {
	*x = *x * 2
	/*It will modify the original variable as we are passing the address of the variable to the function
	and dereferencing it to change its value */
}

func CreateOnStack() int {
	y := 2008
	return y
	/* This variable stays on the stack */
}

func CreateOnHeap() *int {
	x := new(int)
	*x = 3008
	return x
	/* This variable is escapes to the heap */
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

func AnalyzeEscape() {
	CreateOnStack()
	CreateOnHeap()

	/*
		--> Variables captured by closures such as counters and accumulators escape to the heap because they must live beyond the function call.

		-->new(int) escapes to the heap when a pointer is returned from the function CreateOnHeap.

		--> Slices created with make or literals may escape to the heap when their backing array must outlive the function's scope or is returned.

		--> Function literals escape to the heap when they capture variables from their scope.

		--> Some return values which is ~r0 and string literals passed to fmt functions also escape to the heap.

		--> Escaping to the heap means the compiler determined the value(s) must survive beyond the stack frame where it was created.*/
}

func main() {

	fmt.Println("=== Process Information ===")
	ExploreProcess()
	fmt.Println()

	fmt.Println("=== Math Operations ===")

	facts := []int{0, 5, 10}
	for _, n := range facts {
		result, err := Factorial(n)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Factorial(%d) = %d\n", n, result)
		}
	}

	primes := []int{17, 20, 25, 15, -5, 2, 1, 13}
	for _, n := range primes {
		isP, err := isPrime(n)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("IsPrime(%d) = %v\n", n, isP)
		}
	}

	powers := [][2]int{{2, 8}, {5, 3}}
	for _, p := range powers {
		result, err := Power(p[0], p[1])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Power(%d, %d) = %d\n", p[0], p[1], result)
		}
	}

	fmt.Println()

	fmt.Println("=== Closure Demonstration ===")

	counter1 := makeCounter(0)
	counter2 := makeCounter(100)

	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter2:", counter2())
	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter2:", counter2())

	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)

	fmt.Println("Doubler(5):", doubler(5))
	fmt.Println("Tripler(5):", tripler(5))

	fmt.Println()

	fmt.Println("=== Higher-Order Functions ===")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original:", nums)

	squared := Apply(nums, func(x int) int { return x * x })
	fmt.Println("Squared:", squared)

	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println("Evens:", evens)

	sum := Reduce(nums, 0, func(acc, curr int) int {
		return acc + curr
	})
	fmt.Println("Sum:", sum)

	doubleThenAddTen := Compose(
		func(x int) int { return x + 10 },
		func(x int) int { return x * 2 },
	)
	fmt.Println("Compose (double then add 10) on 5:", doubleThenAddTen(5))

	fmt.Println()

	fmt.Println("=== Pointer Demonstration ===")

	a, b := 5, 10
	fmt.Printf("Before SwapValues: a=%d, b=%d\n", a, b)
	a, b = SwapValues(a, b)
	fmt.Printf("After SwapValues: a=%d, b=%d (unchanged originals)\n", a, b)

	x, y := 3, 7
	fmt.Printf("Before SwapPointers: x=%d, y=%d\n", x, y)
	SwapPointers(&x, &y)
	fmt.Printf("After SwapPointers: x=%d, y=%d (modified via pointers)\n", x, y)

	fmt.Println()

	fmt.Println("=== Escape Analysis Demonstration ===")
	AnalyzeEscape()
}
