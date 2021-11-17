package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	tm, _ := time.Parse("1/2/2006 15:04:05", date)
	return tm
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	tm, _ := time.Parse("January 1, 2006 15:04:05", date)
	dur := time.Since(tm)
	return dur > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	tm, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return tm.Hour() >= 12 && tm.Hour() <= 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	tm := Schedule(date)
	str := tm.Format("Monday, January 2, 2006, at 15:04.")
	return fmt.Sprintf("You have an appointment on %s", str)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(2021, 9, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	tm := Schedule("7/13/2020 20:32:00")
	fmt.Printf("%s\n", tm)

	fmt.Println(HasPassed("October 3, 2022 20:32:00"))

	fmt.Println(Description("7/13/2020 20:32:00"))
}
