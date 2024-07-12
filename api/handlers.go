package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// TimeInTimeZone return a current time UTC or time in timezone from paramiter in request
func TimeInTimeZone(w http.ResponseWriter, r *http.Request) {

	timeInTimeZones := make(map[string]string)
	URLtimeZones := r.URL.Query().Get("tz")

	if len(URLtimeZones) == 0 {
		location, err := time.LoadLocation("Etc/UTC")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}

		timeInTimeZones["UTC"] = time.Now().In(location).Format("2006-01-02 15:04:05 -0700 MST")

	} else {
		timezones := strings.Split(URLtimeZones, ",")
		for _, timezone := range timezones {

			location, err := time.LoadLocation(timezone)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "invalid timezone")
				return
			}

			timeInTimeZones[timezone] = time.Now().In(location).Format("2006-01-02 15:04:05 -0700 MST")
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeInTimeZones)
}
