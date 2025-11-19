package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Square
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle (pakai base, height, dan tiga sisi untuk keliling)
type Triangle struct {
	Base   float64
	Height float64
	A      float64 // sisi 1
	B      float64 // sisi 2
	C      float64 // sisi 3
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Trapezoid (trapesium) dengan dua basis dan dua sisi miring serta tinggi
type Trapezoid struct {
	BaseA  float64
	BaseB  float64
	SideC  float64
	SideD  float64
	Height float64
}

func (t Trapezoid) Area() float64 {
	return 0.5 * (t.BaseA + t.BaseB) * t.Height
}

func (t Trapezoid) Perimeter() float64 {
	return t.BaseA + t.BaseB + t.SideC + t.SideD
}

// ProcessShape: cetak nama, luas, keliling
func ProcessShape(sh Shape) {
	var name string
	switch sh.(type) {
	case Square:
		name = "Square"
	case Rectangle:
		name = "Rectangle"
	case Triangle:
		name = "Triangle"
	case Circle:
		name = "Circle"
	case Trapezoid:
		name = "Trapezoid"
	default:
		name = "Unknown"
	}
	fmt.Printf("%s -> Area: %.2f, Perimeter: %.2f\n", name, sh.Area(), sh.Perimeter())
}

func main() {
	sq := Square{Side: 5}
	rec := Rectangle{Width: 4, Height: 6}
	tri := Triangle{Base: 6, Height: 4, A: 5, B: 5, C: 6}
	cir := Circle{Radius: 3}
	trp := Trapezoid{BaseA: 8, BaseB: 5, SideC: 4, SideD: 4, Height: 3}

	ProcessShape(sq)
	ProcessShape(rec)
	ProcessShape(tri)
	ProcessShape(cir)
	ProcessShape(trp)
}
