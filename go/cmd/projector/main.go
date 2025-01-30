package main

import (
  "fmt"
  "log"
  "github.com/cavilarts/aoc/pkg/config"
)

func main()  {
  opts, err := config.GetOpts()

  if err != nil {
    log.Fatalf("unable to get options %v", err)
  }

  fmt.Printf("opts: %+v", opts)
}
