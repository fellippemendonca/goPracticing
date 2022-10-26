package examples

import "fmt"

func Functions() {

	fmt.Println("\n\n ## The following content is exploring Go Functions ## \n")

	// defer executes de function only after all functions inside this scope are executed;
	defer fmt.Println("squareFunc(2):", squareFunc(2))
	num1, num2 := next2ValuesFunc(1)
	fmt.Println("next2ValuesFunc(1):", num1, num2)
	fmt.Println("sumThemUpFunc([]float64{1, 2, 3, 4}):", sumThemUpFunc([]float64{1, 2, 3, 4}))
	fmt.Println("subThemUpFunc(1, 2, 3, 4, 5):", subThemUpFunc(1, 2, 3, 4, 5))
	fmt.Println("factorial(5):", factorial(5))

	// Anonymous Functions
	someVal := 2
	anonymousFunc := func() int {
		return someVal * 3
	}
	fmt.Println("anonymousFunc():", anonymousFunc())

	fmt.Println("div 3 by 0 =", safeDiv(3, 0))
	fmt.Println("div 3 by 3 =", safeDiv(3, 3))

	demPanic()

}

// simple input/output
func squareFunc(input float64) float64 {
	return input * input
}

// Multiple outputs
func next2ValuesFunc(input float64) (float64, float64) {
	return input + 1, input + 2
}

// Multiple inputs
func subThemUpFunc(args ...float64) float64 {

	var sub float64

	for _, value := range args {
		sub -= value
	}

	return sub
}

//array reduce
func sumThemUpFunc(input []float64) float64 {

	var sum float64

	for _, value := range input {
		sum += value
	}

	return sum
}

// Recursive
func factorial(number float64) float64 {

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

// Recover
func safeDiv(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
		// or just recover();
	}()

	result := num1 / num2
	return result

}

// Panic
func demPanic() {

	defer func() {
		fmt.Println(recover())
		// or just recover();
	}()

	panic("PANIC")
}
