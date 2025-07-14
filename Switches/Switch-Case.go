package main

import "fmt"

func main() {

	//1. Normal Switch case with condition
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	//2. Switch with fallthrough, it forces the next case to run even if the current one matched.
	num := 1

	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two") //Only this will print after one due to fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	//3. Swith case without expression acts like else block
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 75:
		fmt.Println("Grade B")
	default:
		fmt.Println("Grade C")
	}

	//4. Multiple values in case
	ch := "a"
	switch ch {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

	//5. Type switch
	var val interface{} = 10.5

	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}

}
