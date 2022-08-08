package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
	volume() float64
	surfaceArea() float64
}
type square struct {
	lenght float64
}

type circle struct {
	radius float64
}

type trapezoid struct {
	side1, side2, width1, width2, height float64
}
type cylinder struct {
	radius, height float64
}

type cone struct {
	radius, height float64
}

func (s square) perimeter() float64 {
	return 4 * s.lenght
}

func (s square) area() float64 {
	return math.Pow(s.lenght, 2)
}
func (s square) volume() float64 {
	//TODO implement me
	panic("implement me")
}

func (s square) surfaceArea() float64 {
	//TODO implement me
	panic("implement me")
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	//TODO implement me
	panic("implement me")
}

func (c circle) volume() float64 {
	//TODO implement me
	panic("implement me")
}

func (c circle) surfaceArea() float64 {
	//TODO implement me
	panic("implement me")
}

func (t trapezoid) perimeter() float64 {
	return t.side1 + t.side2 + t.width1 + t.width2
}

func (t trapezoid) area() float64 {
	return (0.5) * (t.height) * (t.width2 + t.width1)
}
func (t trapezoid) volume() float64 {
	//TODO implement me
	panic("implement me")
}

func (t trapezoid) surfaceArea() float64 {
	//TODO implement me
	panic("implement me")
}

func main() {
	var s, c, t shape
	s = square{7.0}
	c = circle{6.5}
	t = trapezoid{4, 6, 3, 8, 5.5}
	fmt.Println("square :, Area: ", s.area(), " Perimeter: ", s.perimeter())
	fmt.Println("Circle :, Area: ", c.area())
	fmt.Println("Trapezoid:, Area:", t.area(), "Perimeter: ", t.perimeter())

}
