package tempconv

// CToF converts a Celsius temp to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5/9) }

// CToK converts Celsius to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }