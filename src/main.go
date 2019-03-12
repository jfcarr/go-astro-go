package main

import (
	"fmt"
)

func main() {
	var myObserver Observer
	var mySkyObject SkyObject

	myObserver.setLatitude(52, 30, "N")
	myObserver.setLongitude(1, 55, "W")

	mySkyObject.setRightAscension(16, 41.7, 0)
	mySkyObject.setDeclination(36, 28, 0)

	fmt.Printf("Latitude is %.2f, Longitude is %.2f\n", myObserver.Latitude, myObserver.Longitude)
	fmt.Printf("Right Ascension is %.2f, Declination is %.2f\n", mySkyObject.RightAscension, mySkyObject.Declination)
}
