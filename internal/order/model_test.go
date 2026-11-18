package order

import (
	"testing"
)

func TestCanTransition(t *testing.T) {
	all := []Status{StatusCreated, StatusPaid, StatusCancelled, StatusShipped}
	allowed := map[[2]Status]bool{
		{StatusCreated, StatusPaid}:      true,
		{StatusCreated, StatusCancelled}: true,
		{StatusPaid, StatusShipped}:      true,
		{StatusPaid, StatusCancelled}:    true,
	}
	for _, from := range all {
		for _, to := range all {
			got := CanTransition(from, to)
			want := allowed[[2]Status{from, to}]
			if got != want {
				t.Errorf("CanTransition(%s, %s) = %v, want %v", from, to, got, want)
			}
		}
	}
}

func TestValidStatus(t *testing.T) {
	for _, s := range []Status{"created", "paid", "shipped", "cancelled"} {
		if !ValidStatus(s) {
			t.Errorf("ValidStatus(%q) = false, want true", s)
		}
	}
	for _, s := range []Status{"", "canceled", "CREATED", "unknown"} {
		if ValidStatus(s) {
			t.Errorf("ValidStatus(%q) = true, want false", s)
		}
	}
}
