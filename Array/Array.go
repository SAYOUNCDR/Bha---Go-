package main

import "fmt"

func main() {
	var names [10]string
	names[0] = "Sayoun"
	names[2] = "John"

	fmt.Println("Name of blank space :", names[3])            //return "" Output: [Sayoun  John       ]
	fmt.Println("Names in array are :", names)                //Println outputs the name in Simple way not formatted [Sayoun  John       ]
	fmt.Printf("Names in array in formatted way %q\n", names) // %q formats in double-quoted string ["Sayoun" "" "John" "" "" "" "" "" "" ""]
	//Initializing

	var nums = [9]int{1, 2, 3, 4} //empty places will be filled by 0 ->[1 2 3 4 0 0 0 0 0]
	fmt.Println("Numers in array are:", nums)

	fmt.Println("Length of nums array:", len(nums)) //Length of array

	//Initialised with zero/flase,"" acc to its datatype if not assigned any value like for pointer and complex its nil
	var emptyAarray [4]int
	fmt.Println("Empty array:", emptyAarray)
}
