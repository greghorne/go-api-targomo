package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	"os"
	// "database/sql"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"fmt"
	"strings"
)


// ============================================================
func main() {
	
	fmt.Println("main...")
	router := mux.NewRouter()
	router.HandleFunc("/v1/targomo-isochrone/{lng}/{lat}/{time}/{key}", v1TargomoIsochrone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
// ============================================================


// ============================================================
func v1TargomoIsochrone (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	if isochrone, msg := v1DoTargomoIsochrone(params["lng"], params["lat"], params["time"], params["key"]); msg == "" {
		//  success
		fmt.Println(isochrone)
		fmt.Println("msg: " + msg)
	} else {
		//error
		fmt.Println(isochrone)
		fmt.Println("msg: " + msg)
	}

	// jsonResult  := map[string]bool{"intersects": bIntersects}

	json.NewEncoder(w).Encode("name:greg")

}
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
func v1DoTargomoIsochrone(sxLng string, syLat string, sTime string, sKey string) (polygons string, msg string) {

	fmt.Println("v1DoTargomoIsochrone...")

	if region, error_msg := getRegion(sxLng, syLat); error_msg == "" { 
	
		r360_key := os.Getenv("TARGOMO")

    	r360_url := "https://service.route360.net/na_" +
			region + "/v1/polygon?cfg={'sources':[{'lat':" + 
			syLat + ",'lng':" + sxLng + 
			",'id':'Mappy','tm':{'car':{}}}],'polygon':" +
			"{'serializer':'geojson','srid':'4326'," +
			"'values':[" + sTime + "],'buffer':.002,'quadrantSegments':8}}&key=" + r360_key

		response, err := http.Get(r360_url)
		if err == nil {
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {} 

			fmt.Println("myText=====")
			myText := string(body)
			fmt.Println(myText)

			fmt.Println("=====")
			nStart := strings.Index(myText, "geometry\":") + len("geometry\":")
			nEnd := strings.Index(myText, ",\"properties\":{\"time\"")
fmt.Println(nStart, nEnd, len(myText))
			myText2 := myText[nStart:nEnd]
			fmt.Println(myText2)

// fmt.Println(len(text), nEnd)
// fmt.Println(nStart + len("geometry\":"), len(text) - nEnd)

// text2 := text[nStart + len("geometry\":"):]
// 			text3 := text2[:len(text2) - nEnd]
// 			fmt.Println(text3)



		} else {
			fmt.Println(err)
			polygons = ""
			msg = "some error"
		}

		// fmt.Println(polygons, msg)
	
	} else {
		polygons = ""
		msg = error_msg
	}
				  
	// fmt.Println(msg)

	return

}
// ============================================================


