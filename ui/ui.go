// Package ui contains user interface related configuration
package ui

import (
	"fmt"
	"time"
)

// MoreResults provides a consistent message to get more results
const MoreResults = "\nPress enter for more results\n"

// HumanFriendlyDateFormat is a date format string for displaying dates in a
// human friendly way
const HumanFriendlyDateFormat = "Monday 02 January, 2006 at 15:04 MST"

// DateFormat is the standard date format string
const DateFormat = "2006-01-02 15:04:05 -07:00"

// HumanDuration converts a time.Duration to a human-readable string
func HumanDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}
