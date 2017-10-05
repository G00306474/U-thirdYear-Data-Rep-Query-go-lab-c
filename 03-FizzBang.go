//Got From: Own Code
//Fizz Bang

package main

import "fmt"

func main() {//Start Main
	for i := 1;i<=100; i++{//Start for Loop, { to be on this line or wont run
		if i%3==0 && i%5==0 {//if i divided by three equal 0 and i divided by five equal zero go in here
		fmt.Println("Fizz Bang")//print fizz bang to screen 
		} else if i%3 ==0 {//if i divided by three equal 0 go in here
		fmt.Println("Fizz")//print fizz to screen 
		} else if i%5 ==0 {//if i divided by five equal 0 go in here
		fmt.Println("Bang")//print bang to screen 
		}else {//if no condition is met go in here
		fmt.Println(i)}//print number to screen
	}//End for Loop
}//End Main
