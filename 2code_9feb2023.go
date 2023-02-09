package main

import "fmt"

func main() {


	var students [3]string
	students[0] = "jubayer"
	students[1] = "rofik"
	students[2] = "minhaz"

	fmt.Println(students)
	fmt.Println(len(students))
	
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
	
	fmt.Println()
	fmt.Println()
	
// short hand way
	
	students2 := [...]string{"samad", "anik", "rafi", "himel", "opu"}
	fmt.Println(students2, len(students2))




}