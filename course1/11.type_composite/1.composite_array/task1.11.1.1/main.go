package main

func sum(xs [8]int) int {
	total := 0
	for _, val := range xs {
		total += val
	}
	return total
}

func average(xs [8]int) float64 {
	return float64(sum(xs)) / 8
}

func averageFloat(xs [8]float64) float64 {
	var total float64 = 0
	for _, val := range xs {
		total += val
	}
	return total / 8
}

func reverse(xs [8]int) [8]int {
	begin := 0
	end := len(xs) - 1
	for begin < end {
		xs[begin], xs[end] = xs[end], xs[begin]
		begin++
		end--
	}
	return xs
}
