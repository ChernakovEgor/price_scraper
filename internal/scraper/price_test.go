package scraper

import (
	"testing"
)

func TestGetPrice(t *testing.T) {
	tests := []struct {
		input string
		want  float64
		err   error
	}{
		{"12345", 12345, nil},
		{"  12 345 ", 12345, nil},
		{"12345$", 12345, nil},
		{"12 345 $", 12345, nil},
		{"12 345 USD", 12345, nil},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, err := getPrice(test.input)

			if test.err == nil && err != nil {
				t.Errorf("got error on valid: %v", err)
			}
			if test.err != nil && err == nil {
				t.Errorf("no error on incorrect input: %v", err)
			}

			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
