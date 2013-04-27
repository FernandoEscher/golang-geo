package geo

// The Geocoder Interface
// This interfaces specifies the responsibilities of a Geocoder.
// A Geocoder should be able to find a corresponding Point struct from a supplied  address.
// A Geocoder should be able to find a corresponding address when supplied with a Point struct.
type Geocoder interface {
	Geocode(query string) (*Point, error)
	ReverseGeocode(p *Point) (string, error)
}
