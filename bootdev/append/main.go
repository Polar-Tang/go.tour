package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	var floatSlice []float64
	for _, cost := range costs {
		floatSlice = append(floatSlice, float64(cost))

	}
	return floatSlice
}
