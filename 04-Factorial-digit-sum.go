//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced:
package main

import "fmt"
//import "math/big" imports big int 

func main() {

	var num = 0// creates num aas int 
	var total  =  0// creates total aas int 
	var factorial =  0// creates factorial aas int 
	
	//factorial := big.NewInt(0) creates big int 
	

	fmt.Println("Please Enter a number:")//input number 
	fmt.Scanf("%d", &num)//scans number


	if num <= 0{
		fmt.Println("Number cannot be less than 0!")
		for num <= 0{
			fmt.Println("Please Enter a number greater than zero:")//input number 
			fmt.Scanf("%d", &num)//scans number
			fmt.Scan("%d", &num)
		}
	}
	

	for i := 1; i <= num; i++{
		factorial = factorial * i//makes the factorial equal too the factorial * i 
		total = total + i
	
	}
	for i := 1; i <= num; i++{
			factorial = factorial * i//makes the factorial equal too the factorial * i 
			total = total + i
	}
	
	
	fmt.Println("Number Factorial:", factorial)//prints factorial to screen
	fmt.Println("Sum of Factorial:", total)//prints total to screen

}