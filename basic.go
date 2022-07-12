package main

import (
	"fmt"
	"time"
)

// func add(x int, y int) int {
// 	var result = x + y
// 	return result
// }
// func calc(x int, y int) (int, int) {
// 	var result1 = x + y
// 	var result2 = x - y
// 	return result1, result2
// }
func main() {
	// num1:=10
	// num2:=5
	// //var result=add(num1,num2)
	// //fmt.Println(result)
	//  x,y:=calc(num1,num2)
	// fmt.Println(x,y)
	// for  i:=0;i<5;i++ {
	// 	fmt.Println("Pallab")
	// }
	// var num float64
	// fmt.Scanf("%f",&num)
	// var result=math.Sqrt(num)
	// var r=math.Round(result)
	// var c=math.Ceil(result)
	// var f=math.Floor(result)

	// fmt.Printf("Output is %.2f\n",result)

	// fmt.Printf("Round is %.2f\n",r)
	// fmt.Printf("ceil is %.2f\n",c)
	// fmt.Printf("floor is %.2f\n",f)

	// var num int
	// fmt.Scan(&num)
	// if num == 1 {
	// 	fmt.Println("One")
	// } else if num==2 {
	// 	fmt.Println("Two")
	// } else if num==3{
	// 	fmt.Println("Three")
	// } else {
	// 	fmt.Println("None")
	// }

	today:=time.Now().Weekday()
	fmt.Println("Today is->",today)

	fmt.Println("Time ->",time.Now())
}
