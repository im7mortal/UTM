package main

import (
	"github.com/im7mortal/UTM"
	"fmt"
)

func main() {

	easting, northing, zoneNumber, zoneLetter, err := UTM.FromLatLon(40.71435, -74.00597, false)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(
		fmt.Sprintf(
			"Easting: %f; Northing: %f; ZoneNumber: %d; ZoneLetter: %s;",
			easting,
			northing,
			zoneNumber,
			zoneLetter,
		))

	easting, northing, zoneNumber, zoneLetter, err = UTM.FromLatLon(40.71435, -74.00597, true)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(
		fmt.Sprintf(
			"Easting: %f; Northing: %f; ZoneNumber: %d; ZoneLetter: %s;",
			easting,
			northing,
			zoneNumber,
			zoneLetter,
		))

	latitude, longitude, err := UTM.ToLatLon(377486, 6296562, 30, "", true)
	fmt.Println(fmt.Sprintf("Latitude: %.5f; Longitude: %.5f;", latitude, longitude))

	latitude, longitude, err = UTM.ToLatLon(377486, 6296562, 30, "V")
	fmt.Println(fmt.Sprintf("Latitude: %.5f; Longitude: %.5f;", latitude, longitude))

}
