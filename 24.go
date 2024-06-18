package main

import (
	"fmt"
	"math"
)

// Point представляет точку в двумерном пространстве.
type Point struct {
	x, y float64
}

// NewPoint создает новый экземпляр точки с заданными координатами.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// DistanceTo вычисляет расстояние между текущей точкой и переданной точкой.
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
