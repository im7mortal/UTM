// Package UTM is bidirectional UTM-WGS84 converter for golang
package UTM

import (
	"math"
)

//FromLatLon convert a latitude and longitude to Universal Transverse Mercator coordinates
//version with micro optimizations
//panic instead errors
func FromLatLon(lat, lon float64) (easting, northing float64) {
	if !(-80.0 <= lat && lat <= 84.0) {
		panic("latitude out of range (must be between 80 deg S and 84 deg N)")
	}
	if !(-180.0 <= lon && lon <= 180.0) {
		panic("longitude out of range (must be between 180 deg W and 180 deg E)")
	}

	lat_rad := rad(lat)
	lat_sin := math.Sin(lat_rad)
	lat_cos := math.Cos(lat_rad)

	lat_tan := lat_sin / lat_cos
	lat_tan2 := lat_tan * lat_tan
	lat_tan4 := lat_tan2 * lat_tan2
	zoneNumber := int((lon + 180) / 6) + 1
	if 56 <= lat && lat <= 64 && 3 <= lon && lon <= 12 {
		zoneNumber = 32
	}
	if 72 <= lat && lat <= 84 && lon >= 0 {
		if lon <= 9 {
			zoneNumber = 31
		} else if lon <= 21 {
			zoneNumber = 33
		} else if lon <= 33 {
			zoneNumber = 35
		} else if lon <= 42 {
			zoneNumber = 37
		}
	}
	central_lon := (zoneNumber - 1) * 6 - 180 + 3

	lon_rad := rad(lon)
	central_lon_rad := rad(float64(central_lon))

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
	easting = k0 * n * (a +
		a3 / 6 * (1 - lat_tan2 + c) +
		a5 / 120 * (5 - 18 * lat_tan2 + lat_tan4 + 72 * c - 58 * e_p2)) + 500000
	northing = k0 * (m + n * lat_tan * (a2 / 2 +
		a4 / 24 * (5 - lat_tan2 + 9 * c + 4 * c * c) +
		a6 / 720 * (61 - 58 * lat_tan2 + lat_tan4 + 600 * c - 330 * e_p2)))

	if lat < 0 {
		northing += 10000000
	}

	return
}