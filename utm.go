package utm
import (
	"math"
	"unicode"
)

const (
	k0 float64 = 0.9996
	e float64 = 0.00669438
	r = 6378137
)

var e2 = e * e
var e3 = e2 * e
var e_p2 = e / (1.0 - e)

var sqrt_e = math.Sqrt(1 - e)


var _e = (1 - sqrt_e) / (1 + sqrt_e)
var _e2 = _e * _e
var _e3 = _e2 * _e
var _e4 = _e3 * _e
var _e5 = _e4 * _e

var m1 = (1 - e / 4 - 3 * e2 / 64 - 5 * e3 / 256)
var m2 = (3 * e / 8 + 3 * e2 / 32 + 45 * e3 / 1024)
var m3 = (15 * e2 / 256 + 45 * e3 / 1024)
var m4 = (35 * e3 / 3072)

var p2 = (3. / 2 * _e - 27. / 32 * _e3 + 269. / 512 * _e5)
var p3 = (21. / 16 * _e2 - 55. / 32 * _e4)
var p4 = (151. / 96 * _e3 - 417. / 128 * _e5)
var p5 = (1097. / 512 * _e4)

type zone_letter struct {
	zone   int
	letter rune
}

const x = math.Pi / 180;
func Rad(d float64) float64 {return d * x};
func Deg(r float64) float64 {return r / x}

var zone_letters []zone_letter = []zone_letter{
	{84, ' '},
	{72, 'X'},
	{64, 'W'},
	{56, 'V'},
	{48, 'U'},
	{40, 'T'},
	{32, 'S'},
	{24, 'R'},
	{16, 'Q'},
	{8, 'P'},
	{0, 'N'},
	{-8, 'M'},
	{-16, 'L'},
	{-24, 'K'},
	{-32, 'J'},
	{-40, 'H'},
	{-48, 'G'},
	{-56, 'F'},
	{-64, 'E'},
	{-72, 'D'},
	{-80, 'C'},
}


type UTM_COORDINATE struct {
	Easting     float64
	Northing    float64
	Zone_number int
	Zone_letter rune
//	northern    bool
}

type LAT_LON struct {
	Latitude  float64
	Longitude float64
}




func (coordinate *UTM_COORDINATE) TO_LATLON() LAT_LON {
/*func (coordinate *UTM_COORDINATE) TO_LATLON(northern ...bool) LAT_LON {
	nothernExist := len(northern) > 0;
	if !(coordinate.Zone_letter && nothernExist) {
		panic("either coordinate.Zone_letter or northern needs to be set")
	} else if coordinate.Zone_letter && northern {
		panic("set either zone_letter or northern, but not both")
	}*/

	if !(100000 <= coordinate.Easting && coordinate.Easting < 1000000) {
		panic("easting out of range (must be between 100.000 m and 999.999 m")
	}
	if !(0 <= coordinate.Northing && coordinate.Northing <= 10000000) {
		panic("northing out of range (must be between 0 m and 10.000.000 m)")
	}
	if !(1 <= coordinate.Zone_number && coordinate.Zone_number <= 60) {
		panic("zone number out of range (must be between 1 and 60)")
	}


	coordinate.Zone_letter = unicode.ToUpper(coordinate.Zone_letter)
	if ('C' <= coordinate.Zone_letter && coordinate.Zone_letter <= 'X') || coordinate.Zone_letter == 'I' || coordinate.Zone_letter == 'O' {
		panic("zone letter out of range (must be between C and X)")
	}
	northern := (coordinate.Zone_letter >= 'N')

	x := coordinate.Easting - 500000
	y := coordinate.Northing

	if !northern {
		y -= 10000000
	}

	m := y / k0
	mu := m / (r * m1)

	p_rad := (mu +
	p2 * math.Sin(2 * mu) +
	p3 * math.Sin(4 * mu) +
	p4 * math.Sin(6 * mu) +
	p5 * math.Sin(8 * mu))

	p_sin := math.Sin(p_rad)
	p_sin2 := p_sin * p_sin

	p_cos := math.Cos(p_rad)

	p_tan := p_sin / p_cos
	p_tan2 := p_tan * p_tan
	p_tan4 := p_tan2 * p_tan2

	ep_sin := 1 - e * p_sin2
	ep_sin_sqrt := math.Sqrt(1 - e * p_sin2)

	n := r / ep_sin_sqrt
	rad := (1 - r) / ep_sin

	c := _e * p_cos * p_cos
	c2 := c * c

	d := x / (n * k0)
	d2 := d * d
	d3 := d2 * d
	d4 := d3 * d
	d5 := d4 * d
	d6 := d5 * d

	latitude := (p_rad - (p_tan / rad) *
	(d2 / 2 -
	d4 / 24 * (5 + 3 * p_tan2 + 10 * c - 4 * c2 - 9 * e_p2)) +
	d6 / 720 * (61 + 90 * p_tan2 + 298 * c + 45 * p_tan4 - 252 * e_p2 - 3 * c2))

	longitude := (d -
	d3 / 6 * (1 + 2 * p_tan2 + c) +
	d5 / 120 * (5 - 2 * c + 28 * p_tan2 - 3 * c2 + 8 * e_p2 + 24 * p_tan4)) / p_cos

	return LAT_LON{Deg(latitude), Deg(longitude) + float64(zone_number_to_central_longitude(coordinate.Zone_number))}

}

