package examples

import (
	"errors"
	"fmt"
)

const positive int = 1
const negative int = -1

func DecisionBlocks() {

	fmt.Println("\n\n ## The following content is exploring Go Decision Blocks ## \n")

	fmt.Println(ifElser(negative))
	fmt.Println(ifElser(positive))

	switcher(positive)

}

func positiveInts(value int) (int, error) {

	if value < 0 {
		return value, errors.New("int: Negative number")
	}

	return value, nil
}

func ifElser(value int) (string, error) {

	checkVal, err := positiveInts(value)

	if err != nil {
		return "negative value!", errors.New("negative value!")
	}

	if checkVal > 1 {
		return "is bigger than 1", nil
	} else if checkVal == 1 {
		return "is equal to 1", nil
	} /*else {
		return "is smaller than 1";
	}*/

	// This will work as last condition and don't need the final else.
	return "is smaller than 1", nil

}

func switcher(value int) string {
	fmt.Println("switch")

	switch value {
	case 4:
		fmt.Println("integer is four")
	case 5:
		fmt.Println("integer is five")
	default:
		fmt.Println("integer is something else")
	}

	return "done!"

}
