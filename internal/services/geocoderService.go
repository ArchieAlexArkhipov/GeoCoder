package services

import (
	"fmt"
	"os"
    "geocoder/internal/components/geocoder"
    "geocoder/internal/components/geocoder/yandex"
)

const (
	addr         = "Melbourne VIC"
	lat, lng     = -37.813611, 144.963056
	radius       = 50
	zoom         = 18
	addrFR       = "Champs de Mars Paris"
	latFR, lngFR = 48.854395, 2.304770
)

func GetActualGeocoder() geocoder.Geocoder {
    return yandex.Geocoder(os.Getenv("YANDEX_API_KEY"))
}

func Geocode(geocoder geocoder.Geocoder, address string) {
    location, _ := geocoder.Geocode(address)

	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", addr, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
}

func ReverseGeocode(geocoder geocoder.Geocoder, lat, lng float64) {
    address, _ := geocoder.ReverseGeocode(lat, lng)

    if address != nil {
        fmt.Printf("Address of (%.6f,%.6f) is %s\n", lat, lng, address.FormattedAddress)
        fmt.Printf("Detailed address: %#v\n", address)
    } else {
        fmt.Println("got <nil> address")
    }
}
