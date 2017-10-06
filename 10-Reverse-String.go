//Kevin Moran 
//G00306474@gmit.ie
//Code Sourced: https://coderwall.com/p/fw6miw/reverse-text-in-golang

package main

import "fmt"

func main() {
    var input string
    
	fmt.Println("Enter word to see it reversed")
    fmt.Scanf("%s", &input)//stores user input into variable called "input"
	
    fmt.Println(Reverse(input))//passses variable input to function xcalled Reverse
 

}
func Reverse(s string) string {
    var reverse string
    for i := len(s)-1; i >= 0; i-- {//gets length of string and then counts down from that num
        reverse += string(s[i])//stores letter of string s in string reverse. as it is counting down it stores it in backwards 
    }
    return reverse //returns string reverse 
}
