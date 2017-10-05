//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced:http://golangcookbook.blogspot.ie/2012/11/generate-random-number-in-given-range.html
//Gave user option to enter high number cause if i give them the option to enter high and low program breaks
//Also takes first guess without letting user enter input 
package main
import(
    "fmt"
    "math/rand"
    "time"
)

//this generates random number between given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    
    
    guessTaken := 0
    var guess int
	var prevGuess int
	var num1 = 1
	var num2 = 0
	
	//fmt.Println("Please Enter 1st number:")//input number 
	//fmt.Scanf("%d", &num1)//scans number
	
	fmt.Println("Please Enter High number:")//input number 
	fmt.Scanf("%d", &num2)//scans number
	
	myrand := xrand(num1, num2)
 
    fmt.Printf("I am thinking of a number between 1 and %d. You have 6 attempts \n",num2)
    
    //this is the while loop
    for guessTaken < 6 {
		fmt.Println("*********************")
        fmt.Println("Take a guess...")
        fmt.Scanf("%d\n",&guess)
        guessTaken++
		if guess == prevGuess {//does not work for some reason
            guessTaken--
			fmt.Println("Your guess is the same as last time.")
        }
		if guess < myrand {
            fmt.Println("Your guess is too low.")
        }else if guess > myrand {
            fmt.Println("Your guess is too high.")
        }else if guess == myrand {
            break
        }
		
		guess = prevGuess
		fmt.Printf("Count %d \n", guessTaken)
    }
    if guess == myrand {
        fmt.Printf("Good job! You guessed my number in %d tries\n", guessTaken)
    } else {
        fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
    }
}