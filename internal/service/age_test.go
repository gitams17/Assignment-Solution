package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {

	tests := []struct {
		name     string
		dob      string
		expected int
	}{
		{
			name:     "Birthday has passed this year",
			dob:      time.Now().AddDate(-25, -1, 0).Format("2006-01-02"), 
			expected: 25,
		},
		{
			name:     "Birthday has NOT passed this year",
			dob:      time.Now().AddDate(-25, 1, 0).Format("2006-01-02"), 
			expected: 24,
		},
		{
			name:     "Born today",
			dob:      time.Now().Format("2006-01-02"),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dobTime, _ := time.Parse("2006-01-02", tt.dob)
			got := calculateAge(dobTime)
			if got != tt.expected {
				t.Errorf("calculateAge() = %v, want %v", got, tt.expected)
			}
		})
	}
}