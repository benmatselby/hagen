package ui_test

import (
	"testing"
	"time"

	"github.com/benmatselby/hagen/ui"
)

func TestHumanDuration(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		d    time.Duration
		want string
	}{
		{name: "days", d: 49 * time.Hour, want: "2d 1h 0m"},
		{name: "hours", d: 3 * time.Hour, want: "3h 0m"},
		{name: "minutes", d: 3 * time.Minute, want: "3m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ui.HumanDuration(tt.d)

			if got != tt.want {
				t.Errorf("HumanDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
