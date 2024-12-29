package main

import (
	"fmt"
	"io"
	"math/rand/v2"
	"reflect"
	"runtime"
	"runtime/debug"
	"slices"
	"testing"
	"unsafe"
)

// EXERCISE: Fix the backing array problem
//
//	Ensure that changing the elements of the `mine` slice
//	does not change the elements of the `nums` slice.
func TestFixBackingArrayProblem(t *testing.T) {
	// Arrange
	nums := []int{56, 89, 15, 25, 30, 50}
	expectedNums := []int{56, 89, 15}
	// create a nil slice - the backing array is not created in memory yet

	// Act
	// use append to populate and create a new backing array under the hood
	mine := append([]int(nil), nums...)
	mine[0], mine[1], mine[2] = -50, -100, -150

	// Assert
	if !reflect.DeepEqual(nums[:3], expectedNums) {
		t.Errorf("Expected nums to be %v, but got %v", expectedNums, nums[:3])
	}
}

// EXERCISE: Sort the backing array
//
//  1. Sort only the middle 3 items.
//
//  2. All the slices should see your changes.
//
// RESTRICTION
//
//	Do not sort manually. Sort by using the sort package.
//
// EXPECTED OUTPUT
//
//	Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
//
//	Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
//
// HINT:
//
//	Middle items are         : [frogger asteroids simcity]
//
//	After sorting they become: [asteroids frogger simcity]
func TestSortBackingArrayInPlace(t *testing.T) {
	// Arrange
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}
	expectedItems := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "asteroids",
		"frogger", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	// Act
	slices.Sort(items[5:7])

	// Assert
	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Expected items to be %v, but got %v", expectedItems, items)
	}
}

func TestSortSliceWithoutMutatingArray(t *testing.T) {
	// Arrange
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}
	expectedItems := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	// Act
	newItems := append([]string(nil), items...)
	slices.Sort(newItems[5:7])

	// Assert
	if reflect.DeepEqual(items, newItems) { // shouldn't be the same
		t.Errorf("Expected items to be %v, but got %v", expectedItems, items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Expected items to be %v, but got %v", expectedItems, items)
	}
}

// EXERCISE: Observe the memory allocations
//
//	In this exercise, your goal is to observe the memory allocation
//	differences between arrays and slices.
//
//	You will create, assign arrays and slices then you will print
//	the memory usage of your program on each step.
//
//	Please follow the instructions inside the code.
//
// EXPECTED OUTPUT
//
//	Note that, your memory usage numbers may vary. However, the size of the
//	arrays and slices should be the same on your own system as well
//	(if you're on a 64-bit machine).
//
//
//	[initial memory usage]
//	        > Memory Usage: 104 KB
//	[after declaring an array]
//	        > Memory Usage: 78235 KB
//	[after copying the array]
//	        > Memory Usage: 156365 KB
//	[inside passArray]
//	        > Memory Usage: 234495 KB
//	[after slicings]
//	        > Memory Usage: 234497 KB
//	[inside passSlice]
//	        > Memory Usage: 234497 KB
//
//	Array's size : 80000000 bytes.
//	Array2's size: 80000000 bytes.
//	Slice1's size: 24 bytes.
//	Slice2's size: 24 bytes.
//	Slice3's size: 24 bytes.
//
// HINTS
//
//	I've declared a few functions to help you.
//
//	  report function:
//	  - Prints the memory usage.
//	  - Just call it with a message that matches to the expected output.
//
//	  passArray function:
//	  - It automatically prints the memory usage.
//	  - Accepts a [size]int array, so you can pass your array to it.
//
//	  passSlice function:
//	  - It automatically prints the memory usage.
//	  - Accepts an int slice, so you can pass it one of your slices.
const size = 1e7

func TestObserveMemoryAllocations(t *testing.T) {
	// Arrange

	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	var nums [size]int

	fmt.Printf("capacity of nums: %d\n", cap(nums))
	fmt.Printf("length of nums: %d\n", len(nums))

	// 2. print the memory usage (use the report func).
	report("after declaring an array")

	// 3. copy the array to a new array.
	numsCopy := nums

	// 4. print the memory usage
	report("after copying the array")

	// 5. pass the array to the passArray function
	passArray(nums)

	// 6. convert one of the arrays to a slice
	numsSlice := nums[:]

	// 7. slice only the first 1000 elements of the array
	slice1 := nums[:1000]

	// 8. slice only the elements of the array between 1000 and 10000
	slice2 := nums[1000:10000]

	// 9. print the memory usage (report func)
	report("after slicings")

	// 10. pass the one of the slices to the passSlice function
	passSlice(slice1)

	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Printf("\nArray's size : %d bytes.\n", unsafe.Sizeof(nums))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(numsCopy))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(numsSlice))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}

