package UTM

// Don't use it. It panic and return no error.
// FromLatLonF convert a latitude and longitude to Universal Transverse Mercator coordinates.
//
// Deprecated: Use FromLatLon functions to converse instead.
func FromLatLonF(lat, lon float64) (easting, northing float64) {
	var err error
	// Northing always false in this implementation.
	easting, northing, _, _, err = LatLonToUTM(lat, lon, false)
	if err != nil {
		panic(err)
	}
	return
}
