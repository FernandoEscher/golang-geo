package geo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

// Asserts that golang-geo rextracts the correct address 
// From a known successful google geocode response.
func TestExtractAddressFromResponse(t *testing.T) {
	g := &GoogleGeocoder{}

	data, err := GetMockResponse("test/helpers/google_reverse_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	address := g.extractAddressFromResponse(data)
	if address != "285 Bedford Avenue, Brooklyn, NY 11211, USA" {
		t.Error(fmt.Sprintf("Expected: 285 Bedford Avenue, Brooklyn, NY 11211 USA.  Got: %s", address))
	}
}

// Asserts that golang-geo extracts the correct Lat and Lng values 
// From a known successful google geocoder response.
func TestExtractLatLngFromRequest(t *testing.T) {
	g := &GoogleGeocoder{}

	data, err := GetMockResponse("test/helpers/google_geocode_success.json")
	if err != nil {
		t.Error("%v\n", err)
	}

	lat, lng := g.extractLatLngFromResponse(data)
	if lat != 37.615223 && lng != -122.389979 {
		t.Error(fmt.Sprintf("Expected: [37.615223, -122.389979], Got: [%f, %f]", lat, lng))
	}
}

// Returns a byte array that represents the response of the google api request
func GetMockResponse(s string) ([]byte, error) {
	dataPath := path.Join(s)
	_, readErr := os.Stat(dataPath)
	if readErr != nil && os.IsNotExist(readErr) {
		return nil, readErr
	}

	handler, handlerErr := os.Open(dataPath)
	if handlerErr != nil {
		return nil, handlerErr
	}

	data, readErr := ioutil.ReadAll(handler)
	if readErr != nil {
		return nil, readErr
	}

	return data, nil
}
