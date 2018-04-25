package lib;

import (
	"fmt";
	"time";
	"strconv";
)

var pizzaNum = 0;

func Routines() {

	fmt.Println("\n\n ## The following content is exploring Go Routines ## \n");
	//count(10);

	sauceChannel := make(chan string);
	toppingsChannel := make(chan string);
	shippingChannel := make(chan string);
	
	
	for i := 0; i < 10; i++ {

		go makeDough(sauceChannel);
		go addSauce(sauceChannel, toppingsChannel);
		go addToppings(toppingsChannel, shippingChannel);
		go shipping(shippingChannel);
		time.Sleep(time.Millisecond * 1000);
	
	}

}


func makeDough(sauceChannel chan string) {
		
	pizzaNum++;
	
	pizza := "Pizza #" + strconv.Itoa(pizzaNum);

	fmt.Println("Make Dough of", pizza);
	
	time.Sleep(time.Millisecond * 100);

	sauceChannel <- pizza;

}

func addSauce(sauceChannel, toppingsChannel chan string) {
	
	pizza := <- sauceChannel;

	fmt.Println("Add Sauce to", pizza);

	time.Sleep(time.Millisecond * 100);

	toppingsChannel <- pizza;
	
}


func addToppings(toppingsChannel, shippingChannel chan string) {
	
	pizza := <- toppingsChannel;

	fmt.Println("Add Toppings to", pizza);

	time.Sleep(time.Millisecond * 100);

	shippingChannel <- pizza;
	
}


func shipping(shippingChannel chan string) {
	
	pizza := <- shippingChannel;

	fmt.Println("Shipping the", pizza);
	
	time.Sleep(time.Millisecond * 100);
	
}


/*
func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i);
		time.Sleep(time.Millisecond * 1000);
	}
}
*/