package main

import "fmt"

type Student struct {
	Name       string
	Age        int
	Grade      float64
	Email      string
	IsPassed   bool
	Subjects   []string
	University string
}

func (s Student) Print() {
	fmt.Println("===== Student Data =====")
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Grade:", s.Grade)
	fmt.Println("Email:", s.Email)
	fmt.Println("Passed:", s.IsPassed)
	fmt.Println("University:", s.University)

	fmt.Println("Subjects:")
	for _, subject := range s.Subjects {
		fmt.Println("*", subject)
	}
}
