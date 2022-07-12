package main

import "fmt"

type Student struct {
	roll  int
	name  string
	marks int
	
}

func main() {
	var s = Student{50, "pallab", 89}
	
	fmt.Println(s)
}
