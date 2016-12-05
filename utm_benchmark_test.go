package UTM_test

import (
	"testing"

	"github.com/im7mortal/UTM"
)

var coordinate = UTM.Coordinate{466013, 7190568, 6, "W"}

func BenchmarkToLatLon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := coordinate.ToLatLon(); err != nil {
			b.Fatal("benchmark fatal BenchmarkToLatLon")
		}
	}
}

func BenchmarkToLatLonWithNorthern(b *testing.B) {
	coordinate.ZoneLetter = ""
	for i := 0; i < b.N; i++ {
		if _, err := coordinate.ToLatLon(true); err != nil {
			b.Fatal("benchmark fatal BenchmarkToLatLonWithNorthern")
		}
	}
}

var latLon = UTM.LatLon{64.83778, -147.71639}

func BenchmarkFromLatLon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := latLon.FromLatLon(); err != nil {
			b.Fatal("benchmark fatal BenchmarkFromLatLon")
		}
	}
}
