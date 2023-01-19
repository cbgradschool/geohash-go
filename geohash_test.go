package geohash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	expectedLatitude  float64 = 40.72470580906875
	expectedLongitude float64 = -73.99975952911369
	expectedGeohash   string  = "dr5rsjen"
)

func TestEncode(t *testing.T) {
	// it accepts geographical coordinates and returns a 32 bit encoded geohash string
	geohash := Encode(expectedLatitude, expectedLongitude, 5)

	assert.Equal(t, geohash, expectedGeohash, "Geohash string should match")
}

func TestDecode(t *testing.T) {
	// it decodes a 32 bit encoded geohash string and returns geographical coordinates
	latitude, longitude := Decode(expectedGeohash)

	actual := [2]float64{latitude, longitude}

	expected := [2]float64{expectedLatitude, expectedLongitude}

	assert.Equal(t, actual, expected, "The geographical coordinates should be the same")
}
