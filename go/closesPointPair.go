package main

import (
	"fmt"
	"math"
)

func closesPointPair(p [][]int) float64 {
	px := make([][]int, len(p))
	py := make([][]int, len(p))

	for i := 0; i < len(p); i++ {
		px[i] = p[i]
		py[i] = p[i]
	}

	qsort(px, len(p), compareX)
	qsort(py, len(p), compareY)

	return closestUntil(px, py, len(p)-1)
}

func closestUntil(px, py [][]int, n int) float64 {
	if n <= 3 {
		return bruteForce(px, n)
	}

	mid := n / 2
	midPoint := px[mid]

	Pyl := make([][]int, mid)
	Pyr := make([][]int, n-mid)
	var li, ri int

	for i := 0; i < n; i++ {
		if (py[i][0] < midPoint[0] || (py[i][0] == midPoint[0] && py[i][1] < midPoint[1])) && li < mid {
			Pyl[li] = py[i]
			li++
		} else {
			Pyr[ri] = py[i]
			ri++
		}
	}

	dl := closestUntil(px, Pyl, mid)
	dr := closestUntil(px, Pyr, n-mid)

	d := math.Min(dl, dr)

	strip := make([][]int, n-mid)
	j := 0

	for i := 0; i < n; i++ {
		if math.Abs(float64(py[i][0]-midPoint[0])) < d {
			strip[j] = py[i]
			j++
		}
	}

	return stripClosest(strip, j, d)
}

func stripClosest(strip [][]int, n int, d float64) float64 {
	min := d
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && math.Abs(float64(strip[i][1]-strip[j][1])) < min; j++ {
			min = distance(strip[i], strip[j])
		}
	}
	return min
}

func bruteForce(p [][]int, n int) float64 {
	min := math.MaxFloat64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			min = math.Min(min, distance(p[i], p[j]))
		}
	}
	return min
}
func distance(p1, p2 []int) float64 {
	return math.Sqrt(math.Pow(float64(p1[0]-p2[0]), 2) + math.Pow(float64(p1[1]-p2[1]), 2))
}

func main() {
	p := [][]int{{0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}, {-7, 1}, {-5, -3}, {0, 11}}
	fmt.Println(closesPointPair(p))

	fmt.Println(distance(p[0], p[1]))

}
