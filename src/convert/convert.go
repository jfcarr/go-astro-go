package convert

import (
	"math"
)

// RadiansToDegrees  Convert radians to equivalent degrees
func RadiansToDegrees(radians float32) float32 {
	return (radians * 180) / math.Pi
}

// DegreesToRadians  Convert degrees to equivalent radians
func DegreesToRadians(degrees float32) float32 {
	return (degrees * math.Pi) / 180
}
