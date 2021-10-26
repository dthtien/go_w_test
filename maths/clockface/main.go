package main

import (
  "os"
  "time"

  "go_fundamentals/maths"
)

func main() {
  t := time.Now()
  clockface.SVGWriter(os.Stdout, t)
}
