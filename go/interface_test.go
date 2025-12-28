package main

import (
	"fmt"
	"math"
	"testing"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printShapeArea(s Shape) {
	fmt.Printf("%.2f\n", s.Area())
}

func Test_BasicInterface(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 20}
	circle := Circle{Radius: 5}
	printShapeArea(rect)
	printShapeArea(circle)
}

// ❯ go test -v -run Test_BasicInterface
// === RUN   Test_BasicInterface
// 200.00
// 78.54
// --- PASS: Test_BasicInterface (0.00s)
// PASS
// ok      github.com/mridulgain/demo-codes/go     0.589s
