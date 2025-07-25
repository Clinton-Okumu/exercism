package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.

func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("error parsing date:", err)
	}
	return parsedTime
}

// HasPassed returns whether a date has passed.

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("error parsing date:", err)
		return false
	}
	return parsedTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("error parsing date", err)
	}
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.

func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "invalid date"
	}
	dayPart := t.Format("Monday, January 2, 2006")
	timePart := t.Format("15:04")
	return fmt.Sprintf("You have an appointment on %s, at %s.", dayPart, timePart)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
