package UTM

import (
	"testing"
	"math"
//	"fmt"
)

type testData struct {
	LatLon   LAT_LON
	UTM      UTM_COORDINATE
	northern bool
}


var known_values []testData = []testData{
	// Aachen, Germany
	{
		LAT_LON{50.77535, 6.08389},
		UTM_COORDINATE{294409, 5628898, 32, 'U'},
		true,
	},
	// New York, USA
	{
		LAT_LON{40.71435, -74.00597},
		UTM_COORDINATE{583960, 4507523, 18, 'T'},
		true,
	},
	// Wellington, New Zealand
	{
		LAT_LON{-41.28646, 174.77624},
		UTM_COORDINATE{313784, 5427057, 60, 'G'},
		false,
	},
	// Capetown, South Africa
	{
		LAT_LON{-33.92487, 18.42406},
		UTM_COORDINATE{261878, 6243186, 34, 'H'},
		false,
	},
	// Mendoza, Argentina
	{
		LAT_LON{-32.89018, -68.84405},
		UTM_COORDINATE{514586, 6360877, 19, 'H'}, // todo revert to 'h' for test
		false,
	},
	// Fairbanks, Alaska, USA
	{
		LAT_LON{64.83778, -147.71639},
		UTM_COORDINATE{466013, 7190568, 6, 'W'},
		true,
	},
	// Ben Nevis, Scotland, UK
	{
		LAT_LON{56.79680, -5.00601},
		UTM_COORDINATE{377486, 6296562, 30, 'V'},
		true,
	},
}



func TestTO_LATLON(t *testing.T) {
	for i, data := range known_values {
		result, err := data.UTM.TO_LATLON()
		if err != nil {
			t.Fatal(err.Error())
		}
		if Round(data.LatLon.Latitude) != Round(result.Latitude) {
			t.Errorf("Latitude TO_LATLON case %d", i)
		}
		if Round(data.LatLon.Longitude) != Round(result.Longitude) {
			t.Errorf("Longitude TO_LATLON case %d", i)
		}
	}
}


func TestFROM_LATLON(t *testing.T) {

	for i, data := range known_values {
		result, err := data.LatLon.FROM_LATLON()
		if err != nil {
			t.Fatal(err.Error())
		}
		if Round(data.UTM.Easting) != Round(result.Easting) {
			t.Errorf("Easting FROM_LATLON case %d", i)
		}
		if Round(data.UTM.Northing) != Round(result.Northing) {
			t.Errorf("Northing FROM_LATLON case %d", i)
		}
		if data.UTM.Zone_letter != result.Zone_letter {
			t.Errorf("Zone_letter FROM_LATLON case %d", i)
		}
		if data.UTM.Zone_number != result.Zone_number {
			t.Errorf("Zone_number FROM_LATLON case %d", i)
		}
	}
}

func Round(f float64) float64 {
	return math.Floor(f + .5)
}