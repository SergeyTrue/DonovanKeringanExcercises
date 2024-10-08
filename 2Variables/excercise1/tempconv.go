package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroCelsius Celsius = -273.15
	FreezingCelsius     Celsius = 0
	BoilingCelsius      Celsius = 100
)

func (c Celsius) String() string {

	return fmt.Sprintf("%g C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
func CToK(c Celsius) Kelvin {
	return Kelvin((c + 273.15))
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}
