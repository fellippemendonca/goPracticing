package main;

import (
	"fmt";
	"github.com/fellippemendonca/goPracticing/lib"
)

func main() {

	// ----------------------- PRINT -----------------------
	fmt.Printf("hello\n");

	// ----------------------- VAR INIT -----------------------
	const pi float64 = 3.1415926535;
	const oneInteger = 1;
	const hiString = "hi";

	var something string = hiString;

	fmt.Printf("%v \n", something);
	fmt.Printf("%.3f \n", pi);
	fmt.Printf("%T \n", pi);
	fmt.Printf("\a");

	lib.LoopIterations();
	lib.DecisionBlocks();
	lib.DataStructure();
	lib.Functions();

}
