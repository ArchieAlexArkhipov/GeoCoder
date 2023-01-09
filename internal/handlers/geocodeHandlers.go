package handlers

import (
  "strconv"
  "net/http"
  "geocoder/internal/utils"
  "geocoder/internal/services"
)

var GeocodeHandler = func(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()

    addressParam, exists := params["address"]

    if !exists {
        utils.Respond(w, utils.Message(false, "There was an error in your request"))

        return
    }

    geocoder := services.GetActualGeocoder();
    data, error := geocoder.Geocode(addressParam[0])


    resp := utils.Message(true, "success")
    resp["error"] = error

    if data != nil {
        resp["lat"] = data.Lat
        resp["lng"] = data.Lng
	} else {
        resp["error"] = "got <nil> location"
	}

    utils.Respond(w, resp)
}

var ReverseGeocodeHandler = func(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()

    _, existsLat := params["lat"]
    _, existsLng := params["lng"]

    if !existsLat || !existsLng {
        utils.Respond(w, utils.Message(false, "There was an error in your request"))

        return
    }

    lat, errLat := strconv.ParseFloat(params["lat"][0], 64)
    lng, errLng := strconv.ParseFloat(params["lng"][0], 64)

    if errLat != nil || errLng != nil {
        utils.Respond(w, utils.Message(false, "There was an error in your request"))

        return
    }

    geocoder := services.GetActualGeocoder();
    data, error := geocoder.ReverseGeocode(lat, lng)


    resp := utils.Message(true, "success")
    resp["error"] = error

    if data != nil {
        resp["address"] = data.FormattedAddress
        resp["detailedAddress"] = data
    } else {
        resp["error"] = "got <nil> address"
    }

    utils.Respond(w, resp)
}
