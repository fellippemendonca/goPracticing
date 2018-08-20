package lib

import "fmt"


func Variables() {

	fmt.Println("\n\n ## The following content is exploring Go Variables ## \n");

	const pi float64 = 3.1415926535;
	
	fmt.Printf("%.3f \n", pi);
	fmt.Printf("%T \n", pi);

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

	anyvar := "hello there...";
	fmt.Println("Dynamic var type casting:", anyvar);


	x := 0;
	fmt.Println("Value of x is:", x);
	fmt.Println("Address of x is:", &x);
	
	anonymous := func(num *int) *int {
		*num++;
		return num;
	};

	z := anonymous(&x);
	fmt.Printf("new value of x: %v, and value of z: %v\n", x, *z);


	fmt.Println(x<<8 + *z<<16);

}
