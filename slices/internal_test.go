package main

import (
	"fmt"
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
