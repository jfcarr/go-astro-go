package skyobject

// SkyObject  Information about an object in the sky.
type SkyObject struct {
	RightAscension float32
	Declination    float32
}

// SetRightAscension  Set right ascension value for SkyObject
func (so *SkyObject) SetRightAscension(hours float32, minutes float32, seconds float32) {
	so.RightAscension = (hours + (minutes / 60) + (seconds / 60 / 60)) * 15 // convert to degrees
}

// SetDeclination  Set declination value for SkyObject
func (so *SkyObject) SetDeclination(degrees float32, minutes float32, seconds float32) {
	so.Declination = degrees + (minutes / 60) + (seconds / 60 / 60)
}
