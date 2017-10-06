//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced:  

package main

import (
	"fmt"	//import input/output
	"sort"	// import  sort
)


func main(){
		//array1 := []int{10, 40, 60}
		//array2 := []int{20, 30, 50}
	var n int
	var r int
	
	fmt.Println("Please Enter size of array 1:")
  fmt.Scan(&n)
 
  array1 := make([]int, n)
  for i := 0; i < n; i++ {
      fmt.Scan(&array1[i])
    }
  fmt.Println("Nums in array 1",array1)
	
	fmt.Println("Please Enter size of array 2:")
  fmt.Scan(&r)
 
  array2 := make([]int, r)
  for i := 0; i < r; i++ {
      fmt.Scan(&array2[i])
    }
  fmt.Println("Nums in array 2",array2)
	
	
	
	
	
		mergedArray := append(array1, array2...)
	
		sort.Ints(mergedArray)//sorts array by size
		fmt.Println("List AB Merged and Sorted: ",mergedArray)
}//main