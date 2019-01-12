package lib

import "fmt"

func DataStructure() {

	fmt.Println("\n\n ## The following content is exploring Go Data Structures ## \n")

	// Arrays
	var arrayVar [3]float64
	arrayVar[0] = -0.1
	arrayVar[1] = 0.0
	arrayVar[2] = 0.1
	fmt.Println("array:", arrayVar)

	array2Var := [5]int{5, 4, 3, 2, 1} // array has defined size
	fmt.Println("array:", array2Var)

	fmt.Println("slice[1:2]:", arrayVar[1:2])
	fmt.Println("slice[:2]:", arrayVar[:2])
	fmt.Println("slice[1:]:", arrayVar[1:])

	// Slices
	slice1 := []int{5, 4, 3, 2, 1} // slice has not defined size
	fmt.Println("slice1:", slice1)

	slice2 := make([]int, 5, 10)
	fmt.Println("slice2:", slice2)

	copy(slice2, slice1)
	fmt.Println("copy(slice2, slice1)", slice2)

	slice2 = append(slice2, 0, -1, -2, -3, -4)
	fmt.Println("append(slice2, 0, -1)", slice2)

	// MAPs
	sampleMap := make(map[string]int)
	sampleMap["address1"] = 100
	fmt.Println("sampleMap['address1']:", sampleMap["address1"])

	sampleMap["address2"] = 101
	fmt.Println(sampleMap)

	// Structs
	triangle := Triangle{height: 10, width: 10}
	fmt.Printf("Struct triangle type:%T, value:%v: \n", triangle, triangle)
	fmt.Println("Struct triangle area:", triangle.area())
	fmt.Println("Struct triangle area 2:", triangle.area())

	rectangle := Rectangle{height: 10, width: 10}
	circle := Circle{radius: 1}

	fmt.Println("Interface Rectangle Area:", getArea(rectangle))
	fmt.Println("Interface Rectangle Area 2:", getArea(rectangle))
	fmt.Println("Interface Circle Area:", getArea(circle))

}

type Rectangle struct {
	height float64
	width  float64
}

type Triangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Rectangle) area() float64 {
	return s.width * s.height
}

func (s *Triangle) area() float64 {
	s.width = s.width - 1
	s.height = s.height - 1
	return (s.width * s.height) / 2
}

func (s Circle) area() float64 {
	const pi float64 = 3.1415
	return pi * (s.radius * s.radius)
}

type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}
