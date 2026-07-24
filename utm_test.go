package UTM_test

import (
	"math"
	"testing"

	"github.com/im7mortal/UTM"
)

func round(f float64) float64 { return math.Floor(f + .5) }

// emulation for test only.
type testLatLon struct {
	Latitude  float64
	Longitude float64
}

// emulation for test only.
type testCoordinate struct {
	Easting    float64
	Northing   float64
	ZoneNumber int
	ZoneLetter string
}

type testData struct {
	LatLon   testLatLon
	UTM      testCoordinate
	northern bool
}

func getTestValues() []testData {
	return []testData{
		// Aachen, Germany
		{
			testLatLon{50.77535, 6.08389},
			testCoordinate{294409, 5628898, 32, "U"},
			true,
		},
		// New York, USA
		{
			testLatLon{40.71435, -74.00597},
			testCoordinate{583960, 4507523, 18, "T"},
			true,
		},
		// Wellington, New Zealand
		{
			testLatLon{-41.28646, 174.77624},
			testCoordinate{313784, 5427057, 60, "G"},
			false,
		},
		// Capetown, South Africa
		{
			testLatLon{-33.92487, 18.42406},
			testCoordinate{261878, 6243186, 34, "H"},
			false,
		},
		// Mendoza, Argentina
		{
			testLatLon{-32.89018, -68.84405},
			testCoordinate{514586, 6360877, 19, "H"},
			false,
		},
		// Fairbanks, Alaska, USA
		{
			testLatLon{64.83778, -147.71639},
			testCoordinate{466013, 7190568, 6, "W"},
			true,
		},
		// Ben Nevis, Scotland, UK
		{
			testLatLon{56.79680, -5.00601},
			testCoordinate{377486, 6296562, 30, "V"},
			true,
		},
		// Latitude 84
		{
			testLatLon{84, -5.00601},
			testCoordinate{476594, 9328501, 30, "X"},
			true,
		},
	}
}

func TestToLatLon(t *testing.T) {
	t.Parallel()

	for i, data := range getTestValues() {
		latitude, longitude, err := UTM.ToLatLon(
			data.UTM.Easting,
			data.UTM.Northing,
			data.UTM.ZoneNumber,
			data.UTM.ZoneLetter)
		if err != nil {
			t.Fatal(err.Error())
		}

		if round(data.LatLon.Latitude) != round(latitude) {
			t.Errorf("Latitude ToLatLon case %d", i)
		}

		if round(data.LatLon.Longitude) != round(longitude) {
			t.Errorf("Longitude ToLatLon case %d", i)
		}
	}
}

func TestToLatLonWithNorthern(t *testing.T) {
	t.Parallel()

	emptyZoneLetter := ""

	for i, data := range getTestValues() {
		latitude, longitude, err := UTM.ToLatLon(
			data.UTM.Easting,
			data.UTM.Northing,
			data.UTM.ZoneNumber,
			emptyZoneLetter,
			data.northern)
		if err != nil {
			t.Fatal(err.Error())
		}

		if round(data.LatLon.Latitude) != round(latitude) {
			t.Errorf("Latitude TestToLatLonWithNorthern case %d", i)
		}

		if round(data.LatLon.Longitude) != round(longitude) {
			t.Errorf("Longitude TestToLatLonWithNorthern case %d", i)
		}
	}
}

func TestFromLatLon(t *testing.T) {
	t.Parallel()

	for i, data := range getTestValues() {
		easting, northing, zoneNumber, zoneLetter, err := UTM.FromLatLon(data.LatLon.Latitude, data.LatLon.Longitude, false)
		if err != nil {
			t.Fatal(err.Error())
		}

		if round(data.UTM.Easting) != round(easting) {
			t.Errorf("Easting FromLatLon case %d", i)
		}

		if round(data.UTM.Northing) != round(northing) {
			t.Errorf("Northing FromLatLon case %d", i)
		}

		if data.UTM.ZoneLetter != zoneLetter {
			t.Errorf("ZoneLetter FromLatLon case %d", i)
		}

		if data.UTM.ZoneNumber != zoneNumber {
			t.Errorf("ZoneNumber FromLatLon case %d", i)
		}
	}
}

func getBadInputLatLon() []testLatLon {
	return []testLatLon{
		{-81, 0},
		{85, 0},
		{0, -185},
		{0, 185},
	}
}

func TestFromLatLonBadInput(t *testing.T) {
	t.Parallel()

	for i, data := range getBadInputLatLon() {
		_, _, _, _, err := UTM.FromLatLon(data.Latitude, data.Longitude, false)
		if err == nil {
			t.Errorf("Expected error. badInputLatLon TestFromLatLonBadInput case %d", i)
		}
		if _, ok := err.(UTM.InputError); !ok {
			t.Error("Type of error must be UTM.InputError.")
		}
	}

	latLon := testLatLon{}
	latLon.Longitude = 0

	for i := -8000.0; i < 8401.0; i++ {
		latLon.Latitude = i / 100
		_, _, _, _, err := UTM.FromLatLon(latLon.Latitude, latLon.Longitude, false)
		if err != nil {
			t.Errorf("not cover Latitude %f", i/100)
		}
	}

	latLon.Latitude = 0
	for i := -18000.0; i < 18001.0; i++ {
		latLon.Longitude = i / 100
		_, _, _, _, err := UTM.FromLatLon(latLon.Latitude, latLon.Longitude, false)
		if err != nil {
			t.Errorf("not cover Longitude %f", i/100)
		}
	}
}

