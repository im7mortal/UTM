package main

import (
	"github.com/im7mortal/UTM"
	"fmt"
)

func main() {

	latLon := UTM.LatLon{
		Latitude: 40.71435,
		Longitude: -74.00597,
	}
	result, err := latLon.FromLatLon()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(
		fmt.Sprintf(
			"Easting: %d; Northing: %d; ZoneNumber: %d; ZoneLetter: %s;",
			result.Easting,
			result.Northing,
			result.ZoneNumber,
			result.ZoneLetter,
		))

	coordinateUTM := UTM.Coordinate{
		Easting :        377486,
		Northing :        6296562,
		ZoneNumber :    30,
	}

	result1, err1 := coordinateUTM.ToLatLon(true)

	if err1 != nil {
		panic(err1.Error())
	}

	fmt.Println(fmt.Sprintf("Latitude: %.5f; Longitude: %.5f;", result1.Latitude, result1.Longitude))

	coordinateUTM.ZoneLetter = "V"

	result2, err2 := coordinateUTM.ToLatLon()
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println(fmt.Sprintf("Latitude: %.5f; Longitude: %.5f;", result2.Latitude, result2.Longitude))

}