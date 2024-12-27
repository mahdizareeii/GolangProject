package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("Hello Mahdi")

	fmt.Println(swap("a", "b"))
}

func learnDefineVariables() {
	var name string
	var age int
	const PI = 3.14

	name = "Mahdi"
	age = 25

	fmt.Println(name, age)
}

func learnLoops() {
	for i := 0; i < 10; i++ {
		fmt.Println("iteration : ", i)
	}
}

func learnConditions(age int) {
	if age >= 18 {
		fmt.Println("Adult")
	}

	switch age {
	case 18:
		fmt.Println("you are", 18)
	default:
		fmt.Println("you are", age)
	}
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func operateSum(argument []int) int {
	var sum int
	for i := 0; i < len(argument); i++ {
		sum += argument[i]
	}
	return sum
}
