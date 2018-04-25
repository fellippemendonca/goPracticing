package lib

import "fmt"

const times = 3;

func LoopIterations() {


	fmt.Println("\n\n ## The following content is exploring Go Loop Iterations ## \n");

	fmt.Println("simplest for iteration");
	
	var iterator int = 0;

	for iterator < times {
		fmt.Println(iterator);
		iterator++
	}


	fmt.Println("standard for iteration");

	for i := 0; i < times; i++ {
		fmt.Println(i);
	}


	fmt.Println("for arrays iteration");

	arrayVar := [3]float64 {-0.2, 0, 0.2};

	for _, value := range arrayVar {
		fmt.Println(value);
	}

}