//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced: https://play.golang.org/p/5FUOzjBa-o

package main

import "fmt"

func isPalindrome(input string) bool {//boolean function returns true or false
	for i := 0; i < len(input)/2; i++ {//runs a loop for half the length of the word
		if input[i] != input[len(input)-i-1] {//compares each letter against the oppiste to see if they match if they dont returns false
			return false//returns false
		}
	}
	return true//returns false
}

func main() {
	
	var input string
    
    fmt.Println("Enter word to see if it is a palindrome")
    fmt.Scanf("%s", &input)//stores user input into variable called "input"
	 fmt.Println("True if it is/False if it isnt")
    fmt.Println(isPalindrome(input))//returns fuction isPalindrome result
}
