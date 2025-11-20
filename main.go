package main

import (
	"fmt"
	"reflect"
)

const PI = 3.14159

type Shape interface {
	Area() float64      // Menghitung Luas
	Perimeter() float64 //Menghitung Keliling
}

// Persegi
type Square struct {
	Side float64
}

func (square Square) Area() float64 {
	return square.Side * square.Side
}
func (square Square) Perimeter() float64 {
	return square.Side * 4
}

// Persegi Panjang
type Rectangle struct {
	Length, Width float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Length * rectangle.Width
}
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Length + rectangle.Width)
}

// Segitiga
type Triangle struct {
	Base, Height        float64
	Side1, Side2, Side3 float64
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Base * triangle.Height
}
func (triangle Triangle) Perimeter() float64 {
	return triangle.Side1 + triangle.Side2 + triangle.Side3
}

// Circle
type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return PI * (circle.Radius * circle.Radius)
}
func (circle Circle) Perimeter() float64 {
	return 2 * PI * circle.Radius
}

// Trapesium
type Trapezoid struct {
	Base1, Base2, Height float64
	Side1, Side2         float64
}

func (trapezoid Trapezoid) Area() float64 {
	return 0.5 * (trapezoid.Base1 + trapezoid.Base2) * trapezoid.Height
}
func (trapezoid Trapezoid) Perimeter() float64 {
	return trapezoid.Base1 + trapezoid.Base2 + trapezoid.Side1 + trapezoid.Side2
}

func main() {
	// Slice shapes dengan typedata Interface
	shapes := []Shape{
		Square{Side: 5},                 // Persegi
		Rectangle{Length: 10, Width: 5}, // Persegi Panjang
		Triangle{Base: 10, Height: 8, Side1: 10, Side2: 10, Side3: 10}, // Triangle
		Circle{Radius: 25}, //Lingkaran
		Trapezoid{Base1: 15, Base2: 10, Height: 8, Side1: 10, Side2: 10}, //Trapesium
	}

	for _, shape := range shapes {
		ProccessShape(shape)
	}

}

func ProccessShape(sh Shape) {
	name := reflect.TypeOf(sh).Name()
	fmt.Printf("Bangun Datar : %s\n", name)
	fmt.Printf("Luas : %.2f\n", sh.Area())
	fmt.Printf("Keliling : %.2f\n\n", sh.Perimeter())
}
