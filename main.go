package main

import "fmt"

func main() {
  fmt.Println("Enter text: ")
  num := 0
  fmt.Scanln(&num)
  steps := 0
  for num != 1 {
    steps++
    if num % 2 == 0 {
      num /= 2
    } else {
      num = num * 3 + 1
    }
  }
  fmt.Printf("Quantity of steps: %d\n", steps)
}
