package main

import (
	"fmt"

	"github.com/im7mortal/UTM"
)

func main() {
	easting, northing, zoneNumber, zoneLetter, err := UTM.FromLatLon(40.71435, -74.00597, false)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf(
		"Easting: %f; Northing: %f; ZoneNumber: %d; ZoneLetter: %s;\n",
		easting,
		northing,
		zoneNumber,
		zoneLetter,
	)

	easting, northing, zoneNumber, zoneLetter, err = UTM.FromLatLon(40.71435, -74.00597, true)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf(
		"Easting: %f; Northing: %f; ZoneNumber: %d; ZoneLetter: %s;\n",
		easting,
		northing,
		zoneNumber,
		zoneLetter,
	)

	latitude, longitude, err := UTM.ToLatLon(377486, 6296562, 30, "", true)
	fmt.Printf("Latitude: %.5f; Longitude: %.5f; Err: %s\n", latitude, longitude, err)

	latitude, longitude, err = UTM.ToLatLon(377486, 6296562, 30, "V")
	fmt.Printf("Latitude: %.5f; Longitude: %.5f;\f; Err: %s\n", latitude, longitude, err)
}
