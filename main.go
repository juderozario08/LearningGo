package main

import (
	"fmt"
)

// This is a way to create a type that can accept any type of data
type Type interface{}

func main() {
	fmt.Println("################################################")
	fmt.Println()
	fmt.Println("Switch Statement Outputs")
	switches()
	fmt.Println()
	fmt.Println("Array Outputs")
	arrays()
	fmt.Println()
	fmt.Println("Slices Outputs")
	slices()
	fmt.Println()
	fmt.Println("Hashmap outputs")
	hashmap()
	fmt.Println()
	fmt.Println("Variadic Functions")
	variadic_functions(1, 2, "Hello World", 4, 5)
	fmt.Println()
	fmt.Println("Closures")
	close := closures()
	fmt.Println(close())
	fmt.Println(close())
	fmt.Println(close())
	fmt.Println()
	fmt.Println("Tail Recursion")
	fmt.Println(recursion(6))
	fmt.Println()
	fmt.Println("Pointers")
	ptrTest()
	fmt.Println()
	fmt.Println("Struct")
	structs()
}

func arrays() {
	// Initializing an array with values
	// 5:20 means that the 5th index will have the value 20
	// The rest of the values will be 0
	// ... means that the length of the array will be determined by the number of values
	arr := [...]int{5: 20, 5, 6, 7, 8}
	fmt.Println(arr[:6])
}

func slices() {
	// Slices have the option for appending values at the end unlike arrays
	// These are equivalent to array lists
	s := make([]int, 0)
	s = append(s, 2) // Appending a value to the slice
	s = append(s, 4)
	s = append(s, 5)
	fmt.Println(s)
	s = append(s, 10)
	fmt.Println(s)
	// Slices can be sliced like arrays
	fmt.Println(s[0:3])
}

func hashmap() {
	// This is a simple way how a map works
	kvs := make(map[int]string)
	// Assigning a value to a key
	kvs[20] = "Twenty"
	kvs[30] = "Thirty"
	// Iterating through the key value pairs using range
	// A range always returns an index and a value like an enum from python
	for k, v := range kvs {
		fmt.Printf("%d %s\n", k, v)
	}
	fmt.Println("Map: ", kvs)
	// Deleting a KV pair from the map
	delete(kvs, 20)
	fmt.Println("Deleted Map: ", kvs)
	clear(kvs)
	fmt.Println("Cleared Map: ", kvs)
}

// A way to create an anonymous function
func switches() {
	whatAmI := func(i Type /*Switch cases and the interface{} type*/) {
		// interface{} simply accepts any type of data
		switch t := i.(type) {
		case bool:
			fmt.Println("This is a bool")
		case string:
			fmt.Println("This is a string")
		case int:
			fmt.Println("This is a int")
			fmt.Println(t)
		default:
			fmt.Println("Don't know the type: ", t)
		}
	}
	whatAmI(20)
	whatAmI("String")
	whatAmI([]int{5, 6, 7})
}

// These type of functions simply accept the values as an array
func variadic_functions(nums ...Type) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		switch num := num.(type) {
		case int:
			total += num
		}
	}
	fmt.Println(total)

	// When passing an array of arguments to variadic_functions the type must be the same
	arr := []int{1, 2, 3, 4, 5}
	varTest := func(i ...int) {
		for _, num := range i {
			fmt.Print(num, " ")
		}
	}
	varTest(arr...) // This splits the array into individual values
	fmt.Println()
}

func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Basic tail recursion works the same as elixir
func recursion(num int) int {
	return recursionTail(num, 1)
}

func recursionTail(num int, acc int) int {
	if num <= 1 {
		return acc
	}
	return recursionTail(num-1, acc*num)
}

// Pointers work the same way as in C
func ptrTest() {
	a := 10
	b := 5
	zeroptr := func(ptr *int) {
		*ptr = 17
	}
	swap := func(a *int, b *int) {
		temp := *a
		*a = *b
		*b = temp
	}
	fmt.Println("Value of a before ZeroVal Call: ", a)
	zeroptr(&a)
	fmt.Println("Value of a after ZeroVal Call: ", a)
	fmt.Printf("Value of a: %d, Value of b: %d before swap\n", a, b)
	swap(&a, &b)
	fmt.Printf("Value of a: %d, Value of b: %d after swap\n", a, b)
}

// This is a way to create a interface with a method that accepts any type of struct
// Any struct which has a method with the same signature as the interface can be passed to the function
type PERSON interface {
	String() string
}

type Worker struct {
	age  int
	name string
}

type Student struct {
	age  int
	name string
}

// A simple way to create methods for a struct. Similar to impl from rust
func (w Worker) String() string {
	return fmt.Sprintf("%s is %d years old", w.name, w.age)
}

func (w *Worker) StringP() string {
	return fmt.Sprintf("%s is %d years old", w.name, w.age)
}

// A simple way to create methods for a struct. Similar to impl from rust
func (s Student) String() string {
	return fmt.Sprintf("%s is %d years old", s.name, s.age)
}

func (s *Student) StringP() string {
	return fmt.Sprintf("%s is %d years old", s.name, s.age)
}

// Go automatically dereferences pointers when calling methods
func structs() {
	p := Worker{age: 20, name: "John"}
	fmt.Println(p.StringP())
	pp := &p
	fmt.Println(pp.String())
	pr := func(p PERSON) {
		fmt.Println(p.String())
	}
	pr(p)
}
