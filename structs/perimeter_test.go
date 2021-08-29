package main

import "testing"

func TestPerimeter(t *testing.T) {
  checkResult := func(t testing.TB, got, want float64) {
    if got != want {
      t.Errorf("got %.2f want %.2f", got, want)
    }
  }

  t.Run("Regtangle", func(t *testing.T) {
    regtangle := Rectangle{10.0, 10.0}
    got := regtangle.Perimeter()
    want := 40.0
    checkResult(t, got, want)
  })
  // t.Run("Circle", func(t *testing.T) {
    // circle := Circle{10.0}
    // got := circle.Perimeter()
    // want := 62.8318530718
    // checkResult(t, got, want)
  // })
}

func TestArea(t *testing.T) {
  areaTests := []struct {
    name    string
    shape   Shape
    hasArea float64
  }{
    {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
    {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
    {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
  }

  for _, tt := range areaTests {
    // using tt.name from the case to use it as the `t.Run` test name
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Area()
      if got != tt.hasArea {
        t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
      }
    })
  }
}
