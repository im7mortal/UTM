[![Build Status](https://travis-ci.org/im7mortal/UTM.svg)](https://travis-ci.org/im7mortal/UTM)
[![Coverage Status](https://coveralls.io/repos/im7mortal/UTM/badge.svg?branch=master)](https://coveralls.io/r/im7mortal/UTM?branch=master)
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
	Coordinate, err := latLon.FromLatLon()
```
The return has the form

	Coordinate{294409, 5628898, 32, "U"}

Convert a (latitude, longitude) tuple into an UTM coordinate

```
	coordinate := UTM.Coordinate{294409, 5628898, 32, "U"}
	latLon, err := coordinate.ToLatLon()
```
The return has the form

	LatLon{50.77535, 6.08389}
	

Since the zone letter is not strictly needed for the conversion you may also
the ``northern`` parameter instead, which is a named parameter and can be set
to either ``true`` or ``false``. In this case you should define fields clearly(!).
You can't set ZoneLetter or northern both.

```
	coordinate := UTM.UTM.Coordinate{
			Easting :		313784,
			Northing :		5427057,
			ZoneNumber :	60,
		}
	latLon, err := coordinate.ToLatLon(false)
```

The UTM coordinate system is explained on this [Wikipedia page](https://en.wikipedia.org/wiki/Universal_Transverse_Mercator_coordinate_system)

Speed
-----

Benchmark             | Amount of iterations | Average speed
--------------------- | -------------------- | -------------
ToLatLon              | 10000000             | 152 ns/op
ToLatLonWithNorthern  | 10000000             | 150 ns/op
FromLatLon            | 20000000             | 110 ns/op

> go test -bench=.

Data for comparison in oldBenchmark.txt

Development
-----------

Authors
-------

* Petr Lozhkin <im7mortal@gmail.com>

License
-------

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
