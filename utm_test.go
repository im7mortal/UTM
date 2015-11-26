package UTM_test

import (
	"testing"
	"github.com/im7mortal/UTM"
	"math"
)

func round(f float64) float64 { return math.Floor(f + .5) }

type testData struct {
	LatLon   UTM.LatLon
	UTM      UTM.Coordinate
	northern bool
}

var knownValues = []testData{
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

func TestToLatLon(t *testing.T) {
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

func TestToLatLonWithNorthern(t *testing.T) {
	for i, data := range knownValues {
		UTMwithNorthern := UTM.Coordinate{
			Easting :		data.UTM.Easting,
			Northing :		data.UTM.Northing,
			ZoneNumber :	data.UTM.ZoneNumber,
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

func TestFromLatLon(t *testing.T) {

	for i, data := range knownValues {
		result, err := data.LatLon.FromLatLon()
		if err != nil {
			t.Fatal(err.Error())
		}
		if data.UTM.Easting != result.Easting {
			t.Errorf("Easting FromLatLon case %d", i)
		}
		if data.UTM.Northing != result.Northing {
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

var badInputLatLon = []UTM.LatLon{
	UTM.LatLon{-81, 0},
	UTM.LatLon{85, 0},
	UTM.LatLon{0, -185},
	UTM.LatLon{0, 185},
}

func TestFromLatLonBadInput(t *testing.T) {
	for i, data := range badInputLatLon {
		_, err := data.FromLatLon()
		if err == nil {
			t.Errorf("Expected error. badInputLatLon TestFromLatLonBadInput case %d", i)
		}
	}
	latLon := UTM.LatLon{}
	latLon.Longitude = 0
	for i := -8000.0; i < 8401.0; i++ {
		latLon.Latitude = i / 100
		_, err := latLon.FromLatLon()
		if err != nil {
			t.Errorf("not cover Latitude %d", i / 100)
		}
	}
	latLon.Latitude = 0
	for i := -18000.0; i < 18001.0; i++ {
		latLon.Longitude = i / 100
		_, err := latLon.FromLatLon()
		if err != nil {
			t.Errorf("not cover Longitude %d", i / 100)
		}
	}
}


var badInputToLatLon = []UTM.Coordinate{
	// out of range ZoneLetter
	UTM.Coordinate{377486, 6296562, 30, "Y"},
	UTM.Coordinate{377486, 6296562, 30, "B"},
	UTM.Coordinate{377486, 6296562, 30, "I"},
	UTM.Coordinate{377486, 6296562, 30, "i"},
	UTM.Coordinate{377486, 6296562, 30, "O"},
	UTM.Coordinate{377486, 6296562, 30, "o"},
	// out of range ZoneNumber
	UTM.Coordinate{377486, 6296562, 0, "V"},
	UTM.Coordinate{377486, 6296562, 61, "V"},
	// out of range Easting
	UTM.Coordinate{1000000, 6296562, 30, "V"},
	UTM.Coordinate{99999, 6296562, 30, "V"},
	// out of range Northing
	UTM.Coordinate{377486, 10000001, 30, "V"},
	UTM.Coordinate{377486, -1, 30, "V"},
}

func TestToLatLonBadInput(t *testing.T) {
	for i, data := range badInputToLatLon {
		_, err := data.ToLatLon()
		if err == nil {
			t.Errorf("Expected error. badInputToLatLon TestToLatLonBadInput case %d", i)
		}
	}
	coordinate := UTM.Coordinate{
		Easting:	377486,
		Northing:	6296562,
		ZoneNumber:	30,
	}
	_, err := coordinate.ToLatLon()
	if err == nil {
		t.Errorf("Expected error. too few arguments")
	}
	coordinate.ZoneLetter = "V"
	_, err = coordinate.ToLatLon(true)
	if err == nil {
		t.Errorf("Expected error. too many arguments")
	}
	letters := []string{
		"X", "W", "V", "U", "T", "S", "R", "Q", "P", "N", "M", "L", "K", "J", "H", "G", "F", "E", "D", "C",
		"x", "w", "v", "u", "t", "s", "r", "q", "p", "n", "m", "l", "k", "j", "h", "g", "f", "e", "d", "c",
	}

	for _, letter := range letters{
		coordinate.ZoneLetter = letter
		_, err := coordinate.ToLatLon()
		if err != nil {
			t.Errorf("letter isn't covered. %s", letter)
		}
	}
}