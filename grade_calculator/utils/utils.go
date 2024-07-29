package utils

func GetGrade(score float32) string {
	switch {
	case score >= 90:
		return "A+"
	case score >= 85:
		return "A"
	case score >= 80:
		return "B+"
	case score >= 75:
		return "B-"
	case score >= 70:
		return "B"
	case score >= 60:
		return "C"
	case score >= 50:
		return "D"
	default:
		return "F"
	}
}

func GetAverage(grades map[string]float32) float32 {
	var sum float32
	if len(grades) == 0 {
		return 0
	}
	for _, value := range grades {
		sum += value
	}
	average := sum / float32(len(grades))
	return average
}
