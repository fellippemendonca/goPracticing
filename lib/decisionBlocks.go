package lib

import "fmt"

const inputInt = 5;

func DecisionBlocks() {

	fmt.Println("if/else if/else");
	
	// ----------------------- IF -----------------------
	if inputInt > 5 {
		fmt.Println(inputInt, "is bigger than", 5);
	} else if inputInt == 5 {
		fmt.Println(inputInt, "is equal to", 5);
	} else {
		fmt.Println(inputInt, "is smaller than", 5);
	}
	
	fmt.Println("switch");
	// ----------------------- SWITCH -----------------------
	switch inputInt {
		case 4 : fmt.Println("integer is four");
		case 5 : fmt.Println("integer is five");
		default: fmt.Println("integer is something else");
	}

}