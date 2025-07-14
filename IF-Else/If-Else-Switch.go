package main

import "fmt"

func main() {
	// 1. Basic If-Elseif-Else
	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// 2.Short Statement with if You can initialize a variable inside the if itself:
	if x := 5 * 2; x > 5 {
		fmt.Println("x is greater than 5")
	}

	// 🚫 This will NOT compile in Go
	// if (z = 5) { ... }  ❌
	// ✔️ Go requires a **boolean expression** in `if`:
	z := 5
	if z == 5 {
		fmt.Println("This is correct way")
	}


}
