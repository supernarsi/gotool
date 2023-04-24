package douglas

import "math"

type Point struct {
	X, Y float64
}

func DouglasPeucker(points []Point, epsilon float64) []Point {
	if len(points) < 3 {
		return points
	}

	maxDist := 0.0
	index := 0
	ln := line{points[0], points[len(points)-1]}

	for i := 1; i < len(points)-1; i++ {
		dist := ln.pointDistance(points[i])
		if dist > maxDist {
			maxDist = dist
			index = i
		}
	}

	if maxDist >= epsilon {
		left := DouglasPeucker(points[:index+1], epsilon)
		right := DouglasPeucker(points[index:], epsilon)

		return append(left[:len(left)-1], right...)
	}

	return []Point{points[0], points[len(points)-1]}
}

type line struct {
	A, B Point
}

func (l line) pointDistance(p Point) float64 {
	numerator := math.Abs((l.B.Y-l.A.Y)*p.X - (l.B.X-l.A.X)*p.Y + l.B.X*l.A.Y - l.B.Y*l.A.X)
	denominator := math.Sqrt(math.Pow(l.B.Y-l.A.Y, 2) + math.Pow(l.B.X-l.A.X, 2))
	return numerator / denominator
}
