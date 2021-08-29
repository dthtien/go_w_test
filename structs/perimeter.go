package main

import "math"

type Shape interface {
  Area() float64
}

type Circle struct {
  Radius float64
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.Radius
}

type Rectangle struct {
  Width float64
  Height float64
}

func (regtangle Rectangle) Area() float64 {
  return regtangle.Height * regtangle.Width
}

func (regtangle Rectangle) Perimeter() float64 {
  return (regtangle.Height + regtangle.Width) * 2
}

type Triangle struct {
  Base   float64
  Height float64
}

func (t Triangle) Area() float64 {
  return (t.Base * t.Height) * 0.5
}
