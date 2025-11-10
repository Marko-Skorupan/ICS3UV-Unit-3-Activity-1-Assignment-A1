/*
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-10
 * @fileoverview "Program to calculate area of a circle"
 */

package main

import "fmt"

func main(){
//Set the constants
const PI float32 = 3.14159
const radius float32 = 5.6

//INPUT - none

//PROCESS
//calculate the area
area := PI * radius * radius 

//OUTPUT
//display the result
fmt.Print("The area of a circle with a radius of ", + PI ,"cm is " , + area," cm squared")

fmt.Println("\nDone.")
}
