package darts

func Score(x, y float64) int {
	distSq := x*x + y*y

	if distSq <= 1 {
		return 10
	}
	if distSq <= 25 {
		return 5
	}
	if distSq <= 100 {
		return 1
	}
	return 0
}