// EXERCISE: Observe the length and capacity
//
//	Follow the instructions inside the code below to
//	gain more intuition about the length and capacity of a slice.
func TestObserveLengthAndCapacity(t *testing.T) {
	// 1. create a new slice named: games
	games := []string{}

	// 2. print the length and capacity of the games slice
	fmt.Printf("[nil slice] \n\tlen: %d, cap: %d\n", len(games), cap(games))

	// 3. comment out the games slice
	//    then declare it as an empty slice
	games = []string{}

	// 4. print the length and capacity of the games slice
	fmt.Printf("[empty slice] \n\tlen: %d, cap: %d\n", len(games), cap(games))

	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	games = append(games, "pacman", "mario", "tetris", "doom")

	// 6. print the length and capacity of the games slice
	fmt.Printf("[initialize with append]\n\t len: %d, cap: %d\n", len(games), cap(games))

	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 5)
	games = []string{"pacman", "mario", "tetris", "doom"}

	fmt.Printf("[initialize with 4 elements]\n\t len: %d, cap: %d\n", len(games), cap(games))

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	for i := 0; i < 4; i++ {
		s := games[i:]
		// 2. print its length and capacity along the way (in the loop).
		fmt.Printf("[inside loop] \n\t len: %d, cap: %d\n", len(s), cap(s))
	}

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := games[:0]
	// 2. print the games and the new slice's len and cap
	fmt.Printf("[zero slice] \n\t len: %d, cap: %d\n", len(games), cap(games))
	// 3. append a new element to the new slice
	zero = append(zero, "pong")
	// 4. print the new slice's lens and caps
	fmt.Printf("[zero slice after append] \n\t len: %d, cap: %d\n", len(zero), cap(zero))
	// 5. repeat the last two steps 5 times (use a loop)
	for i := 0; i < 5; i++ {
		zero = append(zero, "dagger")
		fmt.Printf("[zero slice after append %d] \n\t len: %d, cap: %d\n", i, len(zero), cap(zero))
	}

	// 6. notice the growth of the capacity after the 5th append

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	for i := range zero {
		s := zero[:i+1]
		fmt.Printf("zero[:%d] len: %d cap: %d\n", i+1, len(s), cap(s))
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	zero = zero[:cap(zero)]

	// 2. print the elements of the zero slice in the loop.
	for i := range zero {
		fmt.Printf("zero[:%d] len: %d cap: %d - %q\n", i+1, len(zero[:i+1]), cap(zero[:i+1]), zero[:i+1])
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	zero[0] = "changed"

	// 2. change the same element of the games slice
	games[0] = "changed"

	// 3. print the games and the zero slices
	fmt.Printf("games: %q\n", games)
	fmt.Printf("zero: %q\n", zero)

	// 4. observe that they don't have the same backing array

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	// Uncommenting the following line will cause a runtime panic
	// fmt.Println(games[:cap(games)+1])
}

// EXERCISE: Limit the backing array sharing
//
//	GOAL
//
//	  Limit the capacity of the slice that is returned
//	  from the `Read` function. Read on for more details.
//
//	WHAT IS THE PROBLEM?
//
//	  `Read` function returns a portion of
//	  its `temps` slice. Below, it saves it to the
//	  `received` slice.
//
//	  `main()` appends to the `received` slice but doing so
//	  also changes the backing array of the `temps` slice.
//	  We don't want that.
//
//	  `main()` can change the part of the `temps` slice
//	  that is returned from the `Read()`, but it shouldn't
//	  be able to change the elements in the rest of the
//	  `temps`.
//
//	WHAT YOU NEED TO DO?
//
//	  So you need to limit the capacity of the returned
//	  slice somehow. Remember: `received` and `temps`
//	  share the same backing array. So, appending to it
//	  can overwrite the same backing array.
//
// CURRENT
//
//	                        | |
//	                        v v
//	temps     : [5 10 3 1 3 80 90]
//	received  : [5 10 3 1 3]
//	                        ^ ^ append changes the `temps`
//	                            slice's backing array.
//
// EXPECTED
//
//	The corrected api package does not allow the `main()` to
//	change unreturned portion of the temps slice's backing array.
//	                        |  |
//	                        v  v
//	temps     : [5 10 3 25 45 80 90]
//	received  : [5 10 3 1 3]
func TestLimitBackingArraySharing(t *testing.T) {
	// Arrange
	expectedTemps := []int{5, 10, 3, 25, 45, 80, 90}
	expectedReceived := []int{5, 10, 3, 1, 3}

	// Act
	received := Read(0, 3)
	received = append(received, []int{1, 3}...)

	// Assert
	if !reflect.DeepEqual(All(), expectedTemps) {
		t.Errorf("Expected api.temps to be %v, but got %v", expectedTemps, All())
	}
	if !reflect.DeepEqual(received, expectedReceived) {
		t.Errorf("Expected main.received to be %v, but got %v", expectedReceived, received)
	}
}

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a slice of elements from the temps slice.
func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA

	portion := append([]int(nil), temps[start:stop]...)

	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	// ----------------------------------------

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}

