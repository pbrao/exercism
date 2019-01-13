package space

// Age converts seconds to planet years
func Age(seconds float64, planet string) float64 {
	var factor float64
	Earth := 1.0
	Mercury := 0.2408467
	Venus := 0.61519726
	Mars := 1.8808158
	Jupiter := 11.862615
	Saturn := 29.447498
	Uranus := 84.016846
	Neptune := 164.79132

	switch planet {
	case "Earth":
		factor = Earth
	case "Mercury":
		factor = Mercury
	case "Venus":
		factor = Venus
	case "Mars":
		factor = Mars
	case "Jupiter":
		factor = Jupiter
	case "Saturn":
		factor = Saturn
	case "Uranus":
		factor = Uranus
	case "Neptune":
		factor = Neptune
	}
	return seconds / 31557600 / factor
}
