package main

import "fmt"
import "reflect"

func main() {


	var students [3]string

	students[0] = "jubayer"
	students[1] = "sagor"
	students[2] = "toufiq"

	x := students[0:3]
	fmt.Println(x)

	y := make([]string, 3)
	fmt.Println(y)
	
	var fruits []string
	fruits = append(fruits, "apple", "banana", "mango")
	fmt.Println(fruits, len(fruits))
	fmt.Println("%T \n", fruits)
	fmt.Println("%T", students)

	a := reflect.TypeOf(students).Kind().String()
	fmt.Println(a)
	
	b := reflect.TypeOf(fruits).Kind().String()
	fmt.Println(b)
	



}