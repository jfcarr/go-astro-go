package main

import (
	"math"
)

func radiansToDegrees(radians float32) float32 {
	return (radians * 180) / math.Pi
}

func degreesToRadians(degrees float32) float32 {
	return (degrees * math.Pi) / 180
}
