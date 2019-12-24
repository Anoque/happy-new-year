package main

import (
  "os/exec"
  "fmt"
  "log"
  "os"
  "strings"
  "strconv"
)

func main() {
  cmd := exec.Command("stty", "size")
  cmd.Stdin = os.Stdin
  out, err := cmd.Output()
  // fmt.Printf("out: %#v", string(out))
  if err != nil {
    log.Fatal(err)
  }

  s := strings.Split(string(out), " ")

  le, err := strconv.Atoi(s[0])

  if err != nil {
    log.Fatal(err)
  }

  s = strings.Split(s[1], "\n")

  l2, err := strconv.Atoi(s[0])

  if err != nil {
    log.Fatal(err)
  }


  for i := 0; i < le - 5; i++ {
    b := l2 / 2 - i

    if i * 2 + 1 >= l2 {
      break
    }

    for j := 0; j < b; j++ {
      fmt.Printf(" ")
    }

    for j := b; j < b + i * 2 + 1; j++ {
      fmt.Printf("*")
    }

    fmt.Printf("\n")
  }

  for i := 0; i < 3; i++ {
    for j := 0; j < l2 / 2 - 1; j++ {
      fmt.Printf(" ")
    }

    fmt.Printf("||\n")
  }
}
