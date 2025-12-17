package service

import "time"


func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age
}