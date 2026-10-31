package main

// type CustomMap[T comparable, V int | string] map[T]V
//
// func main() {
// 	m := make(CustomMap[int, string])
// 	m[3] = "3"
// }

// import "fmt"
//
// type CustomData interface {
// 	int | float64 | string
// }
//
// type User[T CustomData] struct {
// 	ID   int
// 	Name string
// 	Data T
// }
//
// func main() {
// 	u := User[string]{
// 		ID:   0,
// 		Name: "",
// 		Data: "abc",
// 	}
// 	fmt.Printf("user: %+v\n", u)
// }

// import "fmt"

// input: [1, 2, 3], (n) => n * 2
// output: [2, 4, 6]
//
// func MapValues[T int | float64](values []T, mapFunc func(T) T) []T {
// 	var newValues []T
// 	for _, v := range values {
// 		newValue := mapFunc(v)
// 		newValues = append(newValues, newValue)
// 	}
// 	return newValues
// }
//
// func main() {
// 	result := MapValues([]float64{1.1, 2.1, 3.1}, func(n float64) float64 {
// 		return n * 2
// 	})
// 	fmt.Printf("result: %+v\n", result)
// }

// import "fmt"
//
// type UserID int
//
// func Add[T ~int | float64](a T, b T) T {
// 	return a + b
// }
//
// func main() {
// 	a := UserID(1)
// 	b := UserID(2)
// 	result := Add(a, b)
// 	fmt.Printf("result: %+v\n", result)
// }
