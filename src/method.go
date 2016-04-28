package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

// func (r +rect) perim() int {
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 5, height: 10}
	fmt.Println("area=", r.area())
	fmt.Println("perim=", r.perim())

}
