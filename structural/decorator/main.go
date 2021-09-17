package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Area float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle with area %f", c.Area)
}

func (c *Circle) Resize(factor float32) {
	c.Area *= factor
}

type Square struct {
	Area float32
}

func (c *Square) Render() string {
	return fmt.Sprintf("Square with area %f", c.Area)
}

type ColorShape struct {
	Shape Shape
	Color string
}

func (c *ColorShape) Render() string {
	return fmt.Sprintf("%s has %s color", c.Shape.Render(), c.Color)
}

type BorderShape struct {
	Shape  Shape
	Border float32
}

func (c *BorderShape) Render() string {
	return fmt.Sprintf("%s has %fpx border", c.Shape.Render(), c.Border)
}

func main() {
	circle := Circle{
		Area: 10,
	}
	fmt.Println(circle.Render())

	coloredCircle := ColorShape{
		Shape: &circle,
		Color: "black",
	}
	fmt.Println(coloredCircle.Render())

	coloredBorderCircle := BorderShape{
		Shape:  &coloredCircle,
		Border: 13,
	}
	fmt.Println(coloredBorderCircle.Render())
}
