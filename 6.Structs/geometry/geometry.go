package geometry

import "fmt"
import "strconv"

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

// When you do not specify a name you are composing directly!
// type SquareColor struct {
// 	side float64
// 	Color
// }

func CreateShapesAndColors() {

	red := Color{255, 0, 0}
	blue := Color{r: 0, g: 0, b: 255}
	green := Color{g: 255}

	c := Circle{radius: 10.5, color: green}

	t := Triangle{}
	t.color = red
	t.height = 20
	t.base = 10

	s := Square{}
	square := &s
	square.color = blue
	square.side = 10

	fmt.Println("Square: ", s)
	fmt.Println("Triangle: ", t)
	fmt.Println("Circle: ", c)

	// scolor := SquareColor{}
	// scolor.r = 111
	// scolor.g = 111
	// scolor.b = 111
	// scolor.side = 10
}

func (c Circle) String() string {
	return "My radius is: " + strconv.FormatFloat(c.radius, 'f', 2, 64)
}
