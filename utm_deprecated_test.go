package UTM_test

import (
	"testing"
	"github.com/im7mortal/UTM"
)

func TestFromLatLonBadInputF(t *testing.T) {

	suppressPanic := func(i int) {
		defer func() {
			recover()
		}()
		UTM.FromLatLonF(badInputLatLon[i].Latitude, badInputLatLon[i].Longitude)
		t.Errorf("Expected panic. badInputLatLon TestFromLatLonBadInput case %d", i)
	}
	for i := range badInputLatLon {
		suppressPanic(i)
	}

	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			t.Errorf("not cover latitude %s", s)
		}
	}()
	longitude := 0.
	latitude := 0.
	for i := -8000.0; i < 8401.0; i++ {
		latitude = i / 100
		UTM.FromLatLonF(latitude, longitude)
	}
	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			t.Errorf("not cover longitude %s", s)
		}
	}()
	latitude = 0.
	for i := -18000.0; i < 18001.0; i++ {
		longitude = i / 100
		UTM.FromLatLonF(latitude, longitude)
	}
}

func TestToLatLonDeprecated(t *testing.T) {
	for i, data := range knownValues {
		result, err := data.UTM.ToLatLon()
		if err != nil {
			t.Fatal(err.Error())
		}
		if round(data.LatLon.Latitude) != round(result.Latitude) {
			t.Errorf("Latitude ToLatLon case %d", i)
		}
		if round(data.LatLon.Longitude) != round(result.Longitude) {
			t.Errorf("Longitude ToLatLon case %d", i)
		}
	}
}

func TestToLatLonWithDeprecated(t *testing.T) {
	for i, data := range knownValues {
		UTMwithNorthern := UTM.Coordinate{
			Easting:    data.UTM.Easting,
			Northing:   data.UTM.Northing,
			ZoneNumber: data.UTM.ZoneNumber,
		}

		result, err := UTMwithNorthern.ToLatLon(data.northern)
		if err != nil {
			t.Fatal(err.Error())
		}
		if round(data.LatLon.Latitude) != round(result.Latitude) {
			t.Errorf("Latitude TestToLatLonWithNorthern case %d", i)
		}
		if round(data.LatLon.Longitude) != round(result.Longitude) {
			t.Errorf("Longitude TestToLatLonWithNorthern case %d", i)
		}
	}
}
func TestFromLatLonF(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf(r.(string))
		}
	}()

	for i, data := range knownValues {
		e, n := UTM.FromLatLonF(data.LatLon.Latitude, data.LatLon.Longitude)
		if round(data.UTM.Easting) != round(e) {
			t.Errorf("Easting FromLatLon case %d", i)
		}
		if round(data.UTM.Northing) != round(n) {
			t.Errorf("Northing FromLatLon case %d", i)
		}
	}
}

// LatLon.FromLatLon and FromLatLon must calculate the same easting and northing
func TestFromLatLonAndF(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			t.Errorf("not cover longitude %s", s)
		}
	}()
	for i, data := range knownValues {
		result, err := data.LatLon.FromLatLon()
		if err != nil {
			t.Fatal(err.Error())
		}
		e, n := UTM.FromLatLonF(data.LatLon.Latitude, data.LatLon.Longitude)
		if round(e) != round(result.Easting) {
			t.Errorf("Easting FromLatLon case %d", i)
		}
		if round(n) != round(result.Northing) {
			t.Errorf("Northing FromLatLon case %d", i)
		}
	}
}
func TestFromLatLonDeprecated(t *testing.T) {

	for i, data := range knownValues {
		result, err := data.LatLon.FromLatLon()
		if err != nil {
			t.Fatal(err.Error())
		}
		if round(data.UTM.Easting) != round(result.Easting) {
			t.Errorf("Easting FromLatLon case %d", i)
		}
		if round(data.UTM.Northing) != round(result.Northing) {
			t.Errorf("Northing FromLatLon case %d", i)
		}
		if data.UTM.ZoneLetter != result.ZoneLetter {
			t.Errorf("ZoneLetter FromLatLon case %d", i)
		}
		if data.UTM.ZoneNumber != result.ZoneNumber {
			t.Errorf("ZoneNumber FromLatLon case %d", i)
		}
	}
}
