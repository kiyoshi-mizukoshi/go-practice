package main

import (
	"fmt"
	"go-basics/calculator"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))
	PracticeVariables()
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
