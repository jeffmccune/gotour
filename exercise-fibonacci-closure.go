package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  fib := [2]int{0,1}
  f := func() int {
    r := fib[0]
    fib[0] = fib[1]
    fib[1] += r
    return r
  }
  return f
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
