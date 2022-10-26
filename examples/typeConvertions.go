package examples

import (
	"fmt"
	"strconv"
)

func TypeConvertions() {

	fmt.Println("\n\n ## The following content is exploring Go Type Convertions ## \n")

	sampleInt := 5
	sampleFloat := 10.5
	sampleString := "7"
	sampleString2 := "7.2"

	fmt.Println("int(sampleFloat) =", int(sampleFloat))
	fmt.Println("float64(sampleInt) =", float64(sampleInt))

	newInt, _ := strconv.ParseInt(sampleString, 0, 64)
	fmt.Println("strconv.ParseInt(sampleString, 0, 64) =", newInt)

	newFloat, _ := strconv.ParseFloat(sampleString2, 64)
	fmt.Println("strconv.ParseFloat(sampleString2, 0, 64) =", newFloat)

}
