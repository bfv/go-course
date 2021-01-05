package main

 import "fmt"
 import "strings"

 func main() {
   var inputstring string	 
   fmt.Println("Enter a string")
   fmt.Scan(&inputstring)
   inputstring = strings.ToLower(inputstring)
   if (strings.ContainsAny(inputstring, "a") && strings.HasPrefix(inputstring,"i") && strings.HasSuffix(inputstring,"n") )    {
     fmt.Println("Found!")	   
   } else {	     
     fmt.Println("Not Found!")	   
   }
 }

