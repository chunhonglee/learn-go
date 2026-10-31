package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

// Geometry is a composition interface
type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func describeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())
}

// type Circle struct {
// 	radius float64
// }
//
// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func calculateArea(s Shape) float64 {
// 	return s.Area()
// }

// func main() {
// 	rect := Rectangle{width: 5, height: 4}
// 	describeShape(rect)
//
// 	rect := Rectangle{width: 5, height: 4}
// 	circle := Circle{radius: 2}
//
// 	fmt.Println("Rectangle area", calculateArea(rect))
// 	fmt.Println("Circle area", calculateArea(circle))
//
// 	// Unnecessary since generics exist
// 	mysterBox := interface{}("")
// 	describeValue(mysterBox)
//
// 	retrievedInt, ok := mysterBox.(int)
// 	if ok {
// 		fmt.Println("Retrieved int:", retrievedInt)
// 	} else {
// 		fmt.Println("Value is not an integer")
// 	}
// }
//
// Unnecessary since generics exist
// func describeValue(t interface{}) {
// 	fmt.Printf("Type: %T, Value: %v\n", t, t)
// }

type CalculationError struct {
	msg string
}

// Error function implements the underlying interface of an error type
func (ce CalculationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{msg: "Invalid input"}
	}

	return math.Sqrt(val), nil
}
