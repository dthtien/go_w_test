package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
  secondHandLength = 90
  minuteHandLength = 80
  hourHandLength   = 50
  clockCentreX     = 150
  clockCentreY     = 150
)

type Point struct {
  X float64
  Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
  p := SecondHandPoint(t)
  p = Point{p.X * secondHandLength, p.Y * secondHandLength}
  p = Point{p.X, -p.Y}
  p = Point{p.X + clockCentreX, p.Y + clockCentreY} //translate
  return p
}

func SecondsInRadians(t time.Time) float64 {
  return (math.Pi / (30 / (float64(t.Second()))))
}

func SecondHandPoint(t time.Time) Point {
  return angleToPoint(SecondsInRadians(t))
}

//SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
  io.WriteString(w, svgStart)
  io.WriteString(w, bezel)
  secondHand(w, t)
  minuteHand(w, t)
  hourHand(w, t)
  io.WriteString(w, svgEnd)
}

func MinutesInRadians(t time.Time) float64 {
  return (SecondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func HoursInRadians(t time.Time) float64 {
  return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func angleToPoint(angle float64) Point {
  x := math.Sin(angle)
  y := math.Cos(angle)

  return Point{x, y}
}

func HourHandPoint(t time.Time) Point {
  return angleToPoint(HoursInRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
  return angleToPoint(MinutesInRadians(t))
}

func secondHand(w io.Writer, t time.Time) {
  p := makeHand(SecondHandPoint(t), secondHandLength)
  fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
  p := makeHand(MinuteHandPoint(t), minuteHandLength)
  fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p Point, length float64) Point {
  p = Point{p.X * length, p.Y * length}
  p = Point{p.X, -p.Y}
  return Point{p.X + clockCentreX, p.Y + clockCentreY}
}

func hourHand(w io.Writer, t time.Time) {
  p := makeHand(HourHandPoint(t), hourHandLength)
  fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
width="100%"
height="100%"
viewBox="0 0 300 300"
version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`
