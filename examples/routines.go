package examples

import (
	"fmt"
	"time"
)

func Routines() {

	fmt.Println("## The following content is exploring Go Routines ##")

	restaurant := NewRestaurant()
	go restaurant.Kitchen()
	go restaurant.Delivey()

	restaurant.Ordering("Steak")
	// restaurant.Ordering("Pizza")
	// restaurant.Ordering("Burger")
	meal := <-restaurant.CloseCh
	fmt.Println("Delivered the", meal)
}

type Restaurant struct {
	PreparingCh chan string
	ShippingCh  chan string
	CloseCh     chan string
}

func NewRestaurant() Restaurant {
	return Restaurant{
		PreparingCh: make(chan string),
		ShippingCh:  make(chan string),
		CloseCh:     make(chan string),
	}
}

func (r *Restaurant) Ordering(meal string) {
	time.Sleep(time.Millisecond * 1000)
	r.PreparingCh <- meal
}

func (r *Restaurant) Kitchen() {
	meal := <-r.PreparingCh
	fmt.Println("Preparing the", meal)
	time.Sleep(time.Millisecond * 3000)
	r.ShippingCh <- meal
}

func (r *Restaurant) Delivey() {
	meal := <-r.ShippingCh
	fmt.Println("Shipping the", meal)
	time.Sleep(time.Millisecond * 2000)
	r.CloseCh <- meal
}
