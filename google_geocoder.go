package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GoogleGeocoder struct{}

// @param [string] params.  The HTTP params in which to append to the google geocode api request. 
// @return [[]byte].  Returns a []byte representation of a JSON response
//                    From the requested method defined by the passed in paramter string 
func (g *GoogleGeocoder) Request(params string) ([]byte, error) {
	client := &http.Client{}

	fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&%s", params)

	// TODO Potentially refactor out from MapQuestGeocoder as well
	req, _ := http.NewRequest("GET", fullUrl, nil)
	resp, requestErr := client.Do(req)

	if requestErr != nil {
		panic(requestErr)
	}

	data, dataReadErr := ioutil.ReadAll(resp.Body)

	if dataReadErr != nil {
		return nil, dataReadErr
	}

	return data, nil
}

// @param [string] query.  The address to geocode.
// @return [Point] p.  The first {Point} interpreted from the google geocoding api response.
func (g *GoogleGeocoder) Geocode(query string) (*Point, error) {
	url_safe_query := url.QueryEscape(query)
	data, err := g.Request(fmt.Sprintf("address=%s", url_safe_query))
	if err != nil {
		return nil, err
	}

	lat, lng := g.extractLatLngFromResponse(data)
	p := &Point{lat: lat, lng: lng}

	return p, nil
}

// @param [[]byte] data.  The response struct from the earlier google geocode request as an array of bytes.
// @return [float64] lat.  The first point's latitude in the response. 
// @return [float64] lng.  The first point's longitude in the response. 
func (g *GoogleGeocoder) extractLatLngFromResponse(data []byte) (float64, float64) {
	res := make(map[string][]map[string]map[string]map[string]interface{}, 0)
	json.Unmarshal(data, &res)

	lat, _ := res["results"][0]["geometry"]["location"]["lat"].(float64)
	lng, _ := res["results"][0]["geometry"]["location"]["lng"].(float64)

	return lat, lng
}

// @param [Point] p.  The Point struct in which to find the address.
// @return [string] s.  The resulting address from the google geocoding api.
func (g *GoogleGeocoder) ReverseGeocode(p *Point) (string, error) {
	data, err := g.Request(fmt.Sprintf("latlng=%f,%f", p.lat, p.lng))
	if err != nil {
		return "", err
	}

	resStr := g.extractAddressFromResponse(data)

	return resStr, nil
}

// @param [[]byte] data.  The response string from the earlier google geocode request as an array of bytes.
// @return [string] resStr.  The first address from the google geocoding api.
func (g *GoogleGeocoder) extractAddressFromResponse(data []byte) string {
	res := make(map[string][]map[string]interface{}, 0)
	json.Unmarshal(data, &res)

	resStr := res["results"][0]["formatted_address"].(string)
	return resStr
}
