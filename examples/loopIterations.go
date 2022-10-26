package examples

import "fmt"

const times = 3

func LoopIterations() {

	fmt.Println("\n\n ## The following content is exploring Go Loop Iterations ## \n")

	fmt.Println("For iteration as while =", whileFor(times))

	fmt.Println("simple For Iteration =", simpleFor(times))

	arrayVar := []int{1, 2, 3}

	fmt.Println("for arrays iteration", arrayFor(arrayVar))

}

func whileFor(value int) int {

	sum := 0
	iterator := 0

	for iterator < value {
		sum += iterator
		iterator++
	}

	return sum
}

func simpleFor(value int) int {

	sum := 0

	for i := 0; i < value; i++ {
		sum += i
	}

	return sum
}

func arrayFor(value []int) int {

	sum := 0

	for key, value := range value {
		fmt.Println("array[", key, "] =", value)
		sum += value
	}

	for _, value := range value {
		fmt.Println("array[_] =", value)
		sum += value
	}

	for key, _ := range value {
		fmt.Println("array[", key, "] = ?")
		sum += key
	}

	return sum

}
