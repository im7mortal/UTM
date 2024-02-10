package UTM_test

import (
	"testing"

	"github.com/im7mortal/UTM"
)

type testDataDeprecated struct {
	LatLon   UTM.LatLon
	UTM      UTM.Coordinate
	northern bool
}

func getTestValuesDeprecated() []testDataDeprecated {
	return []testDataDeprecated{
		// Aachen, Germany
		{
			UTM.LatLon{50.77535, 6.08389},
			UTM.Coordinate{294409, 5628898, 32, "U"},
			true,
		},
		// New York, USA
		{
			UTM.LatLon{40.71435, -74.00597},
			UTM.Coordinate{583960, 4507523, 18, "T"},
			true,
		},
		// Wellington, New Zealand
		{
			UTM.LatLon{-41.28646, 174.77624},
			UTM.Coordinate{313784, 5427057, 60, "G"},
			false,
		},
		// Capetown, South Africa
		{
			UTM.LatLon{-33.92487, 18.42406},
			UTM.Coordinate{261878, 6243186, 34, "H"},
			false,
		},
		// Mendoza, Argentina
		{
			UTM.LatLon{-32.89018, -68.84405},
			UTM.Coordinate{514586, 6360877, 19, "H"},
			false,
		},
		// Fairbanks, Alaska, USA
		{
			UTM.LatLon{64.83778, -147.71639},
			UTM.Coordinate{466013, 7190568, 6, "W"},
			true,
		},
		// Ben Nevis, Scotland, UK
		{
			UTM.LatLon{56.79680, -5.00601},
			UTM.Coordinate{377486, 6296562, 30, "V"},
			true,
		},
	}
}

func getBadInputLatLonDeprecated() []UTM.LatLon {
	return []UTM.LatLon{
		{-81, 0},
		{85, 0},
		{0, -185},
		{0, 185},
	}
}

func getBadInputToLatLonDeprecated() []UTM.Coordinate {
	return []UTM.Coordinate{
		// out of range ZoneLetter
		{377486, 6296562, 30, "Y"},
		{377486, 6296562, 30, "B"},
		{377486, 6296562, 30, "I"},
		{377486, 6296562, 30, "i"},
		{377486, 6296562, 30, "O"},
		{377486, 6296562, 30, "o"},
		// out of range ZoneNumber
		{377486, 6296562, 0, "V"},
		{377486, 6296562, 61, "V"},
		// out of range Easting
		{1000000, 6296562, 30, "V"},
		{99999, 6296562, 30, "V"},
		// out of range Northing
		{377486, 10000001, 30, "V"},
		{377486, -1, 30, "V"},
	}
}

func TestFromLatLonBadInputF(t *testing.T) {
	t.Parallel()

	suppressPanic := func(i int) {
		defer func() {
			_ = recover()
		}()
		UTM.FromLatLonF(getBadInputLatLonDeprecated()[i].Latitude, getBadInputLatLonDeprecated()[i].Longitude)
		t.Errorf("Expected panic. badInputLatLon TestFromLatLonBadInput case %d", i)
	}

	for i := range getBadInputLatLonDeprecated() {
		suppressPanic(i)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("not cover latitude %s", r)
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
			t.Errorf("not cover longitude %s", r)
		}
	}()

	latitude = 0.

	for i := -18000.0; i < 18001.0; i++ {
		longitude = i / 100
		UTM.FromLatLonF(latitude, longitude)
	}
}

func TestToLatLonDeprecated(t *testing.T) {
	t.Parallel()

	for i, data := range getTestValuesDeprecated() {
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
	t.Parallel()

	for i, data := range getTestValuesDeprecated() {
		UTMWithNorthern := UTM.Coordinate{
			Easting:    data.UTM.Easting,
			Northing:   data.UTM.Northing,
			ZoneNumber: data.UTM.ZoneNumber,
		}

		result, err := UTMWithNorthern.ToLatLon(data.northern)
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
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("%s", r)
		}
	}()

	for i, data := range getTestValuesDeprecated() {
		e, n := UTM.FromLatLonF(data.LatLon.Latitude, data.LatLon.Longitude)

		if round(data.UTM.Easting) != round(e) {
			t.Errorf("Easting FromLatLon case %d", i)
		}

		if round(data.UTM.Northing) != round(n) {
			t.Errorf("Northing FromLatLon case %d", i)
		}
	}
}

// LatLon.FromLatLon and FromLatLon must calculate the same easting and northing.
func TestFromLatLonAndF(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("not cover longitude %s", r)
		}
	}()

	for i, data := range getTestValuesDeprecated() {
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
	t.Parallel()

	for i, data := range getTestValuesDeprecated() {
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

func TestToLatLonBadInputDeprecated(t *testing.T) {
	t.Parallel()

	for i, data := range getBadInputToLatLonDeprecated() {
		_, err := data.ToLatLon()
		if err == nil {
			t.Errorf("Expected error. badInputToLatLon TestToLatLonBadInput case %d", i)
		}
	}

	coordinate := UTM.Coordinate{
		Easting:    377486,
		Northing:   6296562,
		ZoneNumber: 30,
	}
	_, err := coordinate.ToLatLon()

	if err == nil {
		t.Error("Expected error. too few arguments")
	}

	coordinate.ZoneLetter = "V"
	_, err = coordinate.ToLatLon(true)

	if err == nil {
		t.Error("Expected error. too many arguments")
	}

	letters := []string{
		"X", "W", "V", "U", "T", "S", "R", "Q", "P", "N", "M", "L", "K", "J", "H", "G", "F", "E", "D", "C",
		"x", "w", "v", "u", "t", "s", "r", "q", "p", "n", "m", "l", "k", "j", "h", "g", "f", "e", "d", "c",
	}

	for _, letter := range letters {
		coordinate.ZoneLetter = letter

		_, err = coordinate.ToLatLon()
		if err != nil {
			t.Errorf("letter isn't covered. %s", letter)
		}
	}
}
