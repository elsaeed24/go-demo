package main

import "fmt"

// متغير عام (Global Variable)
var age int = 15

// ثابت
const PI = 3.14

// Function
func sayHello() {
	fmt.Println("Hello")
}

func main() {
	fmt.Println(age)
	sayHello()

	student := Student{
		Name:       "Ahmed",
		Age:        24,
		Grade:      90.5,
		Email:      "ahmed@test.com",
		IsPassed:   true,
		University: "Mansoura University",
		Subjects: []string{
			"Math",
			"Physics",
			"Programming",
		},
	}

	student.Print()
}

//+ - * / %
//== != > < >= <=
//&& ||
