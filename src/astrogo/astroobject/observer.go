package astroobject

// Observer  Information about an observer.
type Observer struct {
	Latitude  float32
	Longitude float32
}

// SetLatitude  Set Observer's latitude.
func (obs *Observer) SetLatitude(degrees float32, minutes float32, cardinal string) {
	obs.Latitude = degrees + (minutes / 60)

	if cardinal == "S" {
		obs.Latitude = -(obs.Latitude)
	}
}

// SetLongitude  Set Observer's longitude.
func (obs *Observer) SetLongitude(degrees float32, minutes float32, cardinal string) {
	obs.Longitude = degrees + (minutes / 60)

	if cardinal == "W" {
		obs.Longitude = -(obs.Longitude)
	}
}
