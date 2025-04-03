package main

import (
  "flag"
  "fmt"
  "os"
)

func main() {
  // first parameter - name of the flag (what users type on the cli)
  // second parameter - default value that is used if the user doesnt specify a value
  // third parameter - usage string, appears when someone uses the help flag
  count := flag.Int("count", 1, "Number of times to print Hello World")

  flag.Parse()

  fmt.Println("count has value: ", *count)
  // cli input validation
  if *count < 1 {
    os.Exit(1)
  }

  for i := 0; i < *count; i++ {
    fmt.Println("hello!")
  }

}

