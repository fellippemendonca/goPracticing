package examples

import (
	"fmt"
	"sort"
	"strings"
)

func StringOperations() {

	fmt.Println("\n\n ## The following content is exploring Go String Operations ## \n")

	sample := "Hello World"

	fmt.Println("sample=", sample)
	fmt.Println("strings.Contains(sample, 'lo') =", strings.Contains(sample, "lo"))
	fmt.Println("strings.Index(sample, 'lo') =", strings.Index(sample, "lo"))
	fmt.Println("strings.Count(sample, 'l') =", strings.Count(sample, "l"))
	fmt.Println("strings.Replace(sample, 'l') =", strings.Replace(sample, "l", "x", 2))
	fmt.Println("sample=", sample)

	csvSample := "1,2,3,4,5"
	fmt.Println("sample =", csvSample)
	fmt.Println("strings.Split(csvSample, ',') =", strings.Split(csvSample, ","))

	listOfLetters := []string{"c", "a", "b"}
	fmt.Println("listOfLetters =", listOfLetters)
	sort.Strings(listOfLetters)
	fmt.Println("sort.Strings(listOfLetters) =", listOfLetters)

	listOfNumbers := []string{"3", "2", "1"}
	fmt.Println("listOfNumbers =", listOfNumbers)
	fmt.Println("strings.Join(listOfNumbers, ', ') =", strings.Join(listOfNumbers, ", "))

}
