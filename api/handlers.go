package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// TimeInAnotherTimeZone return a current time UTC or time in timezone from paramiter in request
func TimeInAnotherTimeZone(w http.ResponseWriter, r *http.Request) {

	timeInTimezones := make(map[string]string)
	URLtimezones := r.URL.Query().Get("tz")

	if URLtimezones == "" {
		location, err := time.LoadLocation("UTC")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}

		timeInTimezones["Etc/UTC"] = time.Now().In(location).Format("2006-01-02 15:04:05 -0700 MST")

	} else {
		timezones := strings.Split(URLtimezones, ",")
		for _, timezone := range timezones {

			location, err := time.LoadLocation(timezone)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "invalid timezone")
				return
			}

			timeInTimezones[timezone] = time.Now().In(location).Format("2006-01-02 15:04:05 -0700 MST")
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeInTimezones)
}
