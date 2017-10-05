//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced:http://golangcookbook.blogspot.ie/2012/11/generate-random-number-in-given-range.html

package main
import(
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    
	var num1 = 1
	var num2 = 0
	
	//fmt.Println("Please Enter 1st number:")//input number 
	//fmt.Scanf("%d", &num1)//scans number
	
	fmt.Println("Please Enter High number:")//input number 
	fmt.Scanf("%d", &num2)//scans number
	
	myrand := random(num1, num2)
    fmt.Println(myrand)
	
}