func getBadInputToLatLon() []testCoordinate {
	return []testCoordinate{
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

func TestToLatLonBadInput(t *testing.T) {
	t.Parallel()

	var err error
	for i, data := range getBadInputToLatLon() {
		_, _, err = UTM.ToLatLon(data.Easting, data.Northing, data.ZoneNumber, data.ZoneLetter)
		if err == nil {
			t.Errorf("Expected error. badInputToLatLon TestToLatLonBadInput case %d", i)
		}
		if _, ok := err.(UTM.InputError); !ok {
			t.Error("Type of error must be UTM.InputError.")
		}
	}

	coordinate := testCoordinate{
		Easting:    377486,
		Northing:   6296562,
		ZoneNumber: 30,
	}

	_, _, err = UTM.ToLatLon(coordinate.Easting, coordinate.Northing, coordinate.ZoneNumber, "")
	if err == nil {
		t.Error("Expected error. too few arguments")
	}
	if _, ok := err.(UTM.InputError); !ok {
		t.Error("Type of error must be UTM.InputError.")
	}

	coordinate.ZoneLetter = "V"

	_, _, err = UTM.ToLatLon(coordinate.Easting, coordinate.Northing, coordinate.ZoneNumber, coordinate.ZoneLetter, true)
	if err == nil {
		t.Error("Expected error. too many arguments")
	}
	if _, ok := err.(UTM.InputError); !ok {
		t.Error("Type of error must be UTM.InputError.")
	}

	letters := []string{
		"X", "W", "V", "U", "T", "S", "R", "Q", "P", "N", "M", "L", "K", "J", "H", "G", "F", "E", "D", "C",
		"x", "w", "v", "u", "t", "s", "r", "q", "p", "n", "m", "l", "k", "j", "h", "g", "f", "e", "d", "c",
	}

	for _, letter := range letters {
		coordinate.ZoneLetter = letter

		_, _, err = UTM.ToLatLon(coordinate.Easting, coordinate.Northing, coordinate.ZoneNumber, coordinate.ZoneLetter)
		if err != nil {
			t.Errorf("letter isn't covered. %s", letter)
		}
	}
}

// TestFromLatLonAntimeridianZoneWrap guards the antimeridian boundary: lon == 180
// must map to zone 1 (wrap), not the out-of-range zone 61, matching the Python
// `utm` reference. Both the equator and the Svalbard (72-84N) band are covered.
// northern=false is used so the natural MGRS zone letter is returned.
func TestFromLatLonAntimeridianZoneWrap(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                string
		lat, lon            float64
		wantZone            int
		wantLetter          string
		wantRoundTripErrNil bool
	}{
		{"equator_lon180", 0.0, 180.0, 1, "N", true},
		{"svalbard_lon180", 80.0, 180.0, 1, "X", true},
		{"equator_lon_neg180", 0.0, -180.0, 1, "N", true},
		{"equator_lon179.99", 0.0, 179.99, 60, "N", true},
	}

	for _, c := range cases {
		easting, northing, zone, letter, err := UTM.FromLatLon(c.lat, c.lon, false)
		if err != nil {
			t.Fatalf("%s: FromLatLon returned unexpected error: %v", c.name, err)
		}
		if zone != c.wantZone {
			t.Errorf("%s: zone = %d, want %d", c.name, zone, c.wantZone)
		}
		if letter != c.wantLetter {
			t.Errorf("%s: letter = %q, want %q", c.name, letter, c.wantLetter)
		}
		if !(1 <= zone && zone <= 60) {
			t.Errorf("%s: zone %d out of valid range [1,60]", c.name, zone)
		}
		// Forward output must round-trip through the library's own ToLatLon.
		_, _, rtErr := UTM.ToLatLon(easting, northing, zone, letter)
		if c.wantRoundTripErrNil && rtErr != nil {
			t.Errorf("%s: ToLatLon round-trip failed on forward output (zone=%d): %v",
				c.name, zone, rtErr)
		}
	}
}

// TestFromLatLonNorwayLat64Boundary guards the Norway zone-32V extension boundary.
// MGRS band V is half-open [56, 64); at lat == 64 (band W) the extension does NOT
// apply and the default zone (31 for lon in [3, 6)) is correct. lat == 63.99 stays
// in band V and keeps zone 32. The projected easting at (64, 3) is zone 31's central
// meridian (3E), so easting ~= 500000 m, matching pyproj EPSG:32631.
func TestFromLatLonNorwayLat64Boundary(t *testing.T) {
	t.Parallel()

	// (64, 3): band W, default zone 31, easting near 500000 (central meridian 3E).
	e, n, zone, letter, err := UTM.FromLatLon(64.0, 3.0, false)
	if err != nil {
		t.Fatalf("FromLatLon(64,3) returned unexpected error: %v", err)
	}
	if zone != 31 {
		t.Errorf("FromLatLon(64,3) zone = %d, want 31 (band W, no Norway extension)", zone)
	}
	if letter != "W" {
		t.Errorf("FromLatLon(64,3) letter = %q, want \"W\"", letter)
	}
	if math.Abs(e-500000) > 1.0 {
		t.Errorf("FromLatLon(64,3) easting = %.4f, want ~500000 (pyproj EPSG:32631)", e)
	}
	if math.Abs(n-7097014.16) > 1.0 {
		t.Errorf("FromLatLon(64,3) northing = %.4f, want ~7097014.16 (pyproj EPSG:32631)", n)
	}

	// (63.99, 3): still band V, Norway extension applies -> zone 32.
	_, _, zone2, letter2, err2 := UTM.FromLatLon(63.99, 3.0, false)
	if err2 != nil {
		t.Fatalf("FromLatLon(63.99,3) returned unexpected error: %v", err2)
	}
	if zone2 != 32 {
		t.Errorf("FromLatLon(63.99,3) zone = %d, want 32 (band V, Norway extension)", zone2)
	}
	if letter2 != "V" {
		t.Errorf("FromLatLon(63.99,3) letter = %q, want \"V\"", letter2)
	}
}
