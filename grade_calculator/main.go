package main

import (
	"fmt"
	"grade_calculator/utils"
)

func main() {
	var n int
	scores := make(map[string]float32)

	fmt.Print("How many courses did you take? ")
	fmt.Scan(&n)

	fmt.Println("Enter course names and scores separated by a space (e.g., English 90): for ", n, " courses")

	for i := 0; i < n; i++ {
		var course string
		var score float32
		fmt.Scan(&course, &score)
		if score < 0 || score > 100 {
			fmt.Println("Invalid score. Please enter a score between 0 and 100.")
			i--
			continue
		}
		scores[course] = score
	}

	average_score := utils.GetAverage(scores)
	average_grade := utils.GetGrade(average_score)

	fmt.Println("Grades per course:")
	for course, score := range scores {
		fmt.Printf("%s: %s\n", course, utils.GetGrade(score))
	}

	fmt.Printf("Average Score: %.2f\n", average_score)
	fmt.Printf("Average Grade: %s\n", average_grade)
}
