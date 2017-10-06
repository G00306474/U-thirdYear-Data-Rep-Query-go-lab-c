//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced:  https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go

package main

import (
	"fmt"	//import input/output
	"sort"	// import  sort
)


func main(){
		
	var n int
	var r int
	
	fmt.Println("Please Enter size of array 1:")
  fmt.Scan(&n)
 
  array1 := make([]int, n)//creates array to user sze 
  for i := 0; i < n; i++ {
      fmt.Scan(&array1[i])//takes input for each postition
    }
  fmt.Println("Nums in array 1",array1)//prints contents of array
	
	fmt.Println("Please Enter size of array 2:")
  fmt.Scan(&r)
 
  array2 := make([]int, r)
  for i := 0; i < r; i++ {
      fmt.Scan(&array2[i])
    }
  fmt.Println("Nums in array 2",array2)
	
	
	
	
	
		mergedArray := append(array1, array2...)//merges array1 and array2
	
		sort.Ints(mergedArray)//sorts merged array by size
		fmt.Println("List Array Merged and Sorted: ",mergedArray)//prints merged array
}//main