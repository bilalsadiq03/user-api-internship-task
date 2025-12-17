package service

import (
	"testing"
	"time"

)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(2000, 5, 10, 0, 0, 0, 0, time.UTC)

	age := calculateAge(dob)

	expected := time.Now().Year() - 2000
	if time.Now().Month() < time.May || (time.Now().Month() == time.May && time.Now().Day() < 10) {	
		expected--
	}

	if age != expected {
		t.Errorf("expected age %d, got %d", expected, age)
	}
}