package main

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{name: "freezing point", celsius: 0, expected: 32},
		{name: "boiling point", celsius: 100, expected: 212},
		{name: "negative value", celsius: -40, expected: -40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := celsiusToFahrenheit(tt.celsius); got != tt.expected {
				t.Fatalf("celsiusToFahrenheit(%v) = %v, want %v", tt.celsius, got, tt.expected)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{name: "absolute zero", celsius: -273.15, expected: 0},
		{name: "zero celsius", celsius: 0, expected: 273.15},
		{name: "body temperature", celsius: 37, expected: 310.15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := celsiusToKelvin(tt.celsius); got != tt.expected {
				t.Fatalf("celsiusToKelvin(%v) = %v, want %v", tt.celsius, got, tt.expected)
			}
		})
	}
}
