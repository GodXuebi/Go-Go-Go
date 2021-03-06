package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = 20
	return r.width * r.height
}

func (r rect) perim() int {
	return (r.width + r.height) * 2
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim", rp.perim())
}
