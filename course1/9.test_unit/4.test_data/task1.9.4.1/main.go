package main

func average(xs []float64) float64 {
	var result float64 = 0
	for _, val := range xs {
		result += val
	}

	return result / float64(len(xs))
}

func main() {

}
