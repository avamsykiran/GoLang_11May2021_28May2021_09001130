package main

import "fmt"

type board interface {
	area() float32
	perimeter() float32
}

func estimatePaintingAndBorderingCost(b board) (float32, float32) {
	return b.area() * 5, b.perimeter() * 3
}

type rect_board struct {
	length  int
	breadth int
}

func (r rect_board) area() float32 {
	return float32(r.length * r.breadth)
}

func (r rect_board) perimeter() float32 {
	return float32(2 * (r.length + r.breadth))
}

type circular_board struct {
	radius float32
}

func (c circular_board) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (c circular_board) perimeter() float32 {
	return 2 * 3.14 * c.radius
}

func main() {
	r1 := rect_board{10, 7}
	c1 := circular_board{3.67}

	fmt.Println(estimatePaintingAndBorderingCost(r1))
	fmt.Println(estimatePaintingAndBorderingCost(c1))
}
