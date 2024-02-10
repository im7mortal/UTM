package UTM_test

import (
	"testing"

	"github.com/im7mortal/UTM"
)

func getBenchmarkCoordinate() testCoordinate {
	return testCoordinate{466013, 7190568, 6, "W"}
}

func BenchmarkToLatLon(b *testing.B) {
	var err error
	coordinate := getBenchmarkCoordinate()
	for i := 0; i < b.N; i++ {
		if _, _, err = UTM.ToLatLon(coordinate.Easting, coordinate.Northing, coordinate.ZoneNumber, coordinate.ZoneLetter); err != nil {
			b.Fatal("benchmark fatal BenchmarkToLatLon")
		}
	}
}

func BenchmarkToLatLonWithNorthern(b *testing.B) {
	var err error
	coordinate := getBenchmarkCoordinate()
	coordinate.ZoneLetter = ""
	for i := 0; i < b.N; i++ {
		if _, _, err = UTM.ToLatLon(coordinate.Easting, coordinate.Northing, coordinate.ZoneNumber, coordinate.ZoneLetter, true); err != nil {
			b.Fatal("benchmark fatal BenchmarkToLatLonWithNorthern")
		}
	}
}

func getLatLon() testLatLon {
	return testLatLon{64.83778, -147.71639}
}

func BenchmarkFromLatLon(b *testing.B) {
	var err error
	latLon := getLatLon()
	for i := 0; i < b.N; i++ {
		if _, _, _, _, err = UTM.FromLatLon(latLon.Latitude, latLon.Longitude, false); err != nil {
			b.Fatal("benchmark fatal BenchmarkFromLatLon")
		}
	}
}
