package main

import (
  "strings"
  "fmt"
)

func getInput() string {
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`;
}

type Thing = int

const (
  Tree Thing = iota
  Snow
)

func main() {
  treeCount := 0
  for r, line := range strings.Split(getInput(), "\n") {
    if string(line[r * 3 % len(line)]) == "#" {
      treeCount += 1
    }
  }

  fmt.Printf("treecount %v", treeCount);
}
