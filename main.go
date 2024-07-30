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
	fmt.Println()
	fmt.Println("Enum")
	EnumTester()
	fmt.Println()
	fmt.Println("Struct Embedding")
	struct_embedding()
	fmt.Println()
	fmt.Println("Generics")
	GenericTester()
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
// Calling a function like this is also known as the
// fmt.Stringer interface similar to the ToString method by default in Java
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

/*
	Since the String() method is in both the Worker struct and Student struct, the interface can just
	mimic calling both these functions automatically without having to call for each struct explicitly
*/
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

// Create a type called ServerState
type ServerState int

// This is the ENUM Defined for all the possible server states
// Since there is no definite keywork for enums, we just use the const
// and assign a number to the ENUM using the IOTA
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// This is to Map each state to a certain string to make the readability easier while debugging
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "error",
}

// This transition function simply switches the state based on what state I am in currently
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	}
	return StateConnected
}

func EnumTester() {
	ns := transition(StateIdle)
	fmt.Println(stateName[ns])
	ns2 := transition(ns)
	fmt.Println(stateName[ns2])
}

// This is all a part of Struct Embedding
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func struct_embedding() {
	co := container{
		base: base{
			num: 1,
		},
		str: "Some Name",
	}
	fmt.Printf("co={num: %v , str: %v}\n", co.num, co.str)
	fmt.Println("also num: ", co.num)
	fmt.Println("describe: ", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer: ", d.describe())

	var do describer = base{num: 10}
	fmt.Println("describer: ", do.describe())
}

// Generics
// K has the constraint of being comparable so == and != are valid comparisons for the key and V can be any type
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

type node[T any] struct {
	next *node[T]
	val  T
}

type List[T any] struct {
	head, tail *node[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &node[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &node[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func GenericTester() {
	var m = map[int]string{1: "2", 3: "4", 5: "6"}
	fmt.Println("Keys: ", MapKeys(m))
	fmt.Println("Values: ", MapValues(m))
	lst := List[int]{}
	lst.Push(10)
	lst.Push(11)
	lst.Push(12)
	fmt.Println(lst.GetAll())
}
