package main

import (
	"errors"
	"fmt"
)

var global string = "Global Variable"

func main() {
	// loops()
	// basics()
	// types()
	// arrays()
	// ranges()
	// maps()
	// methods()
	errorCode()
}

func errorExample() (int, error) {
	return 0, errors.New("There is an error")
}

func errorCode() {
	_, err := errorExample()
	if err != nil {
		fmt.Println(err)
	}
}

// type PointInterface interface {
// 	Hello()
// }
//
// type Point struct {
// 	X int
// 	Y int
// }
//
// func (p *Point) Hello() {
// 	fmt.Println("The point is saying hello")
// }
//
// // Equivalent to
// // func sum(p Point) int {}
// // However this allows an instance to call it like Point{3,4}.sum()
// // Also works on general types
// func (p *Point) sum() int {
// 	p.X = 4
// 	return p.X + p.Y
// }
//
// type MyType float64
//
// func (f MyType) hi() {
// 	fmt.Println("Hi")
// }
//
// func methods() {
// 	x := Point{3, 4}
// 	total := x.sum()
// 	fmt.Println(total)
// 	fmt.Println(x.X)
//
// 	y := MyType(3.0)
// 	y.hi()
//
// 	var z PointInterface = Point{4, 5}
// 	z.Hello()
// }
//
// type Tuple struct {
// 	X int
// 	Y int
// }
//
// func maps() {
// 	a := map[string]int{}
// 	a["a"] = 1
// 	// b is used if you already know the relative size
// 	b := make(map[string]int, 3)
// 	b["b"] = 1
// 	fmt.Println(a)
// 	fmt.Println(b)
//
// 	// These are equivalent
// 	// c := map[int]Tuple{
// 	// 	1: Tuple{0, 0},
// 	// 	2: Tuple{X: 2, Y: 1},
// 	// }
// 	c := map[int]Tuple{
// 		1: {0, 0},
// 		2: {2, 1},
// 	}
// 	fmt.Println(c)
// 	c[3] = Tuple{3, 3}
// 	fmt.Println(c)
// 	d := c[1]
// 	fmt.Println(d)
// 	delete(c, 2)
// 	fmt.Println(c)
// 	delete(c, 2) // If the key does not exist it does nothing
// 	fmt.Println(c)
//
// 	// The second variable determines whether the value exist
// 	e, ok := c[1]
// 	fmt.Println(e, ok)
// 	f, ok := c[9]
// 	fmt.Println(f, ok)
//
// 	g := addition
// 	res := g(2, 4)
// 	fmt.Println(res)
//
// 	h := func(a int, b int) int {
// 		return a + b
// 	}
// 	newRes := h(3, 4)
// 	fmt.Println(newRes)
//
// 	// closure is a function that returns a function that returns int (func() int)
// 	closure := func() func() int {
// 		sum := 1
// 		return func() int {
// 			sum += 1
// 			return sum
// 		}
// 	}
//
// 	// now refers to the same underlying sum variable (like an attribute)
// 	now := closure()
// 	fmt.Println(now())
// 	fmt.Println(now())
// }
//
// func addition(a int, b int) int {
// 	return a + b
// }
//
// func ranges() {
// 	a := []string{"a", "b", "c"}
// 	for i, v := range a {
// 		fmt.Println(i, v)
// 	}
// 	for _, v := range a {
// 		fmt.Println("n/a", v)
// 	}
// 	// Only gives the index
// 	for i := range a {
// 		fmt.Println(i, "n/a")
// 	}
// }
//
// func arrays() {
// 	var a [2]string
// 	a[0] = "First"
// 	a[1] = "Second"
// 	// a[2] is out of bound
// 	fmt.Println(a)
//
// 	b := [3]int{1, 2, 3}
// 	fmt.Println(b)
//
// 	// Same as a[1:]
// 	aSlice := a[1:2]
// 	fmt.Println(aSlice)
// 	aSlice[0] = "Nothing"
// 	fmt.Println(a)
//
// 	// Array literal
// 	c := [3]bool{true, true, false}
// 	fmt.Println(c)
// 	fmt.Printf("len=%d cap=%d %v\n", len(c), cap(c), c)
// 	// Array slide
// 	d := []bool{true, true, false}
// 	fmt.Println(d)
// 	fmt.Printf("len=%d cap=%d %v\n", len(d), cap(d), d)
//
// 	// Used if you want to define the len and cap
// 	// e := []int{}
// 	e := make([]int, 5, 5)
// 	fmt.Printf("len=%d cap=%d %v\n", len(e), cap(e), e)
// 	f := append(e, 1)
// 	fmt.Printf("len=%d cap=%d %v\n", len(e), cap(e), e)
// 	fmt.Printf("len=%d cap=%d %v\n", len(f), cap(f), f)
// }
//
// type Example struct {
// 	X int
// 	Y int
// }
//
// func types() {
// 	i := 1
// 	var p *int = &i
// 	q := p
//
// 	fmt.Println(p, q)
// 	fmt.Println(*p, *q)
// 	*q = 2
// 	fmt.Println(*p, *q)
//
// 	// example := Example{1, 2}
// 	example := Example{X: 3, Y: 4}
// 	fmt.Println(example.X)
// 	testing(&example)
// 	fmt.Println(example.X)
// }
//
// func testing(e *Example) {
// 	// Shorthand for (*e).X
// 	e.X = 43
// }
//
// func loops() {
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(i)
// 	// }
//
// 	// sum := 1
// 	// for sum < 5 {
// 	// 	sum += sum
// 	// 	fmt.Println(sum)
// 	// }
//
// 	// Creates an infinite loop
// 	// for {
// 	// 	fmt.Println("Infinite Loop")
// 	// }
//
// 	// x := 0
// 	// if x == 1 {
// 	// 	fmt.Println("x is equal to 1")
// 	// } else if x == 0 {
// 	// 	fmt.Println("x is 0")
// 	// } else {
// 	// 	fmt.Println("x is not equal to 1 or 0")
// 	// }
// 	//
// 	// y := 0
// 	// switch y {
// 	// case 0:
// 	// 	fmt.Println("Y is 0")
// 	// case f():
// 	// 	fmt.Println("Something")
// 	// case 1:
// 	// 	test := "Hello"
// 	// 	fmt.Println(test)
// 	// default:
// 	// 	fmt.Println("Nothing")
// 	// }
//
// 	// Equivalent to switch true {}
// 	// switch {
// 	// case 1 == 0:
// 	// 	fmt.Println("They are not equal")
// 	// case 1 == 1:
// 	// 	fmt.Println("They are equal")
// 	// default:
// 	// 	fmt.Println("Something")
// 	// }
//
// 	// defer fmt.Println("First")
// 	// fmt.Println("Second")
// 	//
// 	// for j := 0; j < 3; j++ {
// 	// 	defer fmt.Println(j)
// 	// }
// }
//
// func f() int {
// 	return 0
// }
//
// func basics() {
// 	fmt.Println("Hello World!")
// 	total := add(1, 2)
// 	abc := float64(total)
// 	fmt.Printf("Something %T\n", abc)
//
// 	a, b := swap("Hello", "World")
// 	fmt.Println(a, b)
//
// 	fmt.Println(global)
//
// 	const kjdf int = 40
// 	fmt.Println(kjdf)
// }
//
// func add(a int, b int) int {
// 	return a + b
// }
//
// func swap(a string, b string) (string, string) {
// 	return b, a
// }
