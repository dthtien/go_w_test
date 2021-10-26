package clockface_test

import (
  "math"
  "testing"
  "time"
	"bytes"
	"encoding/xml"

	"go_fundamentals/maths"
)

func TestSecondsInRadians(t *testing.T) {
  cases := []struct {
    time  time.Time
    angle float64
  }{
    {simpleTime(0, 0, 30), math.Pi},
    {simpleTime(0, 0, 0), 0},
    {simpleTime(0, 0, 45), (math.Pi / 2) * 3},
    {simpleTime(0, 0, 7), (math.Pi / 30) * 7},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := clockface.SecondsInRadians(c.time)
      if got != c.angle {
        t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
      }
    })
  }
}

func TestSecondHandPoint(t *testing.T) {
  cases := []struct {
    time  time.Time
    point clockface.Point
  }{
    {simpleTime(0, 0, 30), clockface.Point{0, -1}},
    {simpleTime(0, 0, 45), clockface.Point{-1, 0}},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := clockface.SecondHandPoint(c.time)
      if !roughlyEqualPoint(got, c.point) {
        t.Fatalf("Wanted %v Point, but got %v", c.point, got)
      }
    })
  }
}

func TestSecondHandAt30Seconds(t *testing.T) {
  tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

  want := clockface.Point{X: 150, Y: 150 + 90}
  got := clockface.SecondHand(tm)

  if got != want {
    t.Errorf("Got %v, wanted %v", got, want)
  }
}

func roughlyEqualFloat64(a, b float64) bool {
  const equalityThreshold = 1e-7
  return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b clockface.Point) bool {
  return roughlyEqualFloat64(a.X, b.X) &&
  roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
  return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
  return t.Format("15:04:05")
}

type SVG struct {
  XMLName xml.Name `xml:"svg"`
  Xmlns   string   `xml:"xmlns,attr"`
  Width   string   `xml:"width,attr"`
  Height  string   `xml:"height,attr"`
  ViewBox string   `xml:"viewBox,attr"`
  Version string   `xml:"version,attr"`
  Circle  Circle   `xml:"circle"`
  Line    []Line   `xml:"line"`
}

type Circle struct {
  Cx float64 `xml:"cx,attr"`
  Cy float64 `xml:"cy,attr"`
  R  float64 `xml:"r,attr"`
}

type Line struct {
  X1 float64 `xml:"x1,attr"`
  Y1 float64 `xml:"y1,attr"`
  X2 float64 `xml:"x2,attr"`
  Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
  tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
  b := bytes.Buffer{}

  clockface.SVGWriter(&b, tm)

  svg := SVG{}

  xml.Unmarshal(b.Bytes(), &svg)

  want := Line{150, 150, 150, 60}

  for _, line := range svg.Line {
    if line == want {
      return
    }
  }

  t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", want, svg.Line)
}

func TestSVGWriterSecondHand(t *testing.T) {
  cases := []struct {
    time time.Time
    line Line
  }{
    {
      simpleTime(0, 0, 0),
      Line{150, 150, 150, 60},
    },
    {
      simpleTime(0, 0, 30),
      Line{150, 150, 150, 240},
    },
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      b := bytes.Buffer{}
      clockface.SVGWriter(&b, c.time)

      svg := SVG{}
      xml.Unmarshal(b.Bytes(), &svg)

      if !containsLine(c.line, svg.Line) {
        t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
      }
    })
  }
}


func TestMinutesInRadians(t *testing.T) {
  cases := []struct {
    time  time.Time
    angle float64
  }{
    {simpleTime(0, 30, 0), math.Pi},
    {simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := clockface.MinutesInRadians(c.time)
      if got != c.angle {
        t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
      }
    })
  }
}

func TestMinuteHandPoint(t *testing.T) {
  cases := []struct {
    time  time.Time
    point clockface.Point
  }{
    {simpleTime(0, 30, 0), clockface.Point{0, -1}},
    {simpleTime(0, 45, 0), clockface.Point{-1, 0}},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := clockface.MinuteHandPoint(c.time)
      if !roughlyEqualPoint(got, c.point) {
        t.Fatalf("Wanted %v Point, but got %v", c.point, got)
      }
    })
  }
}

func TestSVGWriterHourHand(t *testing.T) {
  cases := []struct {
    time time.Time
    line Line
  }{
    {
      simpleTime(6, 0, 0),
      Line{150, 150, 150, 200},
    },
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      b := bytes.Buffer{}
      clockface.SVGWriter(&b, c.time)

      svg := SVG{}
      xml.Unmarshal(b.Bytes(), &svg)

      if !containsLine(c.line, svg.Line) {
        t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
      }
    })
  }
}

func containsLine(l Line, ls []Line) bool {
  for _, line := range ls {
    if line == l {
      return true
    }
  }
  return false
}

func TestHoursInRadians(t *testing.T) {
  cases := []struct {
    time  time.Time
    angle float64
  }{
    {simpleTime(6, 0, 0), math.Pi},
    {simpleTime(0, 0, 0), 0},
    {simpleTime(21, 0, 0), math.Pi * 1.5},
    {simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := clockface.HoursInRadians(c.time)
      if !roughlyEqualFloat64(got, c.angle) {
        t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
      }
    })
  }
}

func TestHourHandPoint(t *testing.T) {
  cases := []struct {
    time  time.Time
    point clockface.Point
  }{
    {simpleTime(6, 0, 0), clockface.Point{0, -1}},
    {simpleTime(21, 0, 0), clockface.Point{-1, 0}},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := clockface.HourHandPoint(c.time)
      if !roughlyEqualPoint(got, c.point) {
        t.Fatalf("Wanted %v Point, but got %v", c.point, got)
      }
    })
  }
}
