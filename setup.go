package main

import (
	"fmt"
	"os"
	"setup/myutil"
	"bufio"
)

func main() {
	fmt.Println("Hello world")

	myutil.PrintMessage("Hey from main")

	// Note: 1 - If want to use some variable in other packages Capatailze its first letter

	// Note: 2 - Name := "Sayoun"  , This is short hand way con=mpiler autmatically assigns
	// corresponding data type generally used when data of function return needs to be stored in a variable

	// Variables
	var Name string = "Sayoun"
	Name = "Sayoun01" // reassignment
	fmt.Println("My name is : ", Name)

	fullName := "Sayoun Parui"
	fmt.Println("My full name is : ", fullName)

	const constantName string = "Some name"
	// constantName = "Changed name" - This is not allowed in constants
	fmt.Println(constantName)

	// Note: 3 - Println() , it adds new line as well as a sinle space before and after variables
	address := "Varanasi"
	age := 20
	height := 5.9223
	fmt.Println("Address:", address, "Age:", age, "Income:", height)

	var income float64 = 45000.23421
	// Note: 4 - Printf() , it is format specifier same as C
	fmt.Printf("The addess is %s\nThe age is %d\nThe income is %.2f\n", address, age, income)

	//Input: 2 Ways
	// 1st fmt.Scan() till whitespace
	var email string
	fmt.Scan(&email)
	fmt.Println("The input email is : ", email)
	// 2nd reader = buffio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	fullAddress,_ := reader.ReadString('\n')
	fmt.Println("The full address is:",fullAddress);

	


}
