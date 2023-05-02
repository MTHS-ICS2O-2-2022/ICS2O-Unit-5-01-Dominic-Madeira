// Created by: Dominic M.
// Created on: April 2023
//
// This program converts fahrenheit to celsius.
package main


import (
  "fmt"
)


func main() {
  var celsius float64
  var fahrenheit float64


  // input
  fmt.Println("This program converts fahrenheit to celsius.")
  fmt.Println()
  fmt.Print("Enter a temperature in fahrenheit: ")
  fmt.Scanln(&fahrenheit)


  // process
  celsius = (fahrenheit - 32) * 5 / 9


  // output
  fmt.Println()
  celsiusFormatted := fmt.Sprintf("%.2f", celsius)
  fmt.Println("The temperature is:", celsiusFormatted, "°C")


  fmt.Println("\nDone.")
}
