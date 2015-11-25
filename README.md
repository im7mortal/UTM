utm
===

Bidirectional UTM-WGS84 converter for golang

Usage
-----

	> go get github.com/im7mortal/UTM

Convert a (latitude, longitude) tuple into an UTM coordinate

```
	import "github.com/im7mortal/UTM"
	latLon := UTM.LAT_LON{50.77535, 6.08389}
	utmCoordinate := latLon.FROM_LATLON()
```
  > UTM_COORDINATE{294409, 5628898, 32, 'U'}

The syntax is **utm.from_latlon(LATITUDE, LONGITUDE)**.

The return has the form **(EASTING, NORTHING, ZONE NUMBER, ZONE LETTER)**.

Convert an UTM coordinate into a (latitude, longitude) tuple::

  utm.to_latlon(340000, 5710000, 32, 'U')
  >>> (51.51852098408468, 6.693872395145327)

The syntax is **utm.to_latlon(EASTING, NORTHING, ZONE NUMBER, ZONE LETTER)**.

The return has the form **(LATITUDE, LONGITUDE)**.

Since the zone letter is not strictly needed for the conversion you may also
the ``northern`` parameter instead, which is a named parameter and can be set
to either ``True`` or ``False``. Have a look at the unit tests to see how it
can be used.

The UTM coordinate system is explained on
`this <https://en.wikipedia.org/wiki/Universal_Transverse_Mercator_coordinate_system>`_
Wikipedia page.

Speed
-----

The library has been compared to the more generic pyproj library by running the
unit test suite through pyproj instead of utm. These are the results:

* with pyproj (without projection cache): 4.0 - 4.5 sec
* with pyproj (with projection cache): 0.9 - 1.0 sec
* with utm: 0.4 - 0.5 sec

Development
-----------

Create a new ``virtualenv`` and install the library via ``pip install -e .``.
After that install the ``pytest`` package via ``pip install pytest`` and run
the unit test suite by calling ``py.test``.

Authors
-------

* Petr Lozhkin <im7mortal@gmail.com>
