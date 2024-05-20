package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// godotenv.Load()
	// fmt.Println(os.Getenv("GO_ENV"))
	// fmt.Println(calculator.Offset)
	// fmt.Println(calculator.Sum(1, 2))
	// fmt.Println(calculator.Multiply(1, 2))

	// PracticeVariables()

	PracticePointer()
}

func PracticeVariables() {

	type Os int

	const (
		Mac Os = iota + 1
		Windows
		Linux
	)

	var (
		int    int
		string string
		bool   bool
	)

	int = 1
	string = "Go"
	bool = true

	fmt.Println(int, string, bool)

	// var i int
	i := 1
	ui := uint16(2)

	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)
}

func PracticePointer() {
	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)
	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2)
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[byte]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of ui1(dereferenced): %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("size of pp1: %d[byte]\n", unsafe.Sizeof(pp1))
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("value of p1(dereferenced): %v\n", *pp1)
	fmt.Printf("value of ui1(dereferenced): %v\n", **pp1)
}
