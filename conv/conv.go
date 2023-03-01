package conv

// Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	
	return (value - 32) * 5 / 9
}


//Celcius til Farhenheit
func CelsiusToFarhenheit(value float64) float64 {

	return value*9/5 + 32
}

// Kelvin til Celcius
func KelvinToCelcius(value float64) float64 {

	return value - 273.15
}

// Celcius til Kelvin
func CelsiusToKelvin(value float64) float64 {

	return value + 273.15
}

//  Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 {

	return (value-273.15)*9/5 + 32
}

//  Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64 {

	return (value-32)*5/9 + 273.15
}
