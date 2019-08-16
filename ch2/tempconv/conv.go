package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CTof(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converst Fahrenheit temperature to Celsius.
func FToc(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
