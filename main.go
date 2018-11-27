package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"strings"
)


// ============================================================
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/targomo-isochrone/{lng}/{lat}/{time}/{key}", v1TargomoIsochrone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", router))

}
// ============================================================


// ============================================================
func v1TargomoIsochrone (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var jsonResult map[string]string

	if isochrone, msg := v1DoTargomoIsochrone(params["lng"], params["lat"], params["time"], params["key"]); msg == "" {
		jsonResult = map[string]string{"targomo": isochrone}
	} else {
		jsonResult = map[string]string{"targomo": msg}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonResult)
}
// ============================================================


// ============================================================
func getRegion(xLng string, yLat string) (region string, msg string) {

	lng, err := strconv.ParseFloat(xLng, 64)
	if err != nil {
		region = "none" 
		msg = "invalid lng"
		return
	}
	lat, err := strconv.ParseFloat(yLat, 64)
	if err != nil {
		region = "none" 
		msg = "invalid lat"
		return
	}

	msg = ""

	if lng <= -100 && lat <= 36 {
		region = "southwest"
	} else if lng <= -100 && lat >= 36 {
		region = "northwest"
	} else if lng >= -100 && lat >= 36 {
		region = "northeast"
	} else if lng >= -100 && lat <= 36 {
		region = "southeast"
	} else {
		msg = "unable to determine region"
	}

	return
}
// ============================================================


// ============================================================
func v1DoTargomoIsochrone(sxLng string, syLat string, sTime string, sKey string) (geojson string, msg string) {

	if region, error_msg := getRegion(sxLng, syLat); error_msg == "" { 
	
		// r360_key := os.Getenv("TARGOMO")
    	r360_url := "https://service.route360.net/na_" +
			region + "/v1/polygon?cfg={'sources':[{'lat':" + 
			syLat + ",'lng':" + sxLng + 
			",'id':'Mappy','tm':{'car':{}}}],'polygon':" +
			"{'serializer':'geojson','srid':'4326'," +
			"'values':[" + sTime + "],'buffer':.002,'quadrantSegments':8}}&key=" + sKey

		startSearchText := ",\"coordinates\":"
		endSearchText   := ",\"properties\":{\"time\""

		geojson = ""
		msg     = ""

		response, err := http.Get(r360_url)
		if err == nil {
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				geojson = ""
				msg     = err.Error()
			} 

			jsonText := string(body)

			nStart   := strings.Index(jsonText, startSearchText) + len(startSearchText)
			nEnd     := strings.Index(jsonText, endSearchText)

			x := strings.Split(jsonText[nStart:nEnd], ",")

			var s []string
			var lng string
			var lat string
			var prefix int
			var suffix int

			for n := 0; n < len(x); n+=2 {
				prefix = strings.Count(x[n], "[")
				lat    = strings.Replace(strings.Replace(strings.Replace(x[n+1], "[", "", -1), "]", "", -1), "}", "", -1)
				lng    = strings.Replace(strings.Replace(x[n], "[", "", -1), "]", "", -1)
				suffix = strings.Count(x[n+1], "]")

				s = append(s, strings.Repeat("[", prefix) + lat + "," + lng + strings.Repeat("]", suffix))
			}
	
			geojson = strings.Join(s, ",")

		} 
	
	} else {
		msg = error_msg
	}

	return
}
// ============================================================


