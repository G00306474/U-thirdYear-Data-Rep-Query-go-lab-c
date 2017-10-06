//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced: https://gist.github.com/abesto/3476594

package main

import (
	"fmt"
	"math"//imports math 
)

const DELTA = 0.0000001//creates variable Delta and assigns it a value  
const INITIAL_Z = 100.0//creates variable  INITIAL_Z and assigns it a value 

func Sqrt(x float64) (z float64) {//start off function called  Sqrt
    z = INITIAL_Z//makes z  = t inital comment 
    
    step := func() float64 {
    	return z - (z*z - x) / (2 * z)//returns sqrt
    }
    
    for zz := step(); math.Abs(zz - z) > DELTA//for loop
    {
    	z = zz
	zz = step()
    }
    return
}//end off function called  Sqrt

func main() {//start off function called   main
	fmt.Println(Sqrt(500))//calls function sqrt and passes value 500
	fmt.Println(math.Sqrt(500))//calls function math & sqrt and passes value 500 
}//end off function called   main