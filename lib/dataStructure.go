package lib

import "fmt"


func DataStructure() {

	var byteVar byte = 1;
	fmt.Println("byte:", byteVar);

	var booleanVar bool = true;
	fmt.Println("bool:", booleanVar);

	var stringVar string = "Hello World";
	fmt.Println("string:", stringVar);

	var uint8Var uint8 = 0;
	fmt.Println("uint8-64:", uint8Var);

	var int8Var int8 = -128;
	fmt.Println("int8-64:", int8Var);

	var float32Var float32 = -2147483648.2147483648;
	fmt.Println("float32-64:", float32Var);
	
	var complex64Var complex64 = -9223372036854775808.1;
	fmt.Println("complex64-128:", complex64Var);


	// Arrays
	var arrayVar[3] float64;
	arrayVar[0] = -0.1;
	arrayVar[1] = 0.0;
	arrayVar[2] = 0.1;
	fmt.Println("array:", arrayVar);

	array2Var := [5]int {5, 4, 3, 2, 1}; // array has defined size
	fmt.Println("array:", array2Var);

	fmt.Println("slice[1:2]:", arrayVar[1:2]);
	fmt.Println("slice[:2]:", arrayVar[:2]);
	fmt.Println("slice[1:]:", arrayVar[1:]);

	// Slices
	slice1 := []int {5, 4, 3, 2, 1}; // slice has not defined size
	fmt.Println("slice1:", slice1);

	slice2 := make([]int, 5, 10);
	fmt.Println("slice2:", slice2);
	
	copy(slice2, slice1);
	fmt.Println("copy(slice2, slice1)", slice2);

	slice2 = append(slice2, 0, -1, -2, -3, -4);
	fmt.Println("append(slice2, 0, -1)", slice2);

	// MAPs
	sampleMap := make(map[string] int);
	sampleMap["address1"] = 100;
	fmt.Println("sampleMap['address1']:", sampleMap["address1"]);

	sampleMap["address2"] = 101;
	fmt.Println(sampleMap);


	// Pointers
	x := 0;
	fmt.Println("Value of x is:", x);
	fmt.Println("Address of x is:", &x);
	
	y := &x;
	fmt.Printf("y := &x, so y = %T\n", y);

	anonymous := func(num *int) {
		*num++;
	};
	
	anonymous(y);
	
	fmt.Println("new value of x:", x);

}