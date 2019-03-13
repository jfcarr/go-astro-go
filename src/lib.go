package main

// Observer  Information about an observer.
type Observer struct {
	Latitude  float32
	Longitude float32
}

// SkyObject  Information about an object in the sky.
type SkyObject struct {
	RightAscension float32
	Declination    float32
}

func (obs *Observer) setLatitude(degrees float32, minutes float32, cardinal string) {
	obs.Latitude = degrees + (minutes / 60)

	if cardinal == "S" {
		obs.Latitude = -(obs.Latitude)
	}
}

func (obs *Observer) setLongitude(degrees float32, minutes float32, cardinal string) {
	obs.Longitude = degrees + (minutes / 60)

	if cardinal == "W" {
		obs.Longitude = -(obs.Longitude)
	}
}

func (so *SkyObject) setRightAscension(hours float32, minutes float32, seconds float32) {
	so.RightAscension = (hours + (minutes / 60) + (seconds / 60 / 60)) * 15 // convert to degrees
}

func (so *SkyObject) setDeclination(degrees float32, minutes float32, seconds float32) {
	so.Declination = degrees + (minutes / 60) + (seconds / 60 / 60)
}
