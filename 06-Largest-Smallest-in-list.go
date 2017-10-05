//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced: https://stackoverflow.com/questions/25510958/go-lang-print-inputted-array 

package main

import "fmt"

func main(){

	var min = 0
	var max = 0
	var n int
	

 fmt.Println("Please Enter size of array:")
  fmt.Scan(&n)
 
  array := make([]int, n)
  for i := 0; i < n; i++ {
      fmt.Scan(&array[i])
    }
  fmt.Println("Nums in array",array)
  
	min = array[0]
	max = array[0]

	for i := range array {
		if array[i] < min {
			min = array[i]
		}
		if array[i] > max {
			max = array[i]
		}
	}


fmt.Println("Lowest Number ",min)
fmt.Println("Highest Number", max)
}