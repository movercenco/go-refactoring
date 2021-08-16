package main

import "fmt"

type color interface {
	Color() string
}

type red struct {
	color string
}

func (r *red) Color() string {
	return r.color
}

func NewRed() *red {
	return &red{color: "red"}
}

type blue struct {
	color string
}

func NewBlue() *blue {
	return &blue{color: "blue"}
}

func (r *blue) Color() string {
	return r.color
}

type shape interface {
	Draw()
}

type Circle struct {
	color color
}

func NewCircle(color color) *Circle {
	return &Circle{color: color}
}

func (c Circle) Draw() {
	fmt.Println("Draw circle of color", c.color.Color())
}

type Square struct {
	color color
}

func (s Square) Draw() {
	fmt.Println("Draw square of color", s.color.Color())
}

func NewSquare(color color) *Square {
	return &Square{color: color}
}

func main() {
	red := NewRed()
	blue := NewBlue()

	circle := NewCircle(red)
	circle.Draw()

	square := NewSquare(blue)
	square.Draw()
}
