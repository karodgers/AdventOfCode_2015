package solutions

func WrappingPaper(l, w, h int) int {
	values := []int{l, w, h}
	surfaceArea := 2*l*w + 2*w*h + 2*h*l

	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	newValues := values[:len(values)-1]

	var result int
	slack := 1
	for _, num := range newValues {
		slack = slack * num
	}
	result = slack + surfaceArea

	return result
}
