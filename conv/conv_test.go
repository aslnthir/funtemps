package conv

import (
	
	"math"
	"testing"
)
/*
	Vi skriver alle funksjoner i henhold til denne malen
*/

func withinTolerance(a, b, error float64) bool {
	// Først sjekker vi tallene
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	//vi skal dele med b
	if b == 0 {
		return difference < error
	}

	// sjekk den relative differanse
	return (difference / math.Abs(b)) < error
}

//  Farhenheit til Celsius
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Tester inputen 
	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}



//  Celcius til Farhenheit
func TestCelsiusToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Test og  inputen 
	tests := []test{
		{input: 56.67, want: 134},
	}
	for _, tc := range tests {
		got := CelsiusToFarhenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// Konverterer Kelvin til Celcius
func TestKelvinToCelcius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
// Test og  input
	tests := []test{
		{input: 329.82, want: 56.67},
	}
	for _, tc := range tests {
		got := KelvinToCelcius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//  Celcius til Kelvin
func TestCelciusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Test og  input
	tests := []test{
		{input: 56.67, want: 329.82},
	}
	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//  Kelvin til Farhenheit
func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Test og  input
	tests := []test{
		{input: 329.82, want: 134.00},
	}
	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//  Farhenheit til Kelvin
func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

// Test og  input
	tests := []test{
		{input: 134.00, want: 329.82},
	}
	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
