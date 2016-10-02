// embed: Exploring embedding within structs.

package main

import "fmt"

type Point struct {
    X, Y int
}

type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}

func main() {
    w := Wheel{Circle{Point{8, 8}, 5}, 20}

    w2 := Wheel{
        Circle: Circle{
            Point: Point{X: 8, Y:8},
            Radius: 5, // note trailing comma is necessary
        },
        Spokes: 20, // note trailing comma is necessary
    }

    fmt.Printf("%#v\n", w)
    fmt.Printf("%#v\n", w2)

    w.X = 42
    fmt.Printf("%#v\n", w)

    w.Circle.Point.X = 8
    fmt.Printf("%#v\n", w)
}