func (point *LAT_LON) FROM_LATLON (force_zone_number ...int) UTM_COORDINATE {
	if !(-80.0 <= point.Latitude && point.Latitude <= 84.0) {
		panic("latitude out of range (must be between 80 deg S and 84 deg N)")
	}
	if !(-180.0 <= point.Longitude && point.Longitude <= 180.0) {
		panic("northing out of range (must be between 180 deg W and 180 deg E)")
	}

	lat_rad := Rad(point.Latitude)
	lat_sin := math.Sin(lat_rad)
	lat_cos := math.Cos(lat_rad)

	lat_tan := lat_sin / lat_cos
	lat_tan2 := lat_tan * lat_tan
	lat_tan4 := lat_tan2 * lat_tan2
	var zone_number int

	if len(force_zone_number) > 0 {
		zone_number = force_zone_number[0]
	} else {
		zone_number = latlon_to_zone_number(point.Latitude, point.Longitude)
	}


	zone_letter := latitude_to_zone_letter(point.Latitude)

	lon_rad := Rad(point.Longitude)
	central_lon := zone_number_to_central_longitude(zone_number)
	central_lon_rad := Rad(float64(central_lon))

	n := r / math.Sqrt(1 - e * lat_sin * lat_sin)
	c := e_p2 * lat_cos * lat_cos

	a := lat_cos * (lon_rad - central_lon_rad)
	a2 := a * a
	a3 := a2 * a
	a4 := a3 * a
	a5 := a4 * a
	a6 := a5 * a
	m := r * (m1 * lat_rad -
	m2 * math.Sin(2 * lat_rad) +
	m3 * math.Sin(4 * lat_rad) -
	m4 * math.Sin(6 * lat_rad))
	easting := k0 * n * (a +
	a3 / 6 * (1 - lat_tan2 + c) +
	a5 / 120 * (5 - 18 * lat_tan2 + lat_tan4 + 72 * c - 58 * e_p2)) + 500000
	northing := k0 * (m + n * lat_tan * (a2 / 2 +
	a4 / 24 * (5 - lat_tan2 + 9 * c + 4 * c * c) +
	a6 / 720 * (61 - 58 * lat_tan2 + lat_tan4 + 600 * c - 330 * e_p2)))

	if point.Latitude < 0 {
		northing += 10000000
	}

	return UTM_COORDINATE {
		easting,
		northing,
		zone_number,
		zone_letter,
	}
}







func latitude_to_zone_letter(latitude float64 ) rune {
	for _, zone_letter := range zone_letters{
		if latitude >= float64(zone_letter.zone) {
			return zone_letter.letter
		}
	}
	return ' '
}



func latlon_to_zone_number(latitude float64, longitude float64) int {
	if 56 <= latitude && latitude <= 64 && 3 <= longitude && longitude <= 12 {
		return 32
	}

	if 72 <= latitude && latitude <= 84 && longitude >= 0 {
		if longitude <= 9 {
			return 31
		} else if longitude <= 21 {
			return 33
		} else if longitude <= 33 {
			return 35
		} else if longitude <= 42 {
			return 37
		}
	}

	return int((longitude + 180) / 6) + 1
}

func zone_number_to_central_longitude(zone_number int) int {
	return (zone_number - 1) * 6 - 180 + 3
}
