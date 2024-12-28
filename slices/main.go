package main

import (
	"fmt"

	"github.com/inancgumus/prettyslice"
)

// A slice is a window to a backing array. It's like a lens that you can use to see a part of an array.
// A slice is a reference type. It doesn't store any data. It just points to an array. Therefore, when you change a slice, you change the array it points to.
func slicesWithBackingArray() {
	ages := [5]int{1, 2, 3, 4, 5}

	// We create a new slice - this will be referencing the same `ages` backing array
	newAge1 := ages[0:3]
	newAge2 := ages[1:4]

	fmt.Println("Before changing")
	prettyslice.Show("newAge1", newAge1)
	prettyslice.Show("newAge2", newAge2)

	ages[2] = 100

	fmt.Println("After changing")

	prettyslice.Show("newAge1", newAge1)
	prettyslice.Show("newAge2", newAge2)

}

// To prevent mutating the original backing the array, we can create a new slice with a new backing array. This can be done by appending the values of the slice into a new slice or by copying the values of the slice into a new slice.
func slicesWithAppendingfBackingArray() {
	ages := [5]int{1, 2, 3, 4, 5}

	var newAge1 = []int{}
	var newAge2 = []int{}
	// We append to the values of the slice into a new slice with a new backing array
	newAge1 = append(newAge1, ages[0:3]...)
	newAge2 = append(newAge2, ages[1:4]...)

	fmt.Println("Before changing")
	prettyslice.Show("newAge1", newAge1)
	prettyslice.Show("newAge2", newAge2)

	ages[2] = 100

	fmt.Println("After changing")

	prettyslice.Show("newAge1", newAge1)
	prettyslice.Show("newAge2", newAge2)
}

// To prevent mutating the original backing the array, we can create a new slice with a new backing array. This can be done by copying the values of the slice into a new slice.
func slicesWithNewCopyOfBackingArray() {
	ages := [5]int{1, 2, 3, 4, 5}

	// Create slices that reference the original array
	newAge1 := ages[0:3]
	newAge2 := ages[1:4]

	// Create copies of the slices to have separate backing arrays
	copyAge1 := make([]int, len(newAge1))
	copy(copyAge1, newAge1)

	copyAge2 := make([]int, len(newAge2))
	copy(copyAge2, newAge2)

	fmt.Println("Before changing")
	prettyslice.Show("newAge1", newAge1)
	prettyslice.Show("newAge2", newAge2)

	ages[2] = 100

	fmt.Println("After changing")

	prettyslice.Show("newAge1", newAge1)
	prettyslice.Show("newAge2", newAge2)
}

func changeArray(data [4]string) {
	// Array are passed by copies!
	data[2] = "X"
	prettyslice.Show("changeArray", data)
}

func changeSlice(data []string) {
	// Slices are passed by copies - but share the same backing array!
	// This wil be just a slice header referencing (and mutating) the same backing array
	data[2] = "X"
	prettyslice.Show("changeSlice", data)
}

// Slices are just headers that reference an array. When you pass a slice to a function, you're passing a copy of the slice header. The backing array is not copied. Therefore, when you change the slice, you change the backing array.
func sliceHeaders() {

	data := [4]string{"A", "B", "C", "D"}
	changeArray(data)
	prettyslice.Show("main array", data)

	dataSlice := []string{"A", "B", "C", "D"}
	changeSlice(dataSlice)
	prettyslice.Show("main slice", dataSlice)
}

// When you append to a slice, if the backing array is not large enough, Go will create a new backing array with a larger capacity, copy the elements from the original backing array to the new backing array, and then append the new elements. This is why the backing array of the original slice is different from the backing array of the extended slice.
func expandCapacity() {
	original := []int{1, 2, 3, 4, 5}
	slice := original[1:3]

	prettyslice.Show("original", original)
	prettyslice.Show("slice", slice)
	// extend capacity beyond the original slice
	extended := append(slice, []int{6, 7, 8, 9, 10}...)
	// show the backing array of extended - this is now different from the original
	prettyslice.Show("extended - after multiple appends", extended)
	extended = append(extended, 11)
	prettyslice.Show("extended - after 1 append", extended)
	extended = append(extended, 12)
	prettyslice.Show("extended - after 2 append", extended)

}

func main() {
	prettyslice.PrintBacking = true
	prettyslice.PrintElementAddr = true
	prettyslice.MaxPerLine = 20
	prettyslice.Width = 150
	// fmt.Println("*********** SLICES REFERENCING ORIGINAL BACKING ARRAY ***********")
	// slicesWithBackingArray()
	// fmt.Println("*********** SLICES COPIED USING APPEND ***********")
	// slicesWithAppendingfBackingArray()
	// fmt.Println("*********** SLICES COPIED USING COPY ***********")
	// slicesWithNewCopyOfBackingArray()
	// sliceHeaders()
	// expandCapacity()
}
