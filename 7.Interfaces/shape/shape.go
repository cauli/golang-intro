package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Color struct {
	r int16 // 0 to 255
	g int16
	b int16
}

type Circle struct {
	radius float64
	color  Color
}

type Triangle struct {
	height float64
	base   float64
	color  Color
}

type Square struct {
	side  float64
	color Color
}

func PaintWithInterface() {

	red := Color{255, 0, 0}
	blue := Color{r: 0, g: 0, b: 255}
	green := Color{g: 255}

	c := Circle{radius: 10.5, color: green}

	t := Triangle{}
	t.color = red
	t.height = 20
	t.base = 13

	s := Square{}
	square := &s
	square.color = blue
	square.side = 10

	fmt.Println("Measuring areas...")
	fmt.Println("Square: ", measure(s))
	fmt.Println("Triangle: ", measure(t))
	fmt.Println("Circle: ", measure(c))

	// measureAll(c, t, s)
}

// Yay for Variadic Funcitions
// func measureAll(shapes ...Shape) {
// 	for _, shape := range shapes {
// 		fmt.Println("Shape area: ", measure(shape))
// 	}
// }

func measure(shape Shape) float64 {
	return shape.area()
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (t Triangle) area() float64 {
	return (t.height * t.base) / 2
}
