package main

import (
	"./astrogo/astroobject"
	"./astrogo/astroutil"
	"fmt"
)

func main() {
	var myObserver astroobject.Observer
	var mySkyObject astroobject.SkyObject

	myObserver.SetLatitude(52, 30, "N")
	myObserver.SetLongitude(1, 55, "W")

	mySkyObject.SetRightAscension(16, 41.7, 0)
	mySkyObject.SetDeclination(36, 28, 0)

	fmt.Printf("Latitude is %.2f, Longitude is %.2f\n", myObserver.Latitude, myObserver.Longitude)
	fmt.Printf("Right Ascension is %.2f, Declination is %.2f\n", mySkyObject.RightAscension, mySkyObject.Declination)

	var degrees float32 = 180
	var radians float32 = 1

	fmt.Printf("%.2f degrees is %.2f radians\n", degrees, astroutil.DegreesToRadians(degrees))
	fmt.Printf("%.2f radians is %.2f degrees\n", radians, astroutil.RadiansToDegrees(radians))
}
