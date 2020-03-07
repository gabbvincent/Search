//Name: Vincent Gabb
//Date: 3/6/2020
//Description: Search


package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  var input int

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

    array := [5]int{r1.Intn(25), r1.Intn(25), r1.Intn(25), r1.Intn(25), r1.Intn(25)}
    
    fmt.Println(array)

    fmt.Println("Please enter a number. (25 and under)")
    
    fmt.Scanln(&input)
  
     if input == array[0] || input == array[1] || input == array[2] || input == array[3] || input == array[4] {
    
     fmt.Println("Nice! You got it.")
    
      } else if input != array[0] && input != array[1] && input != array[2] &&input != array[3] && input != array[4] {
      
      for input != array[0] && input != array[1] && input != array[2] &&input != array[3] && input != array[4] {
     
      fmt.Println("Nope, Try again.")
     
      fmt.Scanln(&input)
     
      if input == array[0] || input == array[1] || input == array[2] || input == array[3] || input == array[4] {
     
      fmt.Println("Nice! You got it.")
    
    }           
  }
 }
}