[![GoDoc](https://godoc.org/github.com/im7mortal/UTM?status.svg)](https://godoc.org/github.com/im7mortal/UTM)

UTM
===

Bidirectional UTM-WGS84 converter for golang. It's port from [UTM python version](https://pypi.python.org/pypi/utm) by Tobias Bieniek

Usage
-----

	go get github.com/im7mortal/UTM

Convert a (latitude, longitude) tuple into an UTM coordinate

```
	import "github.com/im7mortal/UTM"
	latLon := UTM.LatLon{50.77535, 6.08389}
	utmCoordinate := latLon.FromLatLon()
```
The return has the form
	UTMCoordinate{294409, 5628898, 32, 'U'}

Convert a (latitude, longitude) tuple into an UTM coordinate

```
	latLon := UTMCoordinate{294409, 5628898, 32, 'U'}
	utmCoordinate := latLon.ToLatLon()
```
The return has the form
	LatLon{50.77535, 6.08389}
	
	
Not implemented yet

Since the zone letter is not strictly needed for the conversion you may also
the ``northern`` parameter instead, which is a named parameter and can be set
to either ``True`` or ``False``. Have a look at the unit tests to see how it
can be used.

The UTM coordinate system is explained on

[Wikipedia page](https://en.wikipedia.org/wiki/Universal_Transverse_Mercator_coordinate_system)

Speed
-----

Development
-----------

Authors
-------

* Petr Lozhkin <im7mortal@gmail.com>
