package main

import "fmt"

func main() {
	fmt.Println("Slices are better version of array with dynamic size")
	// Basic slice creation
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 1. Creating Slices
	// From Array
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]  // elements from index 1 to 3 (not 4)
	fmt.Println(slice) // [20 30 40]

	//Directly
	names := []string{"Go", "Rust", "Python"}
	fmt.Println(names)

	//Using make
	nums := make([]int, 5) // len = 5, cap = 5
	fmt.Println(nums)      // [0 0 0 0 0]

	// 2. Slice Properties: Length & Capacity
	s := make([]int, 1, 3) //make([]sataType, len, cap)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)    //NOTE: If inserted more than capacity then capacity is increased by double
	fmt.Println(s)      // [0, 1, 2, 3]
	fmt.Println(len(s)) // 4
	fmt.Println(cap(s)) // 6

	// 3. Appending to a Slice
	sl := []int{1, 2}
	sl = append(sl, 3, 4)
	fmt.Println(sl) // [1 2 3 4]

	// 4. Coping a Slice
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(b) // [1 2 3]

	// 5. Slicing a Slice sl := []int{1, 2, 3, 4}
	sub := sl[1:4] // from index 1 to 3
	fmt.Println(sub)

	// 6. Nil vs Empty Slice
	var ax []int  // nil slice (a == nil is true)
	bx := []int{} // empty slice (len=0, cap=0)
	fmt.Println(ax, bx)

	// 7.  7. Iterating over Slices

	It := []string{"apple", "banana", "cherry"}
	for i, o := range It {
		fmt.Println(i, o)
	}

}
