package utils

import (
	"testing"
)

func TestGetGrade(t *testing.T) {
	tests := []struct {
		name     string
		score    float32
		expected string
	}{
		{"A+ grade", 95, "A+"},
		{"A grade", 88, "A"},
		{"B+ grade", 82, "B+"},
		{"B- grade", 78, "B-"},
		{"B grade", 70, "B"},
		{"C grade", 65, "C"},
		{"D grade", 55, "D"},
		{"F grade", 45, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetGrade(tt.score)
			if result != tt.expected {
				t.Errorf("GetGrade(%v) = %v; want %v", tt.score, result, tt.expected)
			}
		})
	}
}

func TestGetAverage(t *testing.T) {
	tests := []struct {
		name     string
		grades   map[string]float32
		expected float32
	}{
		{"Average of three grades", map[string]float32{"Math": 80, "English": 90, "History": 70}, 80},
		{"Average of four grades", map[string]float32{"Math": 60, "English": 70, "History": 80, "Science": 90}, 75},
		{"Average of single grade", map[string]float32{"Math": 100}, 100},
		{"Average of no grades", map[string]float32{}, 0}, // Edge case for no grades
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAverage(tt.grades)
			if result != tt.expected {
				t.Errorf("GetAverage(%v) = %v; want %v", tt.grades, result, tt.expected)
			}
		})
	}
}