// EXERCISE: Fix the memory leak
//
//	WARNING
//
//	  This is a very difficult exercise. You need to
//	  do some research on your own to solve it. Please don't
//	  get discouraged if you can't solve it yet.
//
//	GOAL
//
//	  In this exercise, your goal is to reduce the memory
//	  usage. To do that, you need to find and fix the memory
//	  leak within `main()`.
//
//	PROBLEM
//
//	  `main()` calls `api.Report()` that reports the current
//	  memory usage.
//
//	  After that, `main()` calls `api.Read()` that returns
//	  a slice with 10 millions of elements. But you only need
//	  the last 10 elements of the returned slice.
//
//	WHAT YOU NEED TO DO
//
//	  You only need to change the code in `main()`. Please
//	  do not touch the code in `api/api.go`.
//
//	CURRENT OUTPUT
//
//	  > Memory Usage: 113 KB
//
//	  Last 10 elements: [...]
//
//	  > Memory Usage: 65651 KB
//
//	    + Before `api.Read()` call: It uses 113 KB of memory.
//
//	    + After `api.Read()` call : It uses  65 MB of memory.
//
//	    + This means that, `main()` never releases the memory.
//	      This is the leak.
//
//	    + Your goal is to release the unused memory. Remember,
//	      you only need 10 elements but in the current code
//	      below you have a slice with 10 millions of elements.
//
//	EXPECTED OUTPUT
//
//	  > Memory Usage: 116 KB
//
//	  Last 10 elements: [...]
//
//	  > Memory Usage: 118 KB
//
//	    + In the expected output, `main()` releases the memory.
//
//	      It no longer uses 65 MB of memory. Instead, it only
//	      uses 118 KB of memory. That's why the second
//	      `api.Report()` call reports 118 KB.
//
//	ADDITIONAL NOTE
//
//	  Memory leak means: Your program is using unnecessary
//	  computer memory. It doesn't release memory that is
//	  no longer needed.
//
//	  See this for more information:
//	  https://en.wikipedia.org/wiki/Memory_leak
func TestFixMemoryLeak(t *testing.T) {
	// Arrange
	reportMemoryUsage("initial memory usage")

	// Act
	millions := readLargeSlice()

	reportMemoryUsage("after reading first slice")

	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪
	last10 := millions[len(millions)-10:]
	// I override the slice header pointer to a new backing array that will only have a cap 10 - therefore having Go garbage collecting the backing array with millions of items and 65mb of memory
	millions = append([]int{}, last10...)
	fmt.Printf("\nLast 10 elements: %d\n\n", last10)
	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪

	reportMemoryUsage("after reading large slice")

	// don't worry about this code.
	fmt.Fprintln(io.Discard, millions[0])
}

// readLargeSlice returns a huge slice (allocates ~65 MB of memory)
func readLargeSlice() []int {
	// 2 << 22 means 2^(22 + 1)
	// See this: https://en.wikipedia.org/wiki/Arithmetic_shift

	// Perm function returns a slice with random integers in it.
	// Here it returns a slice with random integers that contains
	// 8,388,608 elements. One int value is 8 bytes.
	// So: 8,388,608 * 8 = ~65MB
	return rand.Perm(2 << 22)
}

// reportMemoryUsage cleans the memory and prints the current memory usage
func reportMemoryUsage(msg string) {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s] > Memory Usage: %v KB\n", msg, m.Alloc/1024)
}
