// Created by: Dominic M.
// Created on: April 2023
//
// This program is a game where you guess a number from 1-6 and try to match the computer's number.
package main


import (
  "fmt"
  "math/rand"
  "time"
)


func main() {
  var input int


  // input
  fmt.Println("This program is a game where you guess a number from 1-6 and try to match the computer's number.")
  fmt.Println()
  fmt.Print("Please enter a number: ")
  fmt.Scanln(&input)


  // process
  s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
  random := r1.Intn(6) + 1

  // This program is an infinite loop until you guess the correct number.
  for input != random {
    fmt.Println("You have guessed incorrectly.")
    fmt.Println()
    fmt.Print("Please enter a number: ")
    fmt.Scanln(&input)
  }

  fmt.Println("You have guessed correctly! The number was", random)

  // // This program is a loop that will only run once.
  // if input == random {
  //   fmt.Println("You guessed correctly!")
  // } 
  
  // if input != random {
  //   fmt.Println("You guessed incorrectly.")
  // }

  fmt.Println("\nDone.")
}